package domain

import (
	"github.com/robfig/cron"
	clientv3 "go.etcd.io/etcd/client/v3"
	"hraft/utils"
)

var (

	//创建创始块，每天的0点 0分 0秒启动任务
	createGenesisBlockSpec = "0 0 0 * * ?"

	//创建区块头，生成1440个，每天的0点 0分 0秒启动任务，第0s的时候
	createBlockHeadSpec = "0 0 0 * * ?"

	//构建分钟块，每天的每点 每分 30秒启动任务，第30s的时候
	createMinBlockSpec = "30 * * * * ?"

	//构建增强块，每天的每点 00,10,20,30,40,50分 40秒启动任务，第40s的时候
	createTenMinBlockSpec = "40 00,10,20,30,40,50 * * * ?"
	//createTenMinBlockSpec = "40 * * * * ?"

	//构建天块，每天的0点 0分 50秒启动任务，第50s的时候
	createDailyBlockSpec = "50 0 0 * * ?"

	// blockHeadDisTime = []int{utils.Conf.Consensus.AsyncTask[LEDGER_TYPE_VIDEO].BlockHeader.Interval,
	// 	utils.Conf.Consensus.AsyncTask[LEDGER_TYPE_USER_BEHAVIOR].BlockHeader.Interval,
	// 	utils.Conf.Consensus.AsyncTask[LEDGER_TYPE_NODE_CREDIBLE].BlockHeader.Interval,
	// 	utils.Conf.Consensus.AsyncTask[LEDGER_TYPE_SENSOR].BlockHeader.Interval,
	// 	utils.Conf.Consensus.AsyncTask[LEDGER_TYPE_SERVICE_ACCESS].BlockHeader.Interval}

	//更新一些全局变量，每天的0点 0分 0秒启动任务
	updateGlobalFieldSpec = "0 0 0 * * ?"
)

func AutoWorkMain(client *clientv3.Client) {
	//开启服务时 首先创建一次
	go AutoCreateGenesisBlockToEtcd(client)
	go AutoCreateBlockHeadToEtcd(client)

	c := cron.New()
	c.AddFunc(createGenesisBlockSpec, func() {
		//创建创始块
		go AutoCreateGenesisBlockToEtcd(client)
	})
	c.AddFunc(createBlockHeadSpec, func() {
		//创建区块头
		go AutoCreateBlockHeadToEtcd(client)
	})
	c.AddFunc(createMinBlockSpec, func() {
		//创建分钟块
		go AutoCreateMinBlockToEtcd(client)
	})

	c.AddFunc(createTenMinBlockSpec, func() {
		//创建增强块
		go AutoCreateTenMinBlockToEtcd(client)
	})

	c.AddFunc(createDailyBlockSpec, func() {
		//创建天块
		go AutoCreateDailyBlockToEtcd(client)
	})

	c.AddFunc(updateGlobalFieldSpec, func() {
		//更新一些全局统计变量
		go func() {
			//全局变量清零
			for i:=0;i<len(ALL_LEDGER_TYPE_ARRAY);i++{
				//数据条目字段清零 最后一个参数表示 是清零操作还是 叠加操作
				utils.StatisticalAllDataCounts(client,ALL_LEDGER_TYPE_ARRAY[i],0,RequestTimeout,true)
				//数据大小字段清零 最后一个参数表示 是清零操作还是 叠加操作
				utils.StatisticalAllDataSize(client,ALL_LEDGER_TYPE_ARRAY[i],0,RequestTimeout,true)
				//当天延时记录数，初始化参数时，先清零
				utils.StatisticalCurDayDelayData(client,ALL_LEDGER_TYPE_ARRAY[i],0,RequestTimeout,true)
				for j:=0;j<len(BLOCK_TYPE_ARRAY);j++{
					utils.SetCurrentDayDataCounts(client,ALL_LEDGER_TYPE_ARRAY[i],BLOCK_TYPE_ARRAY[j],0,RequestTimeout)
					utils.SetCurrentDayDataSize(client,ALL_LEDGER_TYPE_ARRAY[i],BLOCK_TYPE_ARRAY[j],0,RequestTimeout)
				}
			}
		}()
	})
	c.Start()

	select {}
}
