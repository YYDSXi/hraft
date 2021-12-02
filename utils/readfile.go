package utils

import (
	"encoding/json"
	pb "hraft/rpc"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// type Student struct {
// 	Name  string
// 	Age   int
// 	Score float64
// }
//从交易分钟数据区块文件中读取数据到增强块
func ReadTxMinFiletoTenmin(time string, ledger string, index string) pb.MinuteTxBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {

		log.Error("Read file err = ", err1)

	}
	var minBlock pb.MinuteTxBlock
	if err := json.Unmarshal(fileContent, &minBlock); err != nil {

		log.Error("反解析 file error = ", err)
	}
	return minBlock
}

//从存证分钟块文件中读取数据到增强块
func ReadReMinFiletoTenmin(time string, ledger string, index string) pb.MinuteDataBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		log.Error("Read file err = ", err1)

	}
	var minBlock pb.MinuteDataBlock
	if err := json.Unmarshal(fileContent, &minBlock); err != nil {
		log.Error("反解析 file error = ", err)
	}
	return minBlock
}

//读取交易增强块文件到天块
func ReadTxTenMinFiletoDay(time string, ledger string, index string) pb.TenMinuteTxBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		log.Error("Read file err = ", err1)

	}
	var tenMinuteBlock pb.TenMinuteTxBlock
	if err := json.Unmarshal(fileContent, &tenMinuteBlock); err != nil {
		log.Error("反解析 file error = ", err)
	}
	return tenMinuteBlock
}

//读取存证增强块文件到天块
func ReadReTenMinFiletoDay(time string, ledger string, index string) pb.TenMinuteDataBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		log.Error("Read file err = ", err1)

	}
	var tenMinuteBlock pb.TenMinuteDataBlock
	if err := json.Unmarshal(fileContent, &tenMinuteBlock); err != nil {
		log.Error("反解析 file error = ", err)
	}
	return tenMinuteBlock
}
