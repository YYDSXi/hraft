package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "hraft/rpc"
	"hraft/utils"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address                   = "127.0.0.1:8880"
	LEDGER_TYPE_VIDEO         = "video"
	LEDGER_TYPE_USER_BEHAVIOR = "user_behaviour"

	LEDGER_TYPE_NODE_CREDIBLE  = "node_credible"
	LEDGER_TYPE_SENSOR         = "sensor"
	LEDGER_TYPE_SERVICE_ACCESS = "service_access"
)

func main() {

	//初始化rpc连接
	recvSize := 6108890
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(recvSize)))
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewToUpperClient(conn)

	//生成数据

	// var allArgs, dis string
	// for i := 1; i < len(os.Args); i++ {
	// 	allArgs += dis + os.Args[i]
	// 	dis = "#"
	// }
	// if len(os.Args) < 2 {
	// 	fmt.Println("请输入数据！")
	// 	return
	// }
	// allData := strings.Split(allArgs, "#")

	// for i := 0; i < len(allData); i++ {
	// 	fmt.Println("参数是:", allData[i])
	// }

	var datas string
	datas = os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	for i := 0; i < count; i++ {
		globalTime_Now := "\"" + utils.GetNowOneMinTimeStamp2() + "\""
		globalTime_keyid := "\"" + "00" + strconv.Itoa(i) + "\""

		data := fmt.Sprintf(`{"CreateTimestamp":%s,"EntityId":"r21r4f431feqf","KeyId":%s,"ReceiptValue":10.1,"Version":"v1.0","UserName":"ls","OperationType":"111","DataType":"test","ServiceType":"test","FileName":"测试数据1.txt","FileSize":140.20,"FileHash":"0xnofuegwuifgf932fou29f23992effe","Uri":"/usr/local/1.txt","ParentKeyId":"2r45fr","AttachmentFileUris":["fewqf","g52gtttg"],"AttachmentTotalHash":"3f41f431"}`, globalTime_Now, globalTime_keyid)
		datas += "#" + data
	}
	allData := strings.Split(datas, "#")
	// for i := 0; i < len(allData); i++ {
	// 	fmt.Println("参数是:", allData[i])
	// }

	// fmt.Println("=================")
	switch allData[0] {
	//=================1 存证账本类型==================
	case LEDGER_TYPE_VIDEO:
		dataReceipts := string2Receipt(allData)
		var videoData pb.VideoData
		for i := 0; i < len(dataReceipts); i++ {
			videoData.DataReceipts = append(videoData.DataReceipts, &dataReceipts[i])
		}
		start := time.Now().UnixNano()
		r, err := c.Video(context.Background(), &pb.VideoData{DataReceipts: videoData.DataReceipts})
		end := time.Now().UnixNano()
		datacounts := len(dataReceipts)
		fmt.Printf("发送%d条数据总用时：%v纳秒\n", datacounts, (end - start))
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
	for i := 1; i < len(str); i++ {
		var dataReceipt pb.DataReceipt
		json.Unmarshal([]byte(str[i]), &dataReceipt)
		dataReceipts = append(dataReceipts, dataReceipt)
	}
	return dataReceipts
}

func string2Transaction(str []string) []pb.Transaction {
	var transactions []pb.Transaction
	for i := 1; i < len(str); i++ {
		var transaction pb.Transaction
		json.Unmarshal([]byte(str[i]), &transaction)
		transactions = append(transactions, transaction)
	}
	return transactions
}
