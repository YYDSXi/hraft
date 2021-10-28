package utils

import (
	"encoding/json"
	pb "hraft/rpc"
)


func FillTdengineMinBlockHeader(minBlockToTdengine *pb.BlockHeader,ledgerType string,blockType string, currentDataCount int64, currentDataSize int64) *pb.BlockHeader {
	minBlockToTdengine.LedgerType = ledgerType
	minBlockToTdengine.BlockType = blockType
	minBlockToTdengine.CurrentDataCount = minBlockToTdengine.CurrentDataCount + currentDataCount
	minBlockToTdengine.CurrentDataSize = minBlockToTdengine.CurrentDataSize+ currentDataSize
	//填充
	minBlockToTdengineByteArray, _ := json.Marshal(minBlockToTdengine)
	minBlockToTdengine.DataHash = GetStringMD5(string(minBlockToTdengineByteArray))
	minBlockToTdengine.BlockHash = GetStringMD5(string(minBlockToTdengineByteArray))
	return minBlockToTdengine
}


func InitToTdengineTxTenMinBlock(LEDGER_TYPE string, BLOCK_TYPE string, prePreBlockHash string, blockHeight int, keyString string) pb.TenMinuteTxBlock {
	var tenMinBlockToTdengine pb.TenMinuteTxBlock
	//获取 时间戳
	timeCorrect := GetNowOneMinTimeStamp()
	tenMinBlockToTdengine.LedgerType = LEDGER_TYPE
	tenMinBlockToTdengine.BlockType = BLOCK_TYPE
	tenMinBlockToTdengine.CreateTimestamp = timeCorrect
	tenMinBlockToTdengine.BlockHeight = int64(blockHeight)
	tenMinBlockToTdengine.KeyId = keyString
	tenMinBlockToTdengine.PreBlockHash = prePreBlockHash
	return tenMinBlockToTdengine
}
func InitToTdengineDataTenMinBlock(LEDGER_TYPE string, BLOCK_TYPE string, prePreBlockHash string, blockHeight int, keyString string) pb.TenMinuteDataBlock {
	var tenMinBlockToTdengine pb.TenMinuteDataBlock

	//获取 时间戳
	timeCorrect := GetNowOneMinTimeStamp()
	tenMinBlockToTdengine.LedgerType = LEDGER_TYPE
	tenMinBlockToTdengine.BlockType = BLOCK_TYPE
	tenMinBlockToTdengine.CreateTimestamp = timeCorrect
	tenMinBlockToTdengine.BlockHeight = int64(blockHeight)
	tenMinBlockToTdengine.KeyId = keyString
	tenMinBlockToTdengine.PreBlockHash = prePreBlockHash
	return tenMinBlockToTdengine
}

func FillTdengineTxTenMinBlockRemainingFieldsAndDataReceipts(tenMinBlockToTdengine pb.TenMinuteTxBlock) pb.TenMinuteTxBlock {
	tenMinBlockToTdengineByteArray, _ := json.Marshal(tenMinBlockToTdengine)
	tenMinBlockToTdengine.BlockHash = GetStringMD5(string(tenMinBlockToTdengineByteArray))
	return tenMinBlockToTdengine
}
func FillTdengineDataTenMinBlockRemainingFieldsAndDataReceipts(tenMinBlockToTdengine pb.TenMinuteDataBlock) pb.TenMinuteDataBlock {
	tenMinBlockToTdengineByteArray, _ := json.Marshal(tenMinBlockToTdengine)
	tenMinBlockToTdengine.BlockHash = GetStringMD5(string(tenMinBlockToTdengineByteArray))
	return tenMinBlockToTdengine
}
func InitToTdengineTxDailyBlock(LEDGER_TYPE string, BLOCK_TYPE string, prePreBlockHash string, blockHeight int, keyString string) pb.DailyTxBlock {
	var dailyBlockToTdengine pb.DailyTxBlock
	//获取 时间戳
	timeCorrect := GetNowOneMinTimeStamp()
	dailyBlockToTdengine.LedgerType = LEDGER_TYPE
	dailyBlockToTdengine.BlockType = BLOCK_TYPE
	dailyBlockToTdengine.CreateTimestamp = timeCorrect
	dailyBlockToTdengine.BlockHeight = int64(blockHeight)
	dailyBlockToTdengine.KeyId = keyString
	dailyBlockToTdengine.PreBlockHash = prePreBlockHash
	return dailyBlockToTdengine
}
func InitToTdengineDataDailyBlock(LEDGER_TYPE string, BLOCK_TYPE string, prePreBlockHash string, blockHeight int, keyString string) pb.DailyDataBlock {
	var dailyBlockToTdengine pb.DailyDataBlock
	//获取 时间戳
	timeCorrect := GetNowOneMinTimeStamp()
	dailyBlockToTdengine.LedgerType = LEDGER_TYPE
	dailyBlockToTdengine.BlockType = BLOCK_TYPE
	dailyBlockToTdengine.CreateTimestamp = timeCorrect
	dailyBlockToTdengine.BlockHeight = int64(blockHeight)
	dailyBlockToTdengine.KeyId = keyString
	dailyBlockToTdengine.PreBlockHash = prePreBlockHash
	return dailyBlockToTdengine
}
func FillTdengineTxDailyBlockRemainingFieldsAndDataReceipts(dailyBlockToTdengine pb.DailyTxBlock) pb.DailyTxBlock {
	dailyBlockToTdengineByteArray, _ := json.Marshal(dailyBlockToTdengine)
	dailyBlockToTdengine.BlockHash = GetStringMD5(string(dailyBlockToTdengineByteArray))
	return dailyBlockToTdengine
}
func FillTdengineDataDailyBlockRemainingFieldsAndDataReceipts(dailyBlockToTdengine pb.DailyDataBlock) pb.DailyDataBlock{
	dailyBlockToTdengineByteArray, _ := json.Marshal(dailyBlockToTdengine)
	dailyBlockToTdengine.BlockHash = GetStringMD5(string(dailyBlockToTdengineByteArray))
	return dailyBlockToTdengine
}

func CreateMinBlockHeader() *pb.BlockHeader {
	var header pb.BlockHeader
	header.KeyId="1"
	return &header
}