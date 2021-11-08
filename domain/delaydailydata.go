package domain

import (
	"encoding/json"
	pb "hraft/rpc"
	"hraft/utils"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	log "github.com/sirupsen/logrus"
)

var clientDelayDaily *clientv3.Client

//处理延时数据
func DelayDailyBlockOrPackage(client *clientv3.Client, LEDGER_TYPE string) {
	clientDelayDaily = client
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		if DailyChangeVideo {
			go DealDelayDailyData(LEDGER_TYPE_VIDEO, DailyChangeVideo)
		}
	case LEDGER_TYPE_USER_BEHAVIOR:
		if DailyChangeUserBehavior {
			go DealDelayDailyData(LEDGER_TYPE_USER_BEHAVIOR, DailyChangeUserBehavior)
		}
	case LEDGER_TYPE_NODE_CREDIBLE:
		if DailyChangeNodeCredible {
			go DealDelayDailyData(LEDGER_TYPE_NODE_CREDIBLE, DailyChangeNodeCredible)
		}
	case LEDGER_TYPE_SENSOR:
		if DailyChangeSensor {
			go DealDelayDailyData(LEDGER_TYPE_SENSOR, DailyChangeSensor)
		}
	case LEDGER_TYPE_SERVICE_ACCESS:
		if DailyChangeServiceAccess {
			go DealDelayDailyData(LEDGER_TYPE_SERVICE_ACCESS, DailyChangeServiceAccess)
		}
	}
}

func DealDelayDailyData(LEDGER_TYPE string, DailyChange bool) {

	//构建前一天 年月日
	preDayTime := time.Now().AddDate(0, 0, -1)
	preDay := preDayTime.Format("2006-01-02")
	//构建前前一天 年月日
	prePreDayTime := time.Now().AddDate(0, 0, -2)
	prePreDay := prePreDayTime.Format("2006-01-02")

	//获取天块高度 即当前天距离公元多少天
	t, _ := time.Parse("2006-01-02", "0000-00-00")

	//构建key
	//年月日 账本类型 链类型 高度
	preKeyString := preDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)
	prePreKeyString := prePreDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24))

	//得到 前前一天块的BlockHash 填充到前一天块中的 PreBlockHash 字段
	getResponse := utils.GetData(clientDelayTenMin, prePreKeyString, RequestTimeout)

	var value []byte
	switch LEDGER_TYPE {
	case LEDGER_TYPE_NODE_CREDIBLE, LEDGER_TYPE_SENSOR, LEDGER_TYPE_SERVICE_ACCESS:

		var prePreDailyBlock pb.DailyTxBlock
		for _, ev := range getResponse.Kvs {
			err := json.Unmarshal(ev.Value, &prePreDailyBlock)
			if err != nil {
				log.Error("genesisBlock Unmarshal err", err)
			}
		}
		var tenMinBlocksInterfaceArray []*pb.TenMinuteTxBlock
		//遍历144个增强块 填充到tenMinBlocksArray
		for j := 0; j < 144; j++ {
			//keyString := preDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(j)
			var tenMinuteBlock pb.TenMinuteTxBlock
			//根据时间戳 获取原来在etcd中的块数据(从数据库中获取)
			// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
			// for _, ev := range getResponse.Kvs {
			// 	err := json.Unmarshal(ev.Value, &tenMinuteBlock)
			// 	if err != nil {
			// 		log.Error("genesisBlock Unmarshal err", err)
			// 	}
			// }
			//从交易增强块文件中获取
			tenMinuteBlock = utils.ReadTxTenMinFiletoDay(preDay, LEDGER_TYPE, strconv.Itoa(j))
			tenMinBlocksInterfaceArray = append(tenMinBlocksInterfaceArray, &tenMinuteBlock)
		}

		//初始化天块
		//参数分别是  账本类型 链类型 前前一天Blockash 块高度 key
		dailyBlock := utils.InitToTdengineTxDailyBlock(LEDGER_TYPE, BLOCK_TYPE_DAY, prePreDailyBlock.BlockHash, int(time.Now().Sub(t).Hours()/24)+1, preKeyString)
		dailyBlock.Blocks = tenMinBlocksInterfaceArray
		//填充剩余字段
		dailyBlock = utils.FillTdengineTxDailyBlockRemainingFieldsAndDataReceipts(dailyBlock)
		value, _ = json.Marshal(dailyBlock)
		//存入文件
		//preKeyString := preDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)

		utils.WriteTxblocktoDayfile(preDay, LEDGER_TYPE, strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1), dailyBlock)

	case LEDGER_TYPE_VIDEO, LEDGER_TYPE_USER_BEHAVIOR:

		var prePreDailyBlock pb.DailyDataBlock
		for _, ev := range getResponse.Kvs {
			err := json.Unmarshal(ev.Value, &prePreDailyBlock)
			if err != nil {
				log.Error("genesisBlock Unmarshal err", err)
			}
		}
		var tenMinBlocksInterfaceArray []*pb.TenMinuteDataBlock
		//遍历144个增强块 填充到tenMinBlocksArray
		for j := 0; j < 144; j++ {
			//keyString := preDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(j)
			var tenMinuteBlock pb.TenMinuteDataBlock
			//根据时间戳 获取原来在etcd中的块数据（从数据库中）
			// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
			// for _, ev := range getResponse.Kvs {
			// 	err := json.Unmarshal(ev.Value, &tenMinuteBlock)
			// 	if err != nil {
			// 		log.Error("genesisBlock Unmarshal err", err)
			// 	}
			// }
			//从存证增强块文件中获取
			tenMinuteBlock = utils.ReadReTenMinFiletoDay(preDay, LEDGER_TYPE, strconv.Itoa(j))
			tenMinBlocksInterfaceArray = append(tenMinBlocksInterfaceArray, &tenMinuteBlock)
		}

		//初始化天块
		//参数分别是  账本类型 链类型 前前一天Blockash 块高度 key
		dailyBlock := utils.InitToTdengineDataDailyBlock(LEDGER_TYPE, BLOCK_TYPE_DAY, prePreDailyBlock.BlockHash, int(time.Now().Sub(t).Hours()/24)+1, preKeyString)
		dailyBlock.Blocks = tenMinBlocksInterfaceArray
		//填充剩余字段
		dailyBlock = utils.FillTdengineDataDailyBlockRemainingFieldsAndDataReceipts(dailyBlock)
		value, _ = json.Marshal(dailyBlock)
		//存入文件
		//preKeyString := preDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)

		utils.WriteReblocktoDayfile(preDay, LEDGER_TYPE, strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1), dailyBlock)
	}
	//存到etcd
	//utils.PutData(clientDelayTenMin, preKeyString, string(value), RequestTimeout)

	log.Infof("%s账本天块排序打包成功!", LEDGER_TYPE)
	log.Info("天块key=", string(preKeyString))
	log.Info("天块value=", string(value))

	DailyChange = false
}
