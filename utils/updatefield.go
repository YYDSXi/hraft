package utils

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"time"
)

var (
	//当天的数据条目数 统计
	CurrentDayDataCountsKey = "CurrentDayDataCountsKey"
	//当天的数据量 统计
	CurrentDayDataSizeKey = "CurrentDayDataSizeKey"
)
func GetCurrentDayDataCounts(cli *clientv3.Client,ledgerType string,blockType string,requestTimeout time.Duration) int {
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataCountsKey
	getResponse := GetData(cli, key, requestTimeout)
	var currentDayDataCountsString string
	for _, ev := range getResponse.Kvs {
		currentDayDataCountsString = string(ev.Value[:])
	}
	currentDayDataCountsInt := 0
	if currentDayDataCountsString != ""{
		currentDayDataCountsInt, _ = strconv.Atoi(currentDayDataCountsString)
	}
	return currentDayDataCountsInt
}
func GetCurrentDayDataSize(cli *clientv3.Client,ledgerType string,blockType string,requestTimeout time.Duration) int {
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataSizeKey
	getResponse := GetData(cli, key, requestTimeout)
	var currentDayDataSizeString string
	for _, ev := range getResponse.Kvs {
		currentDayDataSizeString = string(ev.Value[:])
	}
	currentDayDataSizeInt := 0
	if currentDayDataSizeString != ""{
		currentDayDataSizeInt, _ = strconv.Atoi(currentDayDataSizeString)
	}
	return currentDayDataSizeInt
}


func SetCurrentDayDataCounts(cli *clientv3.Client,ledgerType string,blockType string,target int,requestTimeout time.Duration) {
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataCountsKey
	PutData(cli,key,targetString,requestTimeout)
}

func SetCurrentDayDataSize(cli *clientv3.Client,ledgerType string,blockType string,target int,requestTimeout time.Duration) {
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataSizeKey
	PutData(cli,key,targetString,requestTimeout)
}


func UpdateCurrentDayDataCounts(cli *clientv3.Client,ledgerType string,blockType string,origin int,requestTimeout time.Duration) {
	target := origin + GetCurrentDayDataCounts(cli,ledgerType,blockType,requestTimeout)
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataCountsKey
	PutData(cli,key,targetString,requestTimeout)
}

func UpdateCurrentDayDataSize(cli *clientv3.Client,ledgerType string,blockType string,origin int,requestTimeout time.Duration) {
	target := origin + GetCurrentDayDataSize(cli,ledgerType,blockType,requestTimeout)
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType+":"+blockType+":"+CurrentDayDataSizeKey
	PutData(cli,key,targetString,requestTimeout)
}

