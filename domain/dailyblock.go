package domain

import (
	"encoding/json"
	pb "hraft/rpc"
	"hraft/utils"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func AutoCreateDailyBlockToEtcd(clientDelayTenMin *clientv3.Client) {
	//遍历账本
	for i := 0; i < len(GlobalLedgerArray); i++ {
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
		preKeyString := preDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)
		prePreKeyString := prePreDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24))

		//得到 前前一天块的BlockHash 填充到前一天块中的 PreBlockHash 字段
		getResponse := utils.GetData(clientDelayTenMin, prePreKeyString, RequestTimeout)

		var value []byte
		switch GlobalLedgerArray[i] {
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
				keyString := preDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(j)
				var tenMinuteBlock pb.TenMinuteTxBlock
				//根据时间戳 获取原来在etcd中的块数据
				getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
				for _, ev := range getResponse.Kvs {
					err := json.Unmarshal(ev.Value, &tenMinuteBlock)
					if err != nil {
						log.Error("genesisBlock Unmarshal err", err)
					}
				}
				tenMinBlocksInterfaceArray = append(tenMinBlocksInterfaceArray, &tenMinuteBlock)
			}

			//初始化天块
			//参数分别是  账本类型 链类型 前前一天Blockash 块高度 key
			dailyBlock := utils.InitToTdengineTxDailyBlock(GlobalLedgerArray[i], BLOCK_TYPE_DAY, prePreDailyBlock.BlockHash, int(time.Now().Sub(t).Hours()/24)+1, preKeyString)
			dailyBlock.Blocks = tenMinBlocksInterfaceArray
			//填充剩余字段
			dailyBlock = utils.FillTdengineTxDailyBlockRemainingFieldsAndDataReceipts(dailyBlock)
			value, _ = json.Marshal(dailyBlock)
			//存入文件
			//preKeyString := preDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)

			utils.WriteTxblocktoDayfile(preDay, GlobalLedgerArray[i], strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1), dailyBlock)

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
				keyString := preDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(j)
				var tenMinuteBlock pb.TenMinuteDataBlock
				//根据时间戳 获取原来在etcd中的块数据
				getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
				for _, ev := range getResponse.Kvs {
					err := json.Unmarshal(ev.Value, &tenMinuteBlock)
					if err != nil {
						log.Error("genesisBlock Unmarshal err", err)
					}
				}
				tenMinBlocksInterfaceArray = append(tenMinBlocksInterfaceArray, &tenMinuteBlock)
			}

			//初始化天块
			//参数分别是  账本类型 链类型 前前一天Blockash 块高度 key
			dailyBlock := utils.InitToTdengineDataDailyBlock(GlobalLedgerArray[i], BLOCK_TYPE_DAY, prePreDailyBlock.BlockHash, int(time.Now().Sub(t).Hours()/24)+1, preKeyString)
			dailyBlock.Blocks = tenMinBlocksInterfaceArray
			//填充剩余字段
			dailyBlock = utils.FillTdengineDataDailyBlockRemainingFieldsAndDataReceipts(dailyBlock)
			value, _ = json.Marshal(dailyBlock)
			//存入文件
			//preKeyString := preDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_DAY + KeySplit + strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1)

			utils.WriteReblocktoDayfile(preDay, GlobalLedgerArray[i], strconv.Itoa(int(time.Now().Sub(t).Hours()/24)+1), dailyBlock)

		}
		//存到etcd
		utils.PutData(clientDelayTenMin, preKeyString, string(value), RequestTimeout)

		log.Infof("%s账本天块排序打包成功!", GlobalLedgerArray[i])
		log.Info("天块key=", string(preKeyString))
		log.Info("天块value=", string(value))

		//存到Tdengine
		////utils.PutDailyBlockToTdengine(dailyBlock, GlobalLedgerArray[i])
	}
}
