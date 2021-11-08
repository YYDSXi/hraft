package utils

import (
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	//统计每种账本每种链实际打包的数据量
	BlockDataCountsKey = "BlockDataCountsKey"
	//统计每种账本每种链实际打包的数据大小
	BlockDataCountsSize = "BlockDataCountsSize"
)

func GetBlockDataCountsKey(cli *clientv3.Client, ledgerType string, blockType string, requestTimeout time.Duration) int {
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsKey
	getResponse := GetData(cli, key, requestTimeout)
	var BlockDataCountsString string
	for _, ev := range getResponse.Kvs {
		BlockDataCountsString = string(ev.Value[:])
	}
	BlockDataCountsInt := 0
	if BlockDataCountsString != "" {
		BlockDataCountsInt, _ = strconv.Atoi(BlockDataCountsString)
	}
	return BlockDataCountsInt
}
func GetBlockDataCountsSize(cli *clientv3.Client, ledgerType string, blockType string, requestTimeout time.Duration) int {
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsSize
	getResponse := GetData(cli, key, requestTimeout)
	var BlockDataCountsSizeString string
	for _, ev := range getResponse.Kvs {
		BlockDataCountsSizeString = string(ev.Value[:])
	}
	BlockDataCountsSizeInt := 0
	if BlockDataCountsSizeString != "" {
		BlockDataCountsSizeInt, _ = strconv.Atoi(BlockDataCountsSizeString)
	}
	return BlockDataCountsSizeInt
}

func SetBlockDataCountsKey(cli *clientv3.Client, ledgerType string, blockType string, target int, requestTimeout time.Duration) {
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsKey
	PutData(cli, key, targetString, requestTimeout)
}

func SetBlockDataCountsSize(cli *clientv3.Client, ledgerType string, blockType string, target int, requestTimeout time.Duration) {
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsSize
	PutData(cli, key, targetString, requestTimeout)
}

func UpdateBlockDataCountsKey(cli *clientv3.Client, ledgerType string, blockType string, origin int, requestTimeout time.Duration) {
	target := origin + GetCurrentDayDataCounts(cli, ledgerType, blockType, requestTimeout)
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsKey
	PutData(cli, key, targetString, requestTimeout)
}

func UpdateBlockDataCountsSize(cli *clientv3.Client, ledgerType string, blockType string, origin int, requestTimeout time.Duration) {
	target := origin + GetCurrentDayDataSize(cli, ledgerType, blockType, requestTimeout)
	targetString := strconv.Itoa(target)
	//账本类型 : 链类型 : 字段值
	key := ledgerType + ":" + blockType + ":" + BlockDataCountsSize
	PutData(cli, key, targetString, requestTimeout)
}
