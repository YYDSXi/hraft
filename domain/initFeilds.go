package domain

import (
	"hraft/util/log"
	"hraft/utils"
	"os"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	BLOCK_TYPE_MIN      = "MINUTE"
	BLOCK_TYPE_TENMINUT = "TENMINUTE"
	BLOCK_TYPE_DAY      = "DAY"
	ErrMsg              = "未知错误"
	SuccessCode         = int32(200)
	ErrCode             = int32(400)
	//用于拼接各字段值 之间分割符
	KeySplit = ":"
	//用于分割 时间戳 和 KeyId 或者 TransactionId
	TIMESTAMP_KEYID = "#"

	LEDGER_TYPE_VIDEO         = "video"
	LEDGER_TYPE_USER_BEHAVIOR = "user_behaviour"

	LEDGER_TYPE_NODE_CREDIBLE  = "node_credible"
	LEDGER_TYPE_SENSOR         = "sensor"
	LEDGER_TYPE_SERVICE_ACCESS = "service_access"

	ALL_LEDGER_TYPE_ARRAY = []string{LEDGER_TYPE_VIDEO, LEDGER_TYPE_USER_BEHAVIOR, LEDGER_TYPE_NODE_CREDIBLE, LEDGER_TYPE_SENSOR, LEDGER_TYPE_SERVICE_ACCESS}
	BLOCK_TYPE_ARRAY      = []string{BLOCK_TYPE_MIN, BLOCK_TYPE_TENMINUT, BLOCK_TYPE_DAY}

	//现在该结点启动账本数组
	GlobalLedgerArray = []string{}

	//2021-04-20 17:18:10.123 # 账本类型 # 存证/交易ID

	//{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
	//存储延时来的Video账本 时间戳 2021-04-20 17:18:10.123 # 账本类型 # 存证ID
	DelayLedgerVideo utils.List = utils.NewArrayList()
	//存储延时来的UserBehavior账本时间戳 2021-04-20 17:18:10.123 # 账本类型 # 存证ID
	DelayLedgerUserBehavior utils.List = utils.NewArrayList()
	//存储延时来的NodeCredible账本时间戳 2021-04-20 17:18:10.123 # 账本类型 # 交易ID
	DelayLedgerNodeCredible utils.List = utils.NewArrayList()
	//存储延时来的Sensor账本时间戳 2021-04-20 17:18:10.123 # 账本类型 # 交易ID
	DelayLedgerSensor utils.List = utils.NewArrayList()
	//存储延时来的ServiceAccess账本时间戳 2021-04-20 17:18:10.123 # 账本类型 # 交易ID
	DelayLedgerServiceAccess utils.List = utils.NewArrayList()

	//定义存储数据的map结构，数据不将存储于数据库
	TransactionData = make(map[string]string)
	ReceiptData     = make(map[string]string)
	MDData          = make(map[string][]string)
	//ReceiptMDData     = make(map[string][]string)

	TenMinBlockChangeVideo         = make([]bool, 144)
	TenMinBlockChangeUserBehavior  = make([]bool, 144)
	TenMinBlockChangeNodeCredible  = make([]bool, 144)
	TenMinBlockChangeSensor        = make([]bool, 144)
	TenMinBlockChangeServiceAccess = make([]bool, 144)

	DailyChangeVideo         = false
	DailyChangeUserBehavior  = false
	DailyChangeNodeCredible  = false
	DailyChangeSensor        = false
	DailyChangeServiceAccess = false
)

var DialTimeout time.Duration
var RequestTimeout time.Duration
var Port string

var GlobalLeaderId uint64
var GlobalLeaderName string

func InitFeilds(cli *clientv3.Client, leaderId uint64) {
	GlobalLeaderId = leaderId

	DialTimeout = time.Duration(utils.Conf.Consensus.CommonConfig.Timeout) * time.Second
	RequestTimeout = time.Duration(utils.Conf.Consensus.CommonConfig.Timeout) * time.Second

	Port = strings.Split(utils.Conf.Consensus.EtcdGroup[GlobalLeaderName].HraftAddress, ":")[1]
	// if ledgerType == LEDGER_TYPE_VIDEO {
	// 	Port = strings.Split(utils.Conf.Consensus.EtcdGroup["leader_video"].GrpcAddress, ":")[1]
	// } else if ledgerType == LEDGER_TYPE_USER_BEBAVIOR {
	// 	Port = strings.Split(utils.Conf.Consensus.EtcdGroup["leader_user"].GrpcAddress, ":")[1]
	// } else if ledgerType == LEDGER_TYPE_NODE_CREDIBLE {
	// 	Port = strings.Split(utils.Conf.Consensus.EtcdGroup["leader_node"].GrpcAddress, ":")[1]
	// } else if ledgerType == LEDGER_TYPE_SENSOR {
	// 	Port = strings.Split(utils.Conf.Consensus.EtcdGroup["leader_sensor"].GrpcAddress, ":")[1]
	// } else if ledgerType == LEDGER_TYPE_SERVICE_ACCESS {
	// 	Port = strings.Split(utils.Conf.Consensus.EtcdGroup["leader_service"].GrpcAddress, ":")[1]
	// } else {
	// 	Port = "-1"
	// }
	Init()
	for i := 0; i < 11; i++ {
		array := make([]string, 0)
		DelayLedgerVideo.Append(array)
		DelayLedgerUserBehavior.Append(array)
		DelayLedgerNodeCredible.Append(array)
		DelayLedgerSensor.Append(array)
		DelayLedgerServiceAccess.Append(array)
	}

	//全局变量清零
	for i := 0; i < len(ALL_LEDGER_TYPE_ARRAY); i++ {
		//数据条目字段清零 最后一个参数表示 是清零操作还是 叠加操作
		utils.StatisticalAllDataCounts(cli, ALL_LEDGER_TYPE_ARRAY[i], 0, RequestTimeout, true)
		//数据大小字段清零 最后一个参数表示 是清零操作还是 叠加操作
		utils.StatisticalAllDataSize(cli, ALL_LEDGER_TYPE_ARRAY[i], 0, RequestTimeout, true)
		//当天延时记录数，初始化参数时，先清零
		utils.StatisticalCurDayDelayData(cli, ALL_LEDGER_TYPE_ARRAY[i], 0, RequestTimeout, true)
		for j := 0; j < len(BLOCK_TYPE_ARRAY); j++ {
			utils.SetCurrentDayDataCounts(cli, ALL_LEDGER_TYPE_ARRAY[i], BLOCK_TYPE_ARRAY[j], 0, RequestTimeout)
			utils.SetCurrentDayDataSize(cli, ALL_LEDGER_TYPE_ARRAY[i], BLOCK_TYPE_ARRAY[j], 0, RequestTimeout)
		}
	}

	//初始化变量  所有key数量
	utils.PutData(cli, "AllKeysCounts", "0", RequestTimeout)
}

func Init() { //这里初始化函数，每个人可能存在差异
	log.Init() //设置调用上面的Init方法初始化
	if !utils.Conf.Common.LogConfig.OutputFile {
		logger.SetOutput(os.Stdout)
	}
}
