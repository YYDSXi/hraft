package domain

import (
	"encoding/json"
	"fmt"
	"hraft/dataStruct"
	pb "hraft/rpc"
	"hraft/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chenhg5/collection"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var clientDelayMin *clientv3.Client

//处理延时数据
func AutoDealDelayDataAndUpdateMinBlock(client *clientv3.Client, LEDGER_TYPE string) {
	clientDelayMin = client
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		go DealDelayData(DelayLedgerVideo, LEDGER_TYPE_VIDEO)
	case LEDGER_TYPE_USER_BEHAVIOR:
		go DealDelayData(DelayLedgerUserBehavior, LEDGER_TYPE_USER_BEHAVIOR)
	case LEDGER_TYPE_NODE_CREDIBLE:
		go DealDelayData(DelayLedgerNodeCredible, LEDGER_TYPE_NODE_CREDIBLE)
	case LEDGER_TYPE_SENSOR:
		go DealDelayData(DelayLedgerSensor, LEDGER_TYPE_SENSOR)
	case LEDGER_TYPE_SERVICE_ACCESS:
		go DealDelayData(DelayLedgerServiceAccess, LEDGER_TYPE_SERVICE_ACCESS)
	}
}

func DealDelayData(ledgerList utils.List, LEDGER_TYPE string) {

	for i := 0; i < ledgerList.Size(); i++ {
		arrayInterface, _ := ledgerList.Get(i)
		if arrayInterface == nil {
			continue
		}
		//将[]interface 类型转换为 []string类型
		//同一个array数组里面 存储的是同一分钟的延时数据
		array := arrayInterface.([]string)
		//该一分钟没有延时数据
		if len(array) == 0 {
			continue
		}
		//数组里存的是每条数据的key   2021-04-20 17:18:10.123 # 账本类型 # KeyId
		timeStamp := strings.Split(array[0], TIMESTAMP_KEYID)[0]
		//2021-04-20 17:18:10.123  get 2021-04-20 and minIndex
		yearMonthDay, indexMinInt := utils.GetMinIntByTimeStamp(timeStamp)

		//表示是那个分钟索引的数据变化
		WhichTenMinBlockChange(indexMinInt, LEDGER_TYPE)
		//上一天的天块的数据是否变化
		WhichDailyBlockChange(indexMinInt, LEDGER_TYPE)

		//构建MD的key 里面存2021-04-20 : 账本类型 : 分钟索引 : 节点ID
		keyMDString := yearMonthDay + KeySplit + "MD" + KeySplit + LEDGER_TYPE + KeySplit + strconv.Itoa(indexMinInt) + KeySplit + strconv.Itoa(int(GlobalLeaderId))
		//获取延时数据所属 MD数据
		getResponse := utils.GetData(clientRedis, keyMDString, RequestTimeout)
		var minuteData dataStruct.MinuteData
		for _, ev := range getResponse.Kvs {
			err := json.Unmarshal(ev.Value, &minuteData)
			if err != nil {
				log.Error("存证数据反序列化【dataStruct.MinuteData】失败：", err)
			}
		}

		//构建延时数据 本来该属于那个分钟块 分钟快的Key
		//年月日 账本类型 链类型 分钟索引
		minBlockKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa(indexMinInt)
		//获取延时数据 本来该属于那个分钟块 分钟块 结构体
		getMinBlockResponse := utils.GetData(clientDelayMin, minBlockKeyString, RequestTimeout)

		//判断是存证还是交易
		switch LEDGER_TYPE {

		case LEDGER_TYPE_VIDEO, LEDGER_TYPE_USER_BEHAVIOR:
			//获取延时数据所属 MD数据 获取存证切片
			receiptTimeStampLedgerTypeKeyIds := minuteData.ReceiptTimeStampLedgerTypeKeyId
			if receiptTimeStampLedgerTypeKeyIds == nil {
				receiptTimeStampLedgerTypeKeyIds = make([]string, 0)
			}

			//将原来MD数据 拼接延时数据
			//新增数据记录数
			newInsertDataCount := 0
			for k := 0; k < len(array); k++ {
				if !collection.Collect(receiptTimeStampLedgerTypeKeyIds).Contains(array[k]) {
					newInsertDataCount++
					receiptTimeStampLedgerTypeKeyIds = append(receiptTimeStampLedgerTypeKeyIds, array[k])
				}
			}

			//排序
			sort.Strings(receiptTimeStampLedgerTypeKeyIds)

			//将新的MD数据存到etcd
			minuteData.ReceiptTimeStampLedgerTypeKeyId = receiptTimeStampLedgerTypeKeyIds
			minuteDataByteArray, _ := json.Marshal(minuteData)
			utils.PutData(clientDelayMin, keyMDString, string(minuteDataByteArray), RequestTimeout)

			//获取延时数据 所属的分钟块(从etcd中取出)
			var minuteBlock pb.MinuteDataBlock
			for _, ev := range getMinBlockResponse.Kvs {
				err := json.Unmarshal(ev.Value, &minuteBlock)
				if err != nil {
					log.Error("pb.MinuteBlock Unmarshal err", err)
				}
			}
			//新建存储每条数据key的数组
			dataReceiptsArrays := []*pb.DataReceipt{}

			temp := ""
			num := 1
			var currentDataSize int64 //数据量大小
			//将延时数据 和 运来已经存在的数据 遍历排序
			for j := 0; j < len(receiptTimeStampLedgerTypeKeyIds); j++ {
				getPerDataResponse := utils.GetData(clientDelayMin, receiptTimeStampLedgerTypeKeyIds[j], RequestTimeout)
				var perDataReceipt pb.DataReceipt
				for _, ev := range getPerDataResponse.Kvs {
					err := json.Unmarshal(ev.Value, &perDataReceipt)
					if err != nil {
						log.Error("pb.MinuteBlock Unmarshal err", err)
					}
				}

				timeStamp := strings.Split(receiptTimeStampLedgerTypeKeyIds[j], TIMESTAMP_KEYID)[0]
				//将同一时间戳的数据排序
				if timeStamp != temp {
					num = 1
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					perDataReceipt.CreateTimestamp = ans
					temp = timeStamp
				} else {
					num++
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					perDataReceipt.CreateTimestamp = ans
				}

				dataReceiptsArrays = append(dataReceiptsArrays, &perDataReceipt)
				dataReceiptByteArray, _ := json.Marshal(perDataReceipt)
				currentDataSize += int64(len(string(dataReceiptByteArray)))

				//将原数组数据 置空
				ledgerList.Update(j, make([]string, 0))
			}
			minuteBlock.DataReceipts = dataReceiptsArrays
			//如果区块头为空
			if minuteBlock.Header == nil {
				//年月日 账本类型 分钟索引
				blockHeadKeyString := time.Now().Format("2006-01-02") + KeySplit + LEDGER_TYPE + KeySplit + strconv.Itoa(indexMinInt)
				//判断区块头是否已存在
				isExistMDResponse := utils.GetData(clientDelayMin, blockHeadKeyString, RequestTimeout)
				var isExistBlockHeader pb.BlockHeader
				for _, ev := range isExistMDResponse.Kvs {
					err := json.Unmarshal(ev.Value, &isExistBlockHeader)
					if err != nil {
						log.Error("preMinuteMDData Unmarshal err", err)
					}
				}
				minuteBlock.Header = &isExistBlockHeader
			}
			//填充剩余字段
			//参数分别是 结构体 目前数据记录量 目前数据量 transactions结构体
			minuteBlock.Header = utils.FillTdengineMinBlockHeader(minuteBlock.Header, LEDGER_TYPE, BLOCK_TYPE_MIN, int64(newInsertDataCount), currentDataSize)
			minuteBlockByteArray, _ := json.Marshal(minuteBlock)
			//存入文件
			//	minBlockKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa(indexMinInt)

			utils.WriteReblocktominfile(yearMonthDay, LEDGER_TYPE, strconv.Itoa(indexMinInt), minuteBlock)

			//存到etcd
			utils.PutData(clientDelayMin, minBlockKeyString, string(minuteBlockByteArray), RequestTimeout)

			log.Infof("%s账本稳定分钟块打包成功！", LEDGER_TYPE)
			log.Info("存到etcd：key = ", minBlockKeyString)
			log.Info("存到etcd：val = ", string(minuteBlockByteArray))
		case LEDGER_TYPE_NODE_CREDIBLE, LEDGER_TYPE_SENSOR, LEDGER_TYPE_SERVICE_ACCESS:
			//获取延时数据所属 MD数据 获取存证切片
			transactionTimeStampLedgerTypeTxIds := minuteData.TransactionTimeStampLedgerTypeTxId
			if transactionTimeStampLedgerTypeTxIds == nil {
				transactionTimeStampLedgerTypeTxIds = make([]string, 0)
			}
			log.Info("keyMDString=", keyMDString)
			log.Info("transactionTimeStampLedgerTypeTxIds=", transactionTimeStampLedgerTypeTxIds)

			//将原来MD数据 拼接延时数据
			//新增数据记录数
			newInsertDataCount := 0
			for k := 0; k < len(array); k++ {
				if !collection.Collect(transactionTimeStampLedgerTypeTxIds).Contains(array[k]) {
					newInsertDataCount++
					transactionTimeStampLedgerTypeTxIds = append(transactionTimeStampLedgerTypeTxIds, array[k])
				}
			}
			log.Info("array=", array)
			log.Info("拼接后transactionTimeStampLedgerTypeTxIds=", transactionTimeStampLedgerTypeTxIds)

			//排序
			sort.Strings(transactionTimeStampLedgerTypeTxIds)
			//将新的MD数据存到etcd
			minuteData.TransactionTimeStampLedgerTypeTxId = transactionTimeStampLedgerTypeTxIds
			minuteDataByteArray, _ := json.Marshal(minuteData)
			utils.PutData(clientDelayMin, keyMDString, string(minuteDataByteArray), RequestTimeout)

			//获取延时数据所属分钟块
			var minuteBlock pb.MinuteTxBlock
			for _, ev := range getMinBlockResponse.Kvs {
				err := json.Unmarshal(ev.Value, &minuteBlock)
				if err != nil {
					log.Error("pb.MinuteBlock Unmarshal err", err)
				}
			}
			//新建存储每条数据key的数组
			transactionsArrays := []*pb.Transaction{}

			temp := ""
			num := 1
			var currentDataSize int64 //数据量大小
			//把数组里的所有key取出来对应数据
			for j := 0; j < len(transactionTimeStampLedgerTypeTxIds); j++ {
				getPerTxResponse := utils.GetData(clientDelayMin, transactionTimeStampLedgerTypeTxIds[j], RequestTimeout)
				var perDataTx pb.Transaction
				for _, ev := range getPerTxResponse.Kvs {
					err := json.Unmarshal(ev.Value, &perDataTx)
					if err != nil {
						log.Error("pb.MinuteBlock Unmarshal err", err)
					}
				}

				timeStamp := strings.Split(transactionTimeStampLedgerTypeTxIds[j], TIMESTAMP_KEYID)[0]
				//将同一时间戳的数据排序
				if timeStamp != temp {
					num = 1
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					perDataTx.CreateTimestamp = ans
					temp = timeStamp
				} else {
					num++
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					perDataTx.CreateTimestamp = ans
				}

				transactionsArrays = append(transactionsArrays, &perDataTx)
				transactionsByteArray, _ := json.Marshal(perDataTx)
				currentDataSize += int64(len(string(transactionsByteArray)))

				//将原数组数据 置空
				ledgerList.Update(j, make([]string, 0))
			}

			minuteBlock.Transactions = transactionsArrays

			//如果区块头为空
			if minuteBlock.Header == nil {
				//年月日 账本类型 分钟索引
				blockHeadKeyString := time.Now().Format("2006-01-02") + KeySplit + LEDGER_TYPE + KeySplit + strconv.Itoa(indexMinInt)
				//判断区块头是否已存在
				isExistMDResponse := utils.GetData(clientDelayMin, blockHeadKeyString, RequestTimeout)
				var isExistBlockHeader pb.BlockHeader
				for _, ev := range isExistMDResponse.Kvs {
					err := json.Unmarshal(ev.Value, &isExistBlockHeader)
					if err != nil {
						log.Error("preMinuteMDData Unmarshal err", err)
					}
				}
				minuteBlock.Header = &isExistBlockHeader
			}

			//填充剩余字段
			//参数分别是 结构体 目前数据记录量 目前数据量 transactions结构体
			minuteBlock.Header = utils.FillTdengineMinBlockHeader(minuteBlock.Header, LEDGER_TYPE, BLOCK_TYPE_MIN, int64(newInsertDataCount), currentDataSize)
			minuteBlockByteArray, _ := json.Marshal(minuteBlock)
			//存入文件
			//	minBlockKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa(indexMinInt)

			utils.WriteTxblocktominfile(yearMonthDay, LEDGER_TYPE, strconv.Itoa(indexMinInt), minuteBlock)
			//存到etcd
			utils.PutData(clientDelayMin, minBlockKeyString, string(minuteBlockByteArray), RequestTimeout)

			log.Infof("%s账本稳定分钟块打包成功！", LEDGER_TYPE)
			log.Info("稳定分钟块存到etcd：key = ", minBlockKeyString)
			log.Info("稳定分钟块存到etcd：val = ", string(minuteBlockByteArray))
		}
	}
}

//表示哪个增强块要变化
func WhichTenMinBlockChange(indexMinInt int, LEDGER_TYPE string) {
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		TenMinBlockChangeVideo[(indexMinInt/10+144)%144] = true
	case LEDGER_TYPE_USER_BEHAVIOR:
		TenMinBlockChangeUserBehavior[(indexMinInt/10+144)%144] = true
	case LEDGER_TYPE_NODE_CREDIBLE:
		TenMinBlockChangeNodeCredible[(indexMinInt/10+144)%144] = true
	case LEDGER_TYPE_SENSOR:
		TenMinBlockChangeSensor[(indexMinInt/10+144)%144] = true
	case LEDGER_TYPE_SERVICE_ACCESS:
		TenMinBlockChangeServiceAccess[(indexMinInt/10+144)%144] = true
	}
}

//表示上一天的天块是否变化
func WhichDailyBlockChange(indexMinInt int, LEDGER_TYPE string) {
	if indexMinInt < 1430 {
		return
	}
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		DailyChangeVideo = true
	case LEDGER_TYPE_USER_BEHAVIOR:
		DailyChangeUserBehavior = true
	case LEDGER_TYPE_NODE_CREDIBLE:
		DailyChangeNodeCredible = true
	case LEDGER_TYPE_SENSOR:
		DailyChangeSensor = true
	case LEDGER_TYPE_SERVICE_ACCESS:
		DailyChangeServiceAccess = true
	}
}
