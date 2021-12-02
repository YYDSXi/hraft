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

func WriteReblocktominfile(time string, ledger string, index string, minReblock pb.MinuteDataBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {

		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {

		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(minReblock)

	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {

		log.Error("Writefile Error = ", err)
		return
	}

	//blockbody []*pb.DataReceipt
	//pb "hraft/rpc"
	//读取文件
	// fileContent, err = ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println("Read file err =", err)
	// 	return
	// }
	// var stuRes Student
	// if err := json.Unmarshal(fileContent, &stuRes); err != nil {
	// 	fmt.Println("Read file error =", err)
	// } else {
	// 	fmt.Println("Read file success =", stuRes)
	// }
}

func WriteTxblocktominfile(time string, ledger string, index string, minTxblock pb.MinuteTxBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/MINUTE"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {
		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(minTxblock)
	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		log.Error("Writefile Error = ", err)
		return
	}
}

func WriteReblocktoTenminfile(time string, ledger string, index string, tenminReblock pb.TenMinuteDataBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {
		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(tenminReblock)
	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		log.Error("Writefile Error = ", err)
		return
	}
}

func WriteTxblocktoTenminfile(time string, ledger string, index string, tenminTxblock pb.TenMinuteTxBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/TENMINUTE"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {
		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(tenminTxblock)
	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		log.Error("Writefile Error = ", err)
		return
	}
}

func WriteReblocktoDayfile(time string, ledger string, index string, dayReblock pb.DailyDataBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/DAY"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {
		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(dayReblock)
	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		log.Error("Writefile Error = ", err)
		return
	}
}

func WriteTxblocktoDayfile(time string, ledger string, index string, dayTxblock pb.DailyTxBlock) {

	mindirectory, err := os.Getwd()

	if err != nil {
		log.Error("创建区块文件主目录失败: ", err)
	}
	var (
		filedirectory = mindirectory + "/scope/" + time + "/" + ledger + "/DAY"
		fileName      = filedirectory + "/" + index
	)
	err1 := os.MkdirAll(filedirectory, os.ModePerm)
	if err1 != nil {
		log.Error("创建区块文件目录失败: ", err1)
	}

	fileContent, err := json.Marshal(dayTxblock)
	log.Info("write file success")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		log.Error("Writefile Error = ", err)
		return
	}
}
