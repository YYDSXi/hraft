package main

import (
	"context"
	"fmt"
	"hraft/domain"
	"hraft/serviceregister"
	"hraft/utils"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	DialTimeout          = 5 * time.Second
	Endpoints            = []string{"127.0.0.1:2379"}
	ConfPath             = "./config.yaml"
	NowLeaderStartLedger = []string{}
)

var isHaveLeader = false
var log = logrus.New()

func main() {

	//加载配置文件
	utils.InitAndGetConf(ConfPath)
	log.Info("配置文件加载完成...")

	//判断输入节点名称合法性 该结点启动那些账本
	InitLedgerAndLeader()

	//判断是否是主节点 主节点启动任务
	for _, ep := range Endpoints {

		client, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{ep},
			DialTimeout: DialTimeout,
		})
		if err != nil {
			log.Error("main-main()获取客户端错误", err)
			return
		}
		defer client.Close()

		resp, err := client.Status(context.Background(), ep)
		if err != nil {
			log.Error("main-main()获取客户端状态错误", err)
			return
		}
		if resp.Header.MemberId == resp.Leader || len(Endpoints) == 1 {

			isHaveLeader = true
			//初始化参数 uint64
			domain.InitFeilds(client, resp.Leader)

			if domain.Port == "-1" {
				log.Error("不存在该账本...")
				return
			}

			//注册服务
			go serviceregister.RegisterService(client, domain.GlobalLeaderName)

			//开启redis接受数据服务
			go domain.RedisServerMain(client)

			//开启自动任务服务
			go domain.AutoWorkMain(client)

		}
	}
	if isHaveLeader == false {
		log.Error("无主节点!")
		return
	}
	select {}
}

//判断输入节点名称合法性 该结点启动那些账本
func InitLedgerAndLeader() {

	//参数
	var allArgs, dis string
	for i := 1; i < len(os.Args); i++ {
		allArgs += dis + os.Args[i]
		dis = "#"
	}
	if len(os.Args) < 2 {
		fmt.Println("请输入命令名称,如:\n[leader/status]")
		return
	}

	if strings.Split(allArgs, "#")[0] == "status" {
		utils.HraftStatus()
		time.Sleep(4 * 1000)
		os.Exit(1)
	}

	//账本类型
	InputLeaderName := strings.Split(allArgs, "#")[0]

	//判断所输入的节点名称在配置文件中是否能找到
	if InputLeaderName == utils.Conf.Common.LedgerName[domain.LEDGER_TYPE_VIDEO].Leader {
		domain.GlobalLedgerArray = append(domain.GlobalLedgerArray, domain.LEDGER_TYPE_VIDEO)
	}
	if InputLeaderName == utils.Conf.Common.LedgerName[domain.LEDGER_TYPE_USER_BEHAVIOR].Leader {
		domain.GlobalLedgerArray = append(domain.GlobalLedgerArray, domain.LEDGER_TYPE_USER_BEHAVIOR)
	}
	if InputLeaderName == utils.Conf.Common.LedgerName[domain.LEDGER_TYPE_NODE_CREDIBLE].Leader {
		domain.GlobalLedgerArray = append(domain.GlobalLedgerArray, domain.LEDGER_TYPE_NODE_CREDIBLE)
	}
	if InputLeaderName == utils.Conf.Common.LedgerName[domain.LEDGER_TYPE_SENSOR].Leader {
		domain.GlobalLedgerArray = append(domain.GlobalLedgerArray, domain.LEDGER_TYPE_SENSOR)
	}
	if InputLeaderName == utils.Conf.Common.LedgerName[domain.LEDGER_TYPE_SERVICE_ACCESS].Leader {
		domain.GlobalLedgerArray = append(domain.GlobalLedgerArray, domain.LEDGER_TYPE_SERVICE_ACCESS)
	}

	if len(domain.GlobalLedgerArray) < 1 {
		fmt.Println("节点名称不匹配！文件中不存在该结点")
		fmt.Println("请输入启动节点名称,如:\n[leader]")
		return
	}

	//全局变量赋值
	domain.GlobalLeaderName = InputLeaderName

	log.Infof("%s节点名称启动账本名称为%s", domain.GlobalLeaderName, domain.GlobalLedgerArray)
}
