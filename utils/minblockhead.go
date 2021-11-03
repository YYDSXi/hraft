package utils

import (
	pb "hraft/rpc"
	"strconv"
)

func GetMinBlockHead(keyString string, LEDGER_TYPE string, BLOCK_TYPE string, blockHeight string) pb.BlockHeader {
	var minuteBlockHead pb.BlockHeader

	timeCorrect := GetNowOneMinTimeStamp()

	minuteBlockHead.CreateTimestamp = timeCorrect
	minuteBlockHead.KeyId = keyString
	blockHeightInt64, _ := strconv.ParseInt(blockHeight, 10, 64)
	minuteBlockHead.BlockHeight = blockHeightInt64
	minuteBlockHead.DataType = "default"
	minuteBlockHead.DataValue = "default"
	minuteBlockHead.UpdateTimestamp = timeCorrect

	var nonceInt32 int32 = 100
	minuteBlockHead.Nonce = nonceInt32
	minuteBlockHead.Target = nonceInt32
	minuteBlockHead.CurrentDataCount = 0
	minuteBlockHead.CurrentDataSize = 0
	minuteBlockHead.Version = "v1.0"
	minuteBlockHead.BlockType = BLOCK_TYPE
	minuteBlockHead.LedgerType = LEDGER_TYPE

	//transaction := make([]pb.Transaction, 0)
	//dataReceipt := make([]pb.DataReceipt, 0)
	//获取区块哈希值
	//minuteBlockHeadByteArray, _ := json.Marshal(minuteBlockHead)
	//
	//minuteBlockHeadMD5 := GetStringMD5(string(minuteBlockHeadByteArray))
	//minuteBlockHead.DataHash = minuteBlockHeadMD5
	//minuteBlockHead.BlockHash = minuteBlockHeadMD5
	//minuteBlockHead.PreBlockHash = minuteBlockHeadMD5
	return minuteBlockHead
}
