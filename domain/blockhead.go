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

//自动创建区块头任务
func AutoCreateBlockHeadToEtcd(clientDelayTenMin *clientv3.Client) {
	//遍历账本
	for i := 0; i < len(GlobalLedgerArray); i++ {
		ledgerName := GlobalLedgerArray[i]
		go func() {
			disTime := utils.Conf.Consensus.AsyncTask[ledgerName].BlockHeader.Interval
			ticker := time.NewTicker(time.Second * time.Duration(disTime)) // 运行时长
			ch := make(chan string)
			go func() {
				indexMinInt := utils.GetIndexMinInt() - 10
				for indexMinInt < 1440 {
					select {
					case <-ticker.C:
						//构建key值 年月日+账本类型+分钟索引
						//年月日 账本类型 分钟索引
						blockHeadKeyString := time.Now().Format("2006-01-02") + KeySplit + ledgerName + KeySplit + strconv.Itoa(indexMinInt)

						//判断区块头是否已存在
						isExistMDResponse := utils.GetData(clientDelayTenMin, blockHeadKeyString, RequestTimeout)
						var isExistBlockHeader pb.BlockHeader
						for _, ev := range isExistMDResponse.Kvs {
							err := json.Unmarshal(ev.Value, &isExistBlockHeader)
							if err != nil {
								log.Error("preMinuteMDData Unmarshal err", err)
							}
						}

						if isExistBlockHeader.CreateTimestamp == "" {
							blockHead := utils.GetMinBlockHead(blockHeadKeyString, ledgerName, BLOCK_TYPE_MIN, strconv.Itoa(indexMinInt))

							log.Infof("%s账本 第%s个区块头初始化成功", ledgerName, strconv.Itoa(indexMinInt))

							//区块头存etcd
							blockHeadByteArray, _ := json.Marshal(blockHead)
							utils.PutData(clientDelayTenMin, blockHeadKeyString, string(blockHeadByteArray), RequestTimeout)
						}
					}
					indexMinInt++
				}
				ticker.Stop()
				ch <- "0"
			}()
			<-ch
		}()
	}
}
