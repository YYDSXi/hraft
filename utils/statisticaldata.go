package utils

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"time"
)

var(
	AllDataCountsKey = "AllDataCountsKey"
	AllDataSizeKey = "AllDataSizeKey"
	AllCurDayDelayData = "AllCurDayDelayData"
)
func StatisticalAllDataCounts(cli *clientv3.Client,ledgerType string,curDataCounts int,requestTimeout time.Duration,isResetZero bool) {
	//账本类型 : 字段值
	key := ledgerType+":"+AllDataCountsKey
	//如果是清零操作
	if isResetZero {
		PutData(cli,key,"0",requestTimeout)
	//如果不是清零操作
	}else {
		getResponse := GetData(cli, key, requestTimeout)
		var allDataCountsString string
		for _, ev := range getResponse.Kvs {
			allDataCountsString = string(ev.Value[:])
		}
		allDayDataCountsInt := 0

		if allDataCountsString != ""{
			allDayDataCountsInt, _ = strconv.Atoi(allDataCountsString)
		}
		allDayDataCountsInt += curDataCounts
		PutData(cli,key,strconv.Itoa(allDayDataCountsInt),requestTimeout)
	}
}
func StatisticalAllDataSize(cli *clientv3.Client,ledgerType string,curDataSize int,requestTimeout time.Duration,isResetZero bool) {
	//账本类型 : 字段值
	key := ledgerType+":"+AllDataSizeKey

	if isResetZero {
		PutData(cli,key,"0",requestTimeout)
	}else{
		getResponse := GetData(cli, key, requestTimeout)
		var allDataSizeString string
		for _, ev := range getResponse.Kvs {
			allDataSizeString = string(ev.Value[:])
		}
		allDayDataSizeInt := 0
		if allDataSizeString != ""{
			allDayDataSizeInt, _ = strconv.Atoi(allDataSizeString)
		}
		allDayDataSizeInt += curDataSize
		PutData(cli,key,strconv.Itoa(allDayDataSizeInt),requestTimeout)
	}
}


func StatisticalCurDayDelayData(cli *clientv3.Client,ledgerType string,curDataSize int,requestTimeout time.Duration,isResetZero bool) {
	//账本类型 : 字段值
	key := ledgerType+":"+AllCurDayDelayData

	if isResetZero {
		PutData(cli,key,"0",requestTimeout)
	}else{
		getResponse := GetData(cli, key, requestTimeout)
		var allDelayDataString string
		for _, ev := range getResponse.Kvs {
			allDelayDataString = string(ev.Value[:])
		}
		allDelayDataInt := 0
		if allDelayDataString != ""{
			allDelayDataInt, _ = strconv.Atoi(allDelayDataString)
		}
		allDelayDataInt += curDataSize
		PutData(cli,key,strconv.Itoa(allDelayDataInt),requestTimeout)
	}
}