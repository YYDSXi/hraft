package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	pb "hraft/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:8880"
	LEDGER_TYPE_VIDEO         = "video"
	LEDGER_TYPE_USER_BEHAVIOR = "user_behaviour"

	LEDGER_TYPE_NODE_CREDIBLE  = "node_credible"
	LEDGER_TYPE_SENSOR         = "sensor"
	LEDGER_TYPE_SERVICE_ACCESS = "service_access"
)

func main()  {

	//初始化rpc连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewToUpperClient(conn)


	//参数
	var allArgs, dis string
	for i := 1; i < len(os.Args); i++ {
		allArgs += dis + os.Args[i]
		dis = "#"
	}
	if len(os.Args) < 2 {
		fmt.Println("请输入数据！")
		return
	}
	allData := strings.Split(allArgs, "#")

	for i:=0;i<len(allData);i++{
		fmt.Println("参数是:",allData[i])
	}
	fmt.Println("=================")
	switch allData[0] {
	//=================1 存证账本类型==================
		case LEDGER_TYPE_VIDEO:
			dataReceipts := string2Receipt(allData)
			var videoData pb.VideoData
			for i := 0; i < len(dataReceipts); i++ {
				videoData.DataReceipts = append(videoData.DataReceipts, &dataReceipts[i])
			}
			r, err := c.Video(context.Background(), &pb.VideoData{DataReceipts: videoData.DataReceipts})
			if err != nil {
				log.Fatalf("cloud not greet: %v", err)
			}
			fmt.Println("Response.ErrCode:", r.ErrCode)
			fmt.Println("Response.ErrMsg:", r.ErrMsg)
	//=================2 存证账本类型==================
		case LEDGER_TYPE_USER_BEHAVIOR:
			dataReceipts := string2Receipt(allData)
			var userBehaviourData pb.UserBehaviourData
			for i := 0; i < len(dataReceipts); i++ {
				userBehaviourData.DataReceipts = append(userBehaviourData.DataReceipts, &dataReceipts[i])
			}
			r, err := c.UserBehaviour(context.Background(), &pb.UserBehaviourData{DataReceipts: userBehaviourData.DataReceipts})
			if err != nil {
				log.Fatalf("cloud not greet: %v", err)
			}
			fmt.Println("Response.ErrCode:", r.ErrCode)
			fmt.Println("Response.ErrMsg:", r.ErrMsg)
	//=================3 交易账本类型==================
		case LEDGER_TYPE_NODE_CREDIBLE:
			transactions := string2Transaction(allData)
			var nodeCredibleData pb.NodeCredibleData
			for i := 0; i < len(transactions); i++ {
				nodeCredibleData.Transactions = append(nodeCredibleData.Transactions, &transactions[i])
			}
			r, err := c.NodeCredible(context.Background(), &pb.NodeCredibleData{Transactions: nodeCredibleData.Transactions})
			if err != nil {
				log.Fatalf("cloud not greet: %v", err)
			}
			fmt.Println("Response.ErrCode:", r.ErrCode)
			fmt.Println("Response.ErrMsg:", r.ErrMsg)
	//=================4 交易账本类型==================
		case LEDGER_TYPE_SENSOR:
			transactions := string2Transaction(allData)
			var sensorData pb.SensorData
			for i := 0; i < len(transactions); i++ {
				sensorData.Transactions = append(sensorData.Transactions, &transactions[i])
			}
			r, err := c.Sensor(context.Background(), &pb.SensorData{Transactions: sensorData.Transactions})
			if err != nil {
				log.Fatalf("cloud not greet: %v", err)
			}
			fmt.Println("Response.ErrCode:", r.ErrCode)
			fmt.Println("Response.ErrMsg:", r.ErrMsg)
	//=================5 交易账本类型==================
		case LEDGER_TYPE_SERVICE_ACCESS:
			transactions := string2Transaction(allData)
			var serviceAccessData pb.ServiceAccessData
			for i := 0; i < len(transactions); i++ {
				serviceAccessData.Transactions = append(serviceAccessData.Transactions, &transactions[i])
			}
			r, err := c.ServiceAccess(context.Background(), &pb.ServiceAccessData{Transactions: serviceAccessData.Transactions})
			if err != nil {
				log.Fatalf("cloud not greet: %v", err)
			}
			fmt.Println("Response.ErrCode:", r.ErrCode)
			fmt.Println("Response.ErrMsg:", r.ErrMsg)

		default:
			fmt.Println("输入账本类型错误！")
	}
}

func string2Receipt(str []string) []pb.DataReceipt {
	var dataReceipts []pb.DataReceipt
	for i:=1;i<len(str);i++ {
		var dataReceipt pb.DataReceipt
		json.Unmarshal([]byte(str[i]),&dataReceipt)
		dataReceipts=append(dataReceipts,dataReceipt)
	}
	return dataReceipts
}

func string2Transaction(str []string) []pb.Transaction {
	var transactions []pb.Transaction
	for i:=1;i<len(str);i++ {
		var transaction pb.Transaction
		json.Unmarshal([]byte(str[i]),&transaction)
		transactions=append(transactions,transaction)
	}
	return transactions
}
