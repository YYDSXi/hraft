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

func WriteReblocktominfile(time string, ledger string, index string, minReblock pb.MinuteDataBlock) {

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/MINUTE"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(minReblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
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

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/MINUTE"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(minTxblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}

func WriteReblocktoTenminfile(time string, ledger string, index string, tenminReblock pb.TenMinuteDataBlock) {

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/TENMINUTE"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(tenminReblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}

func WriteTxblocktoTenminfile(time string, ledger string, index string, tenminTxblock pb.TenMinuteTxBlock) {

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/TENMINUTE"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(tenminTxblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}

func WriteReblocktoDayfile(time string, ledger string, index string, dayReblock pb.DailyDataBlock) {

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/DAY"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(dayReblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}

func WriteTxblocktoDayfile(time string, ledger string, index string, dayTxblock pb.DailyTxBlock) {

	var (
		filedirectory = "/home/wsq/hraft/scope/" + time + "/" + ledger + "/DAY"
		fileName      = filedirectory + "/" + index
	)
	err := os.MkdirAll(filedirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fileContent, err := json.Marshal(dayTxblock)
	fmt.Println("write file success =")
	if err = ioutil.WriteFile(fileName, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
}
