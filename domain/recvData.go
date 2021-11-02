package domain

import (
	"encoding/json"
	"hraft/dataStruct"
	pb "hraft/rpc"
	"hraft/utils"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/chenhg5/collection"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{} //定义一个server结构体

var clientRedis *clientv3.Client //etcd的客户端？？？

func RedisServerMain(client *clientv3.Client) { //redis服务端
	clientRedis = client
	//遍历开启端口
	go StartGrpcPort(":" + Port)
}

//遍历开启端口
func StartGrpcPort(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("开启端口失败: ", err)
	}
	log.Infof("%s端口开启成功！", Port)

	s := grpc.NewServer()
	pb.RegisterToUpperServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Error("端口服务调用失败: ", err)
	}
}

//grpc测试接口
func (s *server) Upper(ctx context.Context, in *pb.UpperRequest) (*pb.UpperReply, error) {
	log.Info("测试rpc，接收到数据:", in.Name)
	return &pb.UpperReply{Message: strings.ToUpper(in.Name)}, nil
}

//Video账本类型  存证数据
func (s *server) Video(ctx context.Context, in *pb.VideoData) (*pb.Response, error) {
	log.Info("Video账本类型，接收到数据: ", in.DataReceipts)
	go ToEtcdDbDataReceipt(in.DataReceipts, ALL_LEDGER_TYPE_ARRAY[0])
	return &pb.Response{ErrCode: SuccessCode, ErrMsg: ""}, nil
}

//UserBehaviour账本类型  存证数据
func (s *server) UserBehaviour(ctx context.Context, in *pb.UserBehaviourData) (*pb.Response, error) {
	log.Info("UserBehaviour账本类型，接收到数据:", in.DataReceipts)
	go ToEtcdDbDataReceipt(in.DataReceipts, ALL_LEDGER_TYPE_ARRAY[1])
	return &pb.Response{ErrCode: SuccessCode, ErrMsg: ""}, nil
}

//NodeCredible账本类型  交易数据
func (s *server) NodeCredible(ctx context.Context, in *pb.NodeCredibleData) (*pb.Response, error) {
	log.Info("NodeCredible账本类型，接收到数据:", in.Transactions)
	go ToEtcdDbTransaction(in.Transactions, ALL_LEDGER_TYPE_ARRAY[2])
	return &pb.Response{ErrCode: SuccessCode, ErrMsg: ""}, nil
}

//Sensor账本类型  交易数据
func (s *server) Sensor(ctx context.Context, in *pb.SensorData) (*pb.Response, error) {
	log.Info("Sensor账本类型，接收到数据:", in.Transactions)
	go ToEtcdDbTransaction(in.Transactions, ALL_LEDGER_TYPE_ARRAY[3])
	return &pb.Response{ErrCode: SuccessCode, ErrMsg: ""}, nil
}

//ServiceAccess账本类型  交易数据
func (s *server) ServiceAccess(ctx context.Context, in *pb.ServiceAccessData) (*pb.Response, error) {
	log.Info("ServiceAccess账本类型，接收到数据:", in.Transactions)
	go ToEtcdDbTransaction(in.Transactions, ALL_LEDGER_TYPE_ARRAY[4])
	return &pb.Response{ErrCode: SuccessCode, ErrMsg: ""}, nil
}

//将接收到的存证数据存入数据库
//[]*proto.DataReceipt
func ToEtcdDbDataReceipt(structArray []*pb.DataReceipt, LEDGER_TYPE string) {

	for i := 0; i < len(structArray); i++ {

		if structArray[i].CreateTimestamp != "" && structArray[i].KeyId != "" {

			dataReceipt := structArray[i] //申请的dataReceipt变量为存证数据结构
			if !utils.TimeTrue(dataReceipt.CreateTimestamp) {
				log.Error("时间戳不合法", dataReceipt.CreateTimestamp)
				log.Error("正确格式为:2021-06-01 12:12:12.123")
				continue
			}
			//2021-04-20 17:18:10.123  get 2021-04-20 and minIndex
			dayTime, indexMinInt := utils.GetMinIntByTimeStamp(dataReceipt.CreateTimestamp)

			//构建每条数据的key   2021-04-20 17:18:10.123 # 账本类型 # KeyId
			perDataKeyString := dataReceipt.CreateTimestamp + TIMESTAMP_KEYID + LEDGER_TYPE + TIMESTAMP_KEYID + dataReceipt.KeyId

			//计算延时时间
			delayMinTimeInt := utils.GetIndexMinInt() - indexMinInt
			//如果是10分钟之前的数据 过滤掉不处理
			if delayMinTimeInt > utils.Conf.Consensus.CommonConfig.Timeout {
				log.Error("数据", structArray[i])
				log.Error("时间戳超过阈值，不做处理！", dataReceipt.CreateTimestamp)
				continue
			} else if delayMinTimeInt > 1 {
				//延时数据里存到数组里的key是2021-04-20 17:18:10.123 # 账本类型 # KeyId
				//临时存储原来里面的数组
				var array interface{}
				switch {
				case LEDGER_TYPE == LEDGER_TYPE_VIDEO:
					array, _ = DelayLedgerVideo.Get(delayMinTimeInt) //list结构，该结构包含追加、更新、删除等方法;获取该延迟索引位置的数据
					array = append(array.([]string), perDataKeyString)
					DelayLedgerVideo.Update(delayMinTimeInt, array)
				case LEDGER_TYPE == LEDGER_TYPE_USER_BEHAVIOR:
					array, _ = DelayLedgerUserBehavior.Get(delayMinTimeInt)
					array = append(array.([]string), perDataKeyString)
					DelayLedgerUserBehavior.Update(delayMinTimeInt, array)
				}
				//这里将数据存入数据库，应该改为用临时存储
				delayDataReceiptByteArray, _ := json.Marshal(dataReceipt)                                       //将数据结构编码成json
				utils.PutData(clientRedis, perDataKeyString, string(delayDataReceiptByteArray), RequestTimeout) //delayDataReceiptByteArray[]byte类型，转化成string类型便于查看
				log.Info("接收到延时数据")
				log.Info("key=", perDataKeyString)
				log.Info("val=", string(delayDataReceiptByteArray))
				//更新变量，整个服务数据量大小
				utils.StatisticalAllDataCounts(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				utils.StatisticalAllDataSize(clientRedis, LEDGER_TYPE, len(string(delayDataReceiptByteArray)), RequestTimeout, false)
				//延时数据 叠加
				utils.StatisticalCurDayDelayData(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				continue
			}

			//处理同一时间戳内 有多个值切片
			//dataReceipt.CreateTimestamp = utils.DevTimestampGenerateIndex(clientRedis, dataReceipt.CreateTimestamp, RequestTimeout)

			//构建MD数据key 2021-04-20 : MD : 账本类型 : 分钟索引 : 节点ID
			//MD里面存2021-04-20 : 账本类型 : 分钟索引 : 节点ID
			keyMDString := dayTime + KeySplit + "MD" + KeySplit + LEDGER_TYPE + KeySplit + strconv.Itoa(indexMinInt) + KeySplit + strconv.Itoa(int(GlobalLeaderId))

			//获取本分钟已存etcd的MD数据
			getResponse := utils.GetData(clientRedis, keyMDString, RequestTimeout)

			var minuteData dataStruct.MinuteData
			for _, ev := range getResponse.Kvs {
				err := json.Unmarshal(ev.Value, &minuteData)
				if err != nil {
					log.Error("存证数据反序列化【dataStruct.MinuteData】失败：", err)
				}
			}
			receiptTimeStampLedgerTypeKeyIds := minuteData.ReceiptTimeStampLedgerTypeKeyId
			if receiptTimeStampLedgerTypeKeyIds == nil {
				receiptTimeStampLedgerTypeKeyIds = make([]string, 0)
			}

			//去重
			if !collection.Collect(receiptTimeStampLedgerTypeKeyIds).Contains(perDataKeyString) {

				//每条数据添加进去
				receiptTimeStampLedgerTypeKeyIds = append(receiptTimeStampLedgerTypeKeyIds, perDataKeyString)
				//这里将数据存入数据库，应该改为用临时存储
				dataReceiptByteArray, _ := json.Marshal(dataReceipt)
				utils.PutData(clientRedis, perDataKeyString, string(dataReceiptByteArray), RequestTimeout)

				log.Info("存证数据 key = ", perDataKeyString)
				log.Info("存证数据 val = ", string(dataReceiptByteArray))
				log.Info("数据成功存储到Etcd！")

				minuteData.ReceiptTimeStampLedgerTypeKeyId = receiptTimeStampLedgerTypeKeyIds

				minuteDataByteArray, _ := json.Marshal(minuteData)

				//更新存储时间戳的 数据结构体
				utils.PutData(clientRedis, keyMDString, string(minuteDataByteArray), RequestTimeout)

				//更新变量，整个服务数据量大小
				utils.StatisticalAllDataCounts(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				utils.StatisticalAllDataSize(clientRedis, LEDGER_TYPE, len(string(dataReceiptByteArray)), RequestTimeout, false)

				//更新创始块  部分字段 如数据量大小
				go func() {
					for i := 0; i < len(BLOCK_TYPE_ARRAY); i++ {
						preGenesisBlock := GetPreGenesisBlock(LEDGER_TYPE, BLOCK_TYPE_ARRAY[i])
						utils.UpdateCurrentDayDataCounts(clientRedis, LEDGER_TYPE, BLOCK_TYPE_ARRAY[i], int(preGenesisBlock.DataCounts), RequestTimeout)
						utils.UpdateCurrentDayDataSize(clientRedis, LEDGER_TYPE, BLOCK_TYPE_ARRAY[i], int(preGenesisBlock.DataSize), RequestTimeout)
						//UpdateGenesisBlockToEtcdAndTdengine(dayTime+KeySplit+LEDGER_TYPE+KeySplit+BLOCK_TYPE_ARRAY[i], len(string(minuteDataByteArray)), LEDGER_TYPE, BLOCK_TYPE_ARRAY[i])
					}
				}()
			}
		} else {
			log.Info("structArray[i]=", structArray[i])
			if structArray[i].CreateTimestamp != "" {
				log.Error("存证数据时间Timestamp为空，此条数据存储失败")
			}
			if structArray[i].KeyId != "" {
				log.Error("存证数据KeyId为空，此条数据存储失败")
			}
		}
	} //for

	//处理延时数据 并标识变化块属于哪个分钟块
	AutoDealDelayDataAndUpdateMinBlock(clientRedis, LEDGER_TYPE)
	time.Sleep(time.Duration(3) * time.Second)
	//表示增强块是否需要重新打包
	DelayTenMinBlockOrPackage(clientRedis, LEDGER_TYPE)
	time.Sleep(time.Duration(3) * time.Second)
	//表示天块是否需要重新打包
	DelayDailyBlockOrPackage(clientRedis, LEDGER_TYPE)
}

//交易数据
func ToEtcdDbTransaction(structArray []*pb.Transaction, LEDGER_TYPE string) {

	for i := 0; i < len(structArray); i++ {

		if structArray[i].CreateTimestamp != "" && structArray[i].TransactionId != "" {

			transaction := structArray[i]
			//判断时间戳是否合法
			if !utils.TimeTrue(transaction.CreateTimestamp) {
				log.Error("时间戳不合法", transaction.CreateTimestamp)
				log.Error("正确格式为:2021-06-01 12:12:12.123")
				continue
			}
			//2021-04-20 17:18:10.123  get 2021-04-20 and minIndex
			dayTime, indexMinInt := utils.GetMinIntByTimeStamp(transaction.CreateTimestamp)

			//构建每条数据的key   2021-04-20 17:18:10.123 # 账本类型 # 交易ID
			perDataKeyString := transaction.CreateTimestamp + TIMESTAMP_KEYID + LEDGER_TYPE + TIMESTAMP_KEYID + transaction.TransactionId

			//计算延时时间
			delayMinTimeInt := utils.GetIndexMinInt() - indexMinInt
			//如果是10分钟之前的数据 过滤掉不处理
			if delayMinTimeInt > 10 {
				log.Info("时间戳超过阈值，不做处理！", transaction.CreateTimestamp)
				continue
			} else if delayMinTimeInt > 1 {
				//延时数据里存的是2021-04-20 17:18:10.123 # 账本类型 # 交易ID
				//临时存储原来里面的数组
				var array interface{}
				switch {
				case LEDGER_TYPE == LEDGER_TYPE_NODE_CREDIBLE:
					array, _ = DelayLedgerNodeCredible.Get(delayMinTimeInt)
					array = append(array.([]string), perDataKeyString)
					DelayLedgerNodeCredible.Update(delayMinTimeInt, array)
				case LEDGER_TYPE == LEDGER_TYPE_SENSOR:
					array, _ = DelayLedgerSensor.Get(delayMinTimeInt)
					array = append(array.([]string), perDataKeyString)
					DelayLedgerSensor.Update(delayMinTimeInt, array)
				case LEDGER_TYPE == LEDGER_TYPE_SERVICE_ACCESS:
					array, _ = DelayLedgerServiceAccess.Get(delayMinTimeInt)
					array = append(array.([]string), perDataKeyString)
					DelayLedgerServiceAccess.Update(delayMinTimeInt, array)
				}
				delayTransactionByteArray, _ := json.Marshal(transaction)
				utils.PutData(clientRedis, perDataKeyString, string(delayTransactionByteArray), RequestTimeout)
				log.Info("接收到延时数据")
				log.Info("key=", perDataKeyString)
				log.Info("val=", string(delayTransactionByteArray))
				//更新变量，整个服务数据量大小
				utils.StatisticalAllDataCounts(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				utils.StatisticalAllDataSize(clientRedis, LEDGER_TYPE, len(string(delayTransactionByteArray)), RequestTimeout, false)
				//延时数据 叠加
				utils.StatisticalCurDayDelayData(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				continue
			}

			//处理同一时间戳内 有多个值切片由2021-04-20 17:18:10.123   变成2021-04-20 17:18:10.123001
			//transaction.CreateTimestamp = utils.DevTimestampGenerateIndex(clientRedis, transaction.CreateTimestamp, RequestTimeout)

			//构建MD数据key 2021-04-20 : MD : 账本类型 : 分钟索引 : 节点ID
			//MD里面存 2021-04-20 17:18:10.123 # 账本类型 # 交易ID
			keyMDString := dayTime + KeySplit + "MD" + KeySplit + LEDGER_TYPE + KeySplit + strconv.Itoa(indexMinInt) + KeySplit + strconv.Itoa(int(GlobalLeaderId))

			//获取本分钟已存etcd的MD数据
			getResponse := utils.GetData(clientRedis, keyMDString, RequestTimeout)

			var minuteData dataStruct.MinuteData
			for _, ev := range getResponse.Kvs {
				//fmt.Printf("%s -> %s\n", ev.Key, ev.Value)
				err := json.Unmarshal(ev.Value, &minuteData)
				if err != nil {
					log.Error("交易数据dataStruct.MinuteData反序列化失败：", err)
				}
			}
			transactionTimeStampLedgerTypeTxIds := minuteData.TransactionTimeStampLedgerTypeTxId
			if transactionTimeStampLedgerTypeTxIds == nil {
				transactionTimeStampLedgerTypeTxIds = make([]string, 0)
			}
			//把每笔数据存到etcd key为2021-04-20 17:18:10.123 # 账本类型 # 交易ID
			transactionByteArray, _ := json.Marshal(transaction)
			utils.PutData(clientRedis, perDataKeyString, string(transactionByteArray), RequestTimeout)

			//去重
			if !collection.Collect(transactionTimeStampLedgerTypeTxIds).Contains(perDataKeyString) {
				transactionTimeStampLedgerTypeTxIds = append(transactionTimeStampLedgerTypeTxIds, perDataKeyString)
				//排序
				//sort.Strings(transactionTimeStamp)

				utils.PutData(clientRedis, perDataKeyString, string(transactionByteArray), RequestTimeout)

				log.Info("交易数据 key = ", perDataKeyString)
				log.Info("交易数据 val = ", string(transactionByteArray))
				log.Info("数据成功存储到Etcd！")

				minuteData.TransactionTimeStampLedgerTypeTxId = transactionTimeStampLedgerTypeTxIds
				minuteDataByteArray, _ := json.Marshal(minuteData)

				//将对应分钟数据 更新好后 存到etcd
				utils.PutData(clientRedis, keyMDString, string(minuteDataByteArray), RequestTimeout)

				//更新变量，整个服务数据量大小
				utils.StatisticalAllDataCounts(clientRedis, LEDGER_TYPE, 1, RequestTimeout, false)
				utils.StatisticalAllDataSize(clientRedis, LEDGER_TYPE, len(string(transactionByteArray)), RequestTimeout, false)

				//更新当天统计字段
				go func() {
					for i := 0; i < len(BLOCK_TYPE_ARRAY); i++ {
						preGenesisBlock := GetPreGenesisBlock(LEDGER_TYPE, BLOCK_TYPE_ARRAY[i])
						utils.UpdateCurrentDayDataCounts(clientRedis, LEDGER_TYPE, BLOCK_TYPE_ARRAY[i], int(preGenesisBlock.DataCounts), RequestTimeout)
						utils.UpdateCurrentDayDataSize(clientRedis, LEDGER_TYPE, BLOCK_TYPE_ARRAY[i], int(preGenesisBlock.DataSize), RequestTimeout)
					}
				}()
			}
		} else {
			log.Info("structArray[i]=", structArray[i])
			if structArray[i].CreateTimestamp != "" {
				log.Error("交易数据时间戳.Timestamp为空，此条数据存储失败")
			}
			if structArray[i].TransactionId != "" {
				log.Error("交易数据Id为空，此条数据存储失败")
			}

		}
	} //for
	//处理延时数据 并标识变化块属于哪个分钟块
	AutoDealDelayDataAndUpdateMinBlock(clientRedis, LEDGER_TYPE)
	time.Sleep(time.Duration(3) * time.Second)
	//表示增强块是否需要重新打包
	DelayTenMinBlockOrPackage(clientRedis, LEDGER_TYPE)
	time.Sleep(time.Duration(3) * time.Second)
	//表示天块是否需要重新打包
	DelayDailyBlockOrPackage(clientRedis, LEDGER_TYPE)
}

func GetPreGenesisBlock(ledgerType string, blockType string) pb.GenesisBlock {
	//获取上一天的创世区块，将其相关字段取出来
	//年月日 账本类型 链类型
	preGenesisBlockKeyString := time.Now().Add(-time.Hour*24).Format("2006-01-02") + KeySplit + ledgerType + KeySplit + blockType
	//判断创始块是否已存在
	isExistPreMDResponse := utils.GetData(clientRedis, preGenesisBlockKeyString, RequestTimeout)
	var isExistPreGenesisBlock pb.GenesisBlock
	for _, ev := range isExistPreMDResponse.Kvs {
		err := json.Unmarshal(ev.Value, &isExistPreGenesisBlock)
		if err != nil {
			log.Error("反序列化结构体，pb.GenesisBlock失败：", err)
		}
	}
	return isExistPreGenesisBlock
}

//每次更新创始链块 数据大小
func UpdateGenesisBlockToEtcdAndTdengine(key string, dataSize int, LEDGER_TYPE string, BLOCK_TYPE string) {
	getResponse := utils.GetData(clientRedis, key, RequestTimeout)
	var genesisBlock pb.GenesisBlock
	for _, ev := range getResponse.Kvs {
		err := json.Unmarshal(ev.Value, &genesisBlock)
		if err != nil {
			log.Error("UpdateGenesiBlockToEtcdAndTdengine函数，pb.GenesisBlock反序列化失败：", err)
		}
	}
	//如果对应创始块不存在 则新建
	if genesisBlock.CreateTimestamp == "" {
		genesisBlock = utils.CreateGenesisBlock(key, LEDGER_TYPE, BLOCK_TYPE, 0, 0, "default")
	}
	genesisBlock.DataCounts += int32(1)
	genesisBlock.DataSize += int64(dataSize)

	genesisBlockByteArray, _ := json.Marshal(genesisBlock)
	genesisBlock.GenesisBlockHash = utils.GetStringMD5(string(genesisBlockByteArray))

	genesisBlockByteArray, _ = json.Marshal(genesisBlock)
	utils.PutData(clientRedis, key, string(genesisBlockByteArray), RequestTimeout)

	//respString := utils.PutGenesisBlockToTdengine(genesisBlock, LEDGER_TYPE)
	//log.Info("PutGenesisBlockToTdengine函数，返回结果：",respString)
}
