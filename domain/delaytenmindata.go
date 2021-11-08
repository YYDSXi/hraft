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

var clientDelayTenMin *clientv3.Client

//处理延时数据
func DelayTenMinBlockOrPackage(client *clientv3.Client, LEDGER_TYPE string) {
	clientDelayTenMin = client
	changeTenMinBlockIndex := ((utils.GetIndexMinInt()/10 - 1) + 144) % 144
	switch LEDGER_TYPE {
	case LEDGER_TYPE_VIDEO:
		if TenMinBlockChangeVideo[changeTenMinBlockIndex] {
			go DealDelayTenMinData(LEDGER_TYPE_VIDEO, TenMinBlockChangeVideo, changeTenMinBlockIndex)
		}
	case LEDGER_TYPE_USER_BEHAVIOR:
		if TenMinBlockChangeUserBehavior[changeTenMinBlockIndex] {
			go DealDelayTenMinData(LEDGER_TYPE_USER_BEHAVIOR, TenMinBlockChangeUserBehavior, changeTenMinBlockIndex)
		}
	case LEDGER_TYPE_NODE_CREDIBLE:
		if TenMinBlockChangeNodeCredible[changeTenMinBlockIndex] {
			go DealDelayTenMinData(LEDGER_TYPE_NODE_CREDIBLE, TenMinBlockChangeNodeCredible, changeTenMinBlockIndex)
		}
	case LEDGER_TYPE_SENSOR:
		if TenMinBlockChangeSensor[changeTenMinBlockIndex] {
			go DealDelayTenMinData(LEDGER_TYPE_SENSOR, TenMinBlockChangeSensor, changeTenMinBlockIndex)
		}
	case LEDGER_TYPE_SERVICE_ACCESS:
		if TenMinBlockChangeServiceAccess[changeTenMinBlockIndex] {
			go DealDelayTenMinData(LEDGER_TYPE_SERVICE_ACCESS, TenMinBlockChangeServiceAccess, changeTenMinBlockIndex)
		}
	}
}

func DealDelayTenMinData(LEDGER_TYPE string, minBlockChangeArray []bool, changeTenMinBlockIndex int) {

	//构建前一个十分钟块的key
	//年月日 账本类型 链类型 十分钟索引
	//2006-01-02:账本类型:链类型:增强块索引
	var yearMonthDay string
	if changeTenMinBlockIndex == 143 { //说明打包的是上一天的最后一个增强块
		yearMonthDay = time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	} else { //说明是打包当天增强块
		yearMonthDay = time.Now().Format("2006-01-02")
	}
	preKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(changeTenMinBlockIndex)

	//构建前前一个十分钟块的key  主要作用是得到 前前一分钟块的BlockHash 填充到前一分钟块中的 PreBlockHash 字段
	prePreKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa((changeTenMinBlockIndex-1+144)%144)
	getResponse := utils.GetData(clientDelayTenMin, prePreKeyString, RequestTimeout)

	//填充剩余字段
	//上一个十分钟key对应的value值
	var value []byte
	switch LEDGER_TYPE {
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
			//keyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa(changeTenMinBlockIndex*10+j)
			//获取每一分钟块数据（从数据库中）
			// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
			// for _, ev := range getResponse.Kvs {
			// 	err := json.Unmarshal(ev.Value, &minBlock)
			// 	if err != nil {
			// 		log.Error("反解析", err)
			// 	}
			// }
			//从交易分钟块文件中获取
			minBlock = utils.ReadTxMinFiletoTenmin(yearMonthDay, LEDGER_TYPE, strconv.Itoa(changeTenMinBlockIndex*10+j))
			minBlocksArray = append(minBlocksArray, &minBlock)
		}
		prePreBlockHash := "default"
		if prePreTenMinuteTxBlock.BlockHash != "" {
			prePreBlockHash = prePreTenMinuteTxBlock.BlockHash
		}
		//账本类型 链类型 前一区块hash 块高度 key值
		tenMinuteBlock := utils.InitToTdengineTxTenMinBlock(LEDGER_TYPE,
			BLOCK_TYPE_TENMINUT, prePreBlockHash, changeTenMinBlockIndex, preKeyString)
		tenMinuteBlock.Blocks = minBlocksArray
		tenMinuteBlock = utils.FillTdengineTxTenMinBlockRemainingFieldsAndDataReceipts(tenMinuteBlock)
		value, _ = json.Marshal(tenMinuteBlock)

		//存入文件
		//preKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(changeTenMinBlockIndex)

		utils.WriteTxblocktoTenminfile(yearMonthDay, LEDGER_TYPE, strconv.Itoa(changeTenMinBlockIndex), tenMinuteBlock)
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
			//keyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa(changeTenMinBlockIndex*10+j)
			//获取每一分钟块数据(从数据库中)
			// getResponse := utils.GetData(clientDelayTenMin, keyString, RequestTimeout)
			// for _, ev := range getResponse.Kvs {
			// 	err := json.Unmarshal(ev.Value, &minBlock)
			// 	if err != nil {
			// 		log.Error("反解析", err)
			// 	}
			// }
			//从存证分钟快文件中获取
			minBlock = utils.ReadReMinFiletoTenmin(yearMonthDay, LEDGER_TYPE, strconv.Itoa(changeTenMinBlockIndex*10+j))
			minBlocksArray = append(minBlocksArray, &minBlock)
		}
		prePreBlockHash := "default"
		if prePreTenMinuteDataBlock.BlockHash != "" {
			prePreBlockHash = prePreTenMinuteDataBlock.BlockHash
		}
		tenMinuteBlock := utils.InitToTdengineDataTenMinBlock(LEDGER_TYPE,
			BLOCK_TYPE_TENMINUT, prePreBlockHash, changeTenMinBlockIndex, preKeyString)
		tenMinuteBlock.Blocks = minBlocksArray
		tenMinuteBlock = utils.FillTdengineDataTenMinBlockRemainingFieldsAndDataReceipts(tenMinuteBlock)
		value, _ = json.Marshal(tenMinuteBlock)
		//存入文件
		//preKeyString := yearMonthDay + KeySplit + LEDGER_TYPE + KeySplit + BLOCK_TYPE_TENMINUT + KeySplit + strconv.Itoa(changeTenMinBlockIndex)

		utils.WriteReblocktoTenminfile(yearMonthDay, LEDGER_TYPE, strconv.Itoa(changeTenMinBlockIndex), tenMinuteBlock)

	}
	//存到etcd
	//utils.PutData(clientDelayTenMin, preKeyString, string(value), RequestTimeout)

	log.Infof("%s账本稳定增强块排序打包成功!", LEDGER_TYPE)
	log.Info("稳定增强块key=", string(preKeyString))
	log.Info("稳定增强块value=", string(value))

	minBlockChangeArray[changeTenMinBlockIndex] = false
}
