package domain

import (
	"encoding/json"
	"fmt"
	pb "hraft/rpc"
	"hraft/utils"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func AutoCreateTenMinBlockToEtcd(clientDelayTenMin *clientv3.Client) {

	//遍历账本
	for i := 0; i < len(GlobalLedgerArray); i++ {
		//遍历账本
		indexMinInt := utils.GetIndexMinInt()

		//打包一次增强块，使得对应增强块表示位置false
		ResetTenMinBlockChange(GlobalLedgerArray[i], indexMinInt)

		//构建前一个十分钟块的key
		//年月日 账本类型 链类型 十分钟索引
		//2006-01-02:账本类型:链类型:增强块索引

		var yearMonthDay string
		if indexMinInt < 10 { //说明打包的是上一天的最后一个增强块
			yearMonthDay = time.Now().Add(-time.Hour * 24).Format("2006-01-02")
		} else { //说明是打包当天增强块
			yearMonthDay = time.Now().Format("2006-01-02")
		}
		//构建增强块的key
		preKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa((indexMinInt/10-1+144)%144)

		//构建前前一个十分钟块的key  主要作用是得到 前前一分钟块的BlockHash 填充到前一分钟块中的 PreBlockHash 字段
		prePreKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa((indexMinInt/10-2+144)%144)
		getResponse := utils.GetData(clientDelayTenMin, prePreKeyString, RequestTimeout)

		//填充剩余字段
		//上一个十分钟key对应的value值
		//var value []byte
		start := time.Now().UnixNano()
		switch GlobalLedgerArray[i] {

		case LEDGER_TYPE_NODE_CREDIBLE, LEDGER_TYPE_SENSOR, LEDGER_TYPE_SERVICE_ACCESS:
			//获取上上一个十分钟块
			var prePreTenMinuteTxBlock pb.TenMinuteTxBlock
			for _, ev := range getResponse.Kvs {
				err := json.Unmarshal(ev.Value, &prePreTenMinuteTxBlock)
				if err != nil {
					log.Error("pb.TenMinuteBlock Unmarshal err", err)
				}
			}

			//定义 前十分钟 每分钟分钟块数据
			var minBlocksArray []*pb.MinuteTxBlock
			for j := 0; j < 10; j++ {
				var minBlock pb.MinuteTxBlock
				//keyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa((((indexMinInt/10)-1+144)%144)*10+j)
				//获取每一分钟块数据(从数据库中读取)
				// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
				// for _, ev := range getResponse.Kvs {
				// 	err := json.Unmarshal(ev.Value, &minBlock)
				// 	if err != nil {
				// 		log.Error("反解析", err)
				// 	}
				// }
				//获取每一分钟块数据
				minBlock = utils.ReadTxMinFiletoTenmin(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((((indexMinInt/10)-1+144)%144)*10+j))
				minBlocksArray = append(minBlocksArray, &minBlock)
			}
			prePreBlockHash := "default"
			if prePreTenMinuteTxBlock.BlockHash != "" {
				prePreBlockHash = prePreTenMinuteTxBlock.BlockHash
			}
			tenMinuteBlock := utils.InitToTdengineTxTenMinBlock(GlobalLedgerArray[i],
				BLOCK_TYPE_TENMINUT, prePreBlockHash, (indexMinInt/10-1+144)%144, preKeyString)
			tenMinuteBlock.Blocks = minBlocksArray
			tenMinuteBlock = utils.FillTdengineTxTenMinBlockRemainingFieldsAndDataReceipts(tenMinuteBlock)
			//value, _ = json.Marshal(tenMinuteBlock)
			//存入文件
			//preKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa((indexMinInt/10-1+144)%144)

			utils.WriteTxblocktoTenminfile(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((indexMinInt/10-1+144)%144), tenMinuteBlock)

		case LEDGER_TYPE_VIDEO, LEDGER_TYPE_USER_BEHAVIOR:

			//获取上上一个十分钟块
			var prePreTenMinuteDataBlock pb.TenMinuteDataBlock
			for _, ev := range getResponse.Kvs {
				err := json.Unmarshal(ev.Value, &prePreTenMinuteDataBlock)
				if err != nil {
					log.Error("pb.TenMinuteBlock Unmarshal err", err)
				}
			}
			//定义装 前十分钟 每分钟分钟块数据
			var minBlocksArray []*pb.MinuteDataBlock
			for j := 0; j < 10; j++ {
				var minBlock pb.MinuteDataBlock
				//keyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa((((indexMinInt/10)-1+144)%144)*10+j)
				//获取每一分钟块数据从数据库中)
				// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
				// for _, ev := range getResponse.Kvs {
				// 	err := json.Unmarshal(ev.Value, &minBlock)
				// 	if err != nil {
				// 		log.Error("反解析", err)
				// 	}
				// }
				//从分钟文件中读取数据
				minBlock = utils.ReadReMinFiletoTenmin(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((((indexMinInt/10)-1+144)%144)*10+j))
				minBlocksArray = append(minBlocksArray, &minBlock)
			}
			prePreBlockHash := "default"
			if prePreTenMinuteDataBlock.BlockHash != "" {
				prePreBlockHash = prePreTenMinuteDataBlock.BlockHash
			}
			tenMinuteBlock := utils.InitToTdengineDataTenMinBlock(GlobalLedgerArray[i],
				BLOCK_TYPE_TENMINUT, prePreBlockHash, (indexMinInt/10-1+144)%144, preKeyString)
			tenMinuteBlock.Blocks = minBlocksArray
			tenMinuteBlock = utils.FillTdengineDataTenMinBlockRemainingFieldsAndDataReceipts(tenMinuteBlock)
			//value, _ = json.Marshal(tenMinuteBlock)
			//存入文件
			//preKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa((indexMinInt/10-1+144)%144)

			utils.WriteReblocktoTenminfile(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((indexMinInt/10-1+144)%144), tenMinuteBlock)

		}

		//增强块存到etcd
		//utils.PutData(clientDelayTenMin, preKeyString, string(value), RequestTimeout)

		log.Infof("%s账本增强块排序打包成功!", GlobalLedgerArray[i])
		log.Info("增强块key=", string(preKeyString))
		//log.Info("增强块value=", string(value))
		end := time.Now().UnixNano()
		fmt.Printf("打包%s增强块总用时：%v毫秒\n", string(preKeyString), (end-start)/1000000)
	}
}

//打包一次增强块，使得对应增强块表示位置false
func ResetTenMinBlockChange(LEDGER_TYPE string, indexMinInt int) {
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		TenMinBlockChangeVideo[(indexMinInt/10-1+144)%144] = false
	case LEDGER_TYPE_USER_BEHAVIOR:
		TenMinBlockChangeUserBehavior[(indexMinInt/10-1+144)%144] = false
	case LEDGER_TYPE_NODE_CREDIBLE:
		TenMinBlockChangeNodeCredible[(indexMinInt/10-1+144)%144] = false
	case LEDGER_TYPE_SENSOR:
		TenMinBlockChangeSensor[(indexMinInt/10-1+144)%144] = false
	case LEDGER_TYPE_SERVICE_ACCESS:
		TenMinBlockChangeServiceAccess[(indexMinInt/10-1+144)%144] = false
	}
}
