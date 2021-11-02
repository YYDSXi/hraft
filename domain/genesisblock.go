package domain

import (
	"encoding/json"
	pb "hraft/rpc"
	"hraft/utils"
	"time"

	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

//创始块 自动任务
func AutoCreateGenesisBlockToEtcd(client *clientv3.Client) {

	time.Sleep(time.Duration(1) * time.Second)
	//遍历账本
	for i := 0; i < len(GlobalLedgerArray); i++ {
		//遍历三个链
		for j := 0; j < len(BLOCK_TYPE_ARRAY); j++ {

			//构建创始块的key
			//年月日 账本类型 链类型
			genesisBlockKeyString := time.Now().Format("2006-01-02") + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_ARRAY[j]

			//判断创始块是否已存在
			isExistMDResponse := utils.GetData(client, genesisBlockKeyString, RequestTimeout)
			var isExistGenesisBlock pb.GenesisBlock
			for _, ev := range isExistMDResponse.Kvs {
				err := json.Unmarshal(ev.Value, &isExistGenesisBlock)
				if err != nil {
					log.Error("反序列化结构体，pb.GenesisBlock失败：", err)
				}
			}
			if isExistGenesisBlock.CreateTimestamp == "" {
				//获取上一天的创世区块，将其相关字段取出来
				//年月日 账本类型 链类型
				preGenesisBlockKeyString := time.Now().Add(-time.Hour*24).Format("2006-01-02") + KeySplit + GlobalLedgerArray[i] + KeySplit + BLOCK_TYPE_ARRAY[j]
				//判断创始块是否已存在
				isExistPreMDResponse := utils.GetData(client, preGenesisBlockKeyString, RequestTimeout)
				var isExistPreGenesisBlock pb.GenesisBlock
				for _, ev := range isExistPreMDResponse.Kvs {
					err := json.Unmarshal(ev.Value, &isExistPreGenesisBlock)
					if err != nil {
						log.Error("反序列化结构体，pb.GenesisBlock失败：", err)
					}
				}
				//当前创世块的 数据条目数 大小  = 上一创世块字段值 + 上一天的统计量
				curDataCounts := isExistPreGenesisBlock.DataCounts + int32(utils.GetCurrentDayDataCounts(client, GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j], RequestTimeout))
				curDataSize := isExistPreGenesisBlock.DataSize + int64(utils.GetCurrentDayDataSize(client, GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j], RequestTimeout))

				//将上一天的统计量置零
				utils.SetCurrentDayDataCounts(client, GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j], 0, RequestTimeout)
				utils.SetCurrentDayDataSize(client, GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j], 0, RequestTimeout)

				preGenesisBlockHash := isExistPreGenesisBlock.GenesisBlockHash
				genesisBlock := utils.CreateGenesisBlock(genesisBlockKeyString, GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j],
					curDataCounts, curDataSize, preGenesisBlockHash)

				//创始块存etcd
				genesisBlockByteArray, _ := json.Marshal(genesisBlock)
				utils.PutData(client, genesisBlockKeyString, string(genesisBlockByteArray), RequestTimeout)

				log.Infof("%s账本,%s链类型创世块创建成功！", GlobalLedgerArray[i], BLOCK_TYPE_ARRAY[j])
				log.Info("创世块key=", genesisBlockKeyString)
				log.Info("创世块val=", string(genesisBlockByteArray))
			}

			//创始块存tdengine
			////respString := utils.PutGenesisBlockToTdengine(genesisBlock, GlobalLedgerArray[i])
			////fmt.Println("存至tdengine返回信息：", respString)
		}
	}

}
