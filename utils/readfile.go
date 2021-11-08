package utils

import (
	"encoding/json"
	"fmt"
	pb "hraft/rpc"
	"io/ioutil"
	"os"
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
		fmt.Println(err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		fmt.Println("Read file err =", err)

	}
	var minBlock pb.MinuteTxBlock
	if err := json.Unmarshal(fileContent, &minBlock); err != nil {
		fmt.Println("反解析 file error =", err)
	}
	return minBlock
}

//从存证分钟块文件中读取数据到增强块
func ReadReMinFiletoTenmin(time string, ledger string, index string) pb.MinuteDataBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		fmt.Println("Read file err =", err)

	}
	var minBlock pb.MinuteDataBlock
	if err := json.Unmarshal(fileContent, &minBlock); err != nil {
		fmt.Println("反解析 file error =", err)
	}
	return minBlock
}

//读取交易增强块文件到天块
func ReadTxTenMinFiletoDay(time string, ledger string, index string) pb.TenMinuteTxBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		fmt.Println("Read file err =", err)

	}
	var tenMinuteBlock pb.TenMinuteTxBlock
	if err := json.Unmarshal(fileContent, &tenMinuteBlock); err != nil {
		fmt.Println("反解析 file error =", err)
	}
	return tenMinuteBlock
}

//读取存证增强块文件到天块
func ReadReTenMinFiletoDay(time string, ledger string, index string) pb.TenMinuteDataBlock {

	mindirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	var (
		fileName = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE" + "/" + index
	)

	//读取文件
	fileContent, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		fmt.Println("Read file err =", err)

	}
	var tenMinuteBlock pb.TenMinuteDataBlock
	if err := json.Unmarshal(fileContent, &tenMinuteBlock); err != nil {
		fmt.Println("反解析 file error =", err)
	}
	return tenMinuteBlock
}
