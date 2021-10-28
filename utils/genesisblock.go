package utils

import (
	"encoding/json"
	"fmt"
	pb "hraft/rpc"
	"time"
)

//若创始块不存在 则新建
func CreateGenesisBlock(key string, ledgerType string, blockChainType string,preDataCounts int32,preDataSize int64,preGenesisBlockHash string) (block pb.GenesisBlock) {
	var naosecond = time.Now().Nanosecond() / 1e6
	timeFormatString := time.Now().Format("2006-01-02 15:04:05")
	timeCorrect := fmt.Sprintf("%s.%d", timeFormatString, naosecond)
	block.CreateTimestamp = timeCorrect
	block.KeyId = key
	block.Height = 0
	block.DataCounts = int32(preDataCounts)
	block.DataSize = int64(preDataSize)
	block.ChildBlockCount = 0
	block.CumulativeBlock = 0
	block.Version = "v1.0"
	block.BlockChainType = blockChainType
	block.CumulativeValue = 0
	block.CumulativeNode = 0
	block.CumulativeUser = 0
	block.LedgerType = ledgerType
	block.UpdateTimestamp = timeCorrect
	block.GroupMasterNodeCount = 1
	block.GroupSlaveNodeCount = 2
	block.CreateChainTimestamp = timeCorrect
	//获取区块哈希值
	blockByteArray, _ := json.Marshal(block)
	block.GenesisBlockHash = GetStringMD5(string(blockByteArray))
	return
}
