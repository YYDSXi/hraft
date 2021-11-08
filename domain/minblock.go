package domain

import (
	"encoding/json"
	"fmt"
	pb "hraft/rpc"
	"hraft/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

//自动任务创建分钟块
func AutoCreateMinBlockToEtcd(clientDelayTenMin *clientv3.Client) {

	//遍历账本
	for i := 0; i < len(GlobalLedgerArray); i++ {

		//获取当前分钟索引
		indexMinInt := utils.GetIndexMinInt()

		var yearMonthDay string
		if indexMinInt == 0 { //说明打包的是上一天的最后一个增强块
			yearMonthDay = time.Now().Add(-time.Hour * 24).Format("2006-01-02")
		} else { //说明是打包当天增强块
			yearMonthDay = time.Now().Format("2006-01-02")
		}

		//构建前一分钟未打包数据的key 获取原来在etcd中的数据
		//2021-04-20 : MD : 账本类型 : 分钟索引 : 节点ID
		preKeyMDString := yearMonthDay + KeySplit + "MD" + KeySplit + GlobalLedgerArray[i] + KeySplit + strconv.Itoa((indexMinInt-1+1440)%1440) + KeySplit + strconv.Itoa(int(GlobalLeaderId))

		//原有的获取分钟数据
		// getMDResponse := utils.GetData(clientDelayTenMin, preKeyMDString, RequestTimeout)
		// var preMinuteMDData dataStruct.MinuteData
		// for _, ev := range getMDResponse.Kvs {
		// 	err := json.Unmarshal(ev.Value, &preMinuteMDData)
		// 	if err != nil {
		// 		log.Error("preMinuteMDData Unmarshal err", err)
		// 	}
		// }

		//构建前一分钟MinBlock的key
		//年月日：账本类型：链类型：分钟索引
		preKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa((indexMinInt-1+1440)%1440)

		//获取前前一分钟块 为了得到Blockash 赋值到前一分钟的preBlockash字段
		prePreKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_MIN + KeySplit + strconv.Itoa((indexMinInt-2+1440)%1440)
		getPrePreResponse := utils.GetData(clientDelayTenMin, prePreKeyString, RequestTimeout)

		//获取前前一分钟结构体
		var prePreMinuteBlock pb.MinuteDataBlock
		for _, ev := range getPrePreResponse.Kvs {
			err := json.Unmarshal(ev.Value, &prePreMinuteBlock)
			if err != nil {
				log.Error("pb.MinuteBlock Unmarshal err", err)
			}
		}

		//获取前区块头
		getBlockHeadKeyString := yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + strconv.Itoa((indexMinInt-1+1440)%1440)

		getPreBlockHeadResponse := utils.GetData(clientDelayTenMin, getBlockHeadKeyString, RequestTimeout)
		var blockHeader = pb.BlockHeader{}
		for _, ev := range getPreBlockHeadResponse.Kvs {
			err := json.Unmarshal(ev.Value, &blockHeader)
			if err != nil {
				log.Error("pb.DataReceipt Unmarshal err", err)
			}
		}
		if blockHeader.CreateTimestamp == "" {

			blockHeader.CreateTimestamp = utils.GetNowOneMinTimeStamp()
		}
		//判断当前账本是存证还是交易
		//存证
		if GlobalLedgerArray[i] == LEDGER_TYPE_VIDEO || GlobalLedgerArray[i] == LEDGER_TYPE_USER_BEHAVIOR {

			//初始化分钟块
			//参数分别是 账本类型 链类型 块高度 key
			var minBlockToTdengine = pb.MinuteDataBlock{} //结构体=块头+数据
			timeCorrect := utils.GetNowOneMinTimeStamp()
			blockHeader.UpdateTimestamp = timeCorrect

			if prePreMinuteBlock.Header != nil {
				blockHeader.PreBlockHash = prePreMinuteBlock.Header.BlockHash
			}
			minBlockToTdengine.Header = &blockHeader
			dataReceipts := minBlockToTdengine.DataReceipts
			receiptTimeStampLedgerTypeKeyIds := MDData[preKeyMDString]

			//排序
			sort.Strings(receiptTimeStampLedgerTypeKeyIds)

			temp := ""
			num := 1
			//当前数据量大小
			var currentDataSize int
			var blockDatacounts int
			//遍历每个时间戳  获取数据
			for j := 0; j < len(receiptTimeStampLedgerTypeKeyIds); j++ {
				//根据时间戳 获取原来在etcd中的块数据

				var dataReceipt pb.DataReceipt
				//原有的从数据库中读取具体数据
				// getResponse := utils.GetData(clientDelayTenMin, receiptTimeStampLedgerTypeKeyIds[j], RequestTimeout)
				// for _, ev := range getResponse.Kvs {
				// 	err := json.Unmarshal(ev.Value, &dataReceipt)
				// 	if err != nil {
				// 		log.Error("pb.DataReceipt Unmarshal err", err)
				// 	}
				// }
				ReceiptDatamu.RLock()
				getResponse := ReceiptData[receiptTimeStampLedgerTypeKeyIds[j]]
				ReceiptDatamu.RUnlock()
				err := json.Unmarshal([]byte(getResponse), &dataReceipt)
				if err != nil {
					log.Error("pb.DataReceipt Unmarshal err", err)
				}

				timeStamp := strings.Split(receiptTimeStampLedgerTypeKeyIds[j], TIMESTAMP_KEYID)[0]
				//将同一时间戳的数据排序
				if timeStamp != temp {
					num = 1
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					dataReceipt.CreateTimestamp = ans
					temp = timeStamp
				} else {
					num++
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					dataReceipt.CreateTimestamp = ans
				}

				log.Info("构建的时间戳为：", dataReceipt.CreateTimestamp)
				dataReceipts = append(dataReceipts, &dataReceipt)
				dataReceiptByteArray, _ := json.Marshal(dataReceipt)
				//数据量大小叠加
				currentDataSize += len(string(dataReceiptByteArray))
				blockDatacounts = blockDatacounts + 1
			}

			//填充剩余字段
			//参数分别是 结构体 目前数据记录量 目前数据量 dataReceipts结构体
			minBlockToTdengine.Header = utils.FillTdengineMinBlockHeader(minBlockToTdengine.Header, GlobalLedgerArray[i], BLOCK_TYPE_MIN, int64(len(receiptTimeStampLedgerTypeKeyIds)), int64(currentDataSize))
			minBlockToTdengine.DataReceipts = dataReceipts
			//写入文件
			//yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + strconv.Itoa((indexMinInt-1+1440)%1440)
			utils.WriteReblocktominfile(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((indexMinInt-1+1440)%1440), minBlockToTdengine)
			log.Info("正在写入文件")
			//只存区块头
			minBlockToTdengineByteArray, _ := json.Marshal(minBlockToTdengine.Header)
			//minBlockToTdengineByteArray, _ := json.Marshal(minBlockToTdengine)

			//存到etcd
			utils.PutData(clientDelayTenMin, preKeyString, string(minBlockToTdengineByteArray), RequestTimeout)
			//打包成功的数据量
			//utils.UpdateBlockDataCountsKey(clientDelayTenMin, GlobalLedgerArray[i], BLOCK_TYPE_MIN, len(receiptTimeStampLedgerTypeKeyIds), RequestTimeout)
			utils.UpdateBlockDataCountsKey(clientDelayTenMin, GlobalLedgerArray[i], BLOCK_TYPE_MIN, blockDatacounts, RequestTimeout)

			log.Infof("%s账本存证数据分钟块排序打包成功！", GlobalLedgerArray[i])
			log.Info("存到etcd：key = ", preKeyString)
			log.Info("存到etcd：val = ", string(minBlockToTdengineByteArray))

			//存到Tdengine
			////utils.PutMinBlockToTdengine(minBlockToTdengine, GlobalLedgerArray[i])

		} else { //交易

			//初始化分钟块
			//参数分别是 账本类型 链类型 块高度 key
			var minBlockToTdengine = pb.MinuteTxBlock{}
			transactions := minBlockToTdengine.Transactions
			timeCorrect := utils.GetNowOneMinTimeStamp()
			blockHeader.UpdateTimestamp = timeCorrect

			if prePreMinuteBlock.Header != nil {
				blockHeader.PreBlockHash = prePreMinuteBlock.Header.BlockHash
			} else {
				blockHeader.PreBlockHash = "default"
			}

			minBlockToTdengine.Header = &blockHeader
			transactionTimeStampLedgerTypeTxIds := MDData[preKeyMDString]

			//排序
			sort.Strings(transactionTimeStampLedgerTypeTxIds)

			temp := ""
			num := 1
			//当前数据量大小
			var currentDataSize int
			var blockDatacounts int
			for j := 0; j < len(transactionTimeStampLedgerTypeTxIds); j++ {
				//根据时间戳 获取原来在etcd中的块数据
				//原有的从数据库中获取具体数据

				var transaction pb.Transaction
				TransactionDatamu.RLock()
				getResponse := TransactionData[transactionTimeStampLedgerTypeTxIds[j]]
				TransactionDatamu.RUnlock()
				err := json.Unmarshal([]byte(getResponse), &transaction)
				if err != nil {
					log.Error("pb.Transaction Unmarshal err", err)
				}
				// getResponse := utils.GetData(clientDelayTenMin, transactionTimeStampLedgerTypeTxIds[j], RequestTimeout)
				// for _, ev := range getResponse.Kvs {
				// 	err := json.Unmarshal(ev.Value, &transaction)
				// 	if err != nil {
				// 		log.Error("pb.Transaction Unmarshal err", err)
				// 	}
				// }

				timeStamp := strings.Split(transactionTimeStampLedgerTypeTxIds[j], TIMESTAMP_KEYID)[0]
				//将同一时间戳的数据排序
				if timeStamp != temp {
					num = 1
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					transaction.CreateTimestamp = ans
					temp = timeStamp
				} else {
					num++
					ms := strings.Split(timeStamp, ".")[1]
					strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
					curIndexStr := fmt.Sprintf(strFormat, num)
					ans := timeStamp + curIndexStr
					transaction.CreateTimestamp = ans
				}
				transactions = append(transactions, &transaction)
				transactionByteArray, _ := json.Marshal(transaction)
				currentDataSize += len(string(transactionByteArray))
				blockDatacounts = blockDatacounts + 1
			}

			//填充剩余字段
			//参数分别是 结构体 目前数据记录量 目前数据量 transactions结构体
			minBlockToTdengine.Header = utils.FillTdengineMinBlockHeader(minBlockToTdengine.Header, GlobalLedgerArray[i], BLOCK_TYPE_MIN, int64(len(transactionTimeStampLedgerTypeTxIds)),
				int64(currentDataSize))
			minBlockToTdengine.Transactions = transactions
			//只存区块头
			minBlockToTdengineByteArray, _ := json.Marshal(minBlockToTdengine.Header)
			//minBlockToTdengineByteArray, _ := json.Marshal(minBlockToTdengine)
			//写入文件
			//yearMonthDay + KeySplit + GlobalLedgerArray[i] + KeySplit + strconv.Itoa((indexMinInt-1+1440)%1440)
			utils.WriteTxblocktominfile(yearMonthDay, GlobalLedgerArray[i], strconv.Itoa((indexMinInt-1+1440)%1440), minBlockToTdengine)

			//存到etcd
			utils.PutData(clientDelayTenMin, preKeyString, string(minBlockToTdengineByteArray), RequestTimeout)
			//打包成功的数据量
			//utils.UpdateBlockDataCountsKey(clientDelayTenMin, GlobalLedgerArray[i], BLOCK_TYPE_MIN, len(transactionTimeStampLedgerTypeTxIds), RequestTimeout)
			utils.UpdateBlockDataCountsKey(clientDelayTenMin, GlobalLedgerArray[i], BLOCK_TYPE_MIN, blockDatacounts, RequestTimeout)

			log.Infof("%s账本交易数据分钟块排序打包成功！", GlobalLedgerArray[i])
			log.Info("存到etcd：key = ", preKeyString)
			log.Info("存到etcd：val = ", string(minBlockToTdengineByteArray))
			//存到Tdengine
			////utils.PutMinBlockToTdengine(minBlockToTdengine, GlobalLedgerArray[i])
		}
	} //for
}
