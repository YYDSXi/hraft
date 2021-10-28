package main

import (
	"fmt"
	"log"

	pb "hraft/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:8880"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewToUpperClient(conn)


	//var data pb.VideoData
	//var dat1 pb.DataReceipt
	//var dat2 pb.DataReceipt
	//
	//dat1.EntityId = "123"
	//dat1.CreateTimestamp = "2021-06-02 15:48:20.456"
	//dat1.KeyId = "aaaaaaaaaaaa456"
	//
	//dat2.EntityId = "456"
	//dat2.CreateTimestamp = "2021-06-02 15:48:20.456"
	//dat2.KeyId = "bbbbbbbbb456"
	//
	//data.DataReceipts = append(data.DataReceipts, &dat1)
	//data.DataReceipts = append(data.DataReceipts, &dat2)
	//r, err := c.Video(context.Background(), &pb.VideoData{DataReceipts: data.DataReceipts})
	//if err != nil {
	//	log.Fatalf("cloud not greet: %v", err)
	//}
	//fmt.Println("Response.ErrCode:", r.ErrCode)
	//fmt.Println("Response.ErrMsg:", r.ErrMsg)


	//name1 := "{\"EntityId\":\"1\",\"Timestamp\":\"2021-04-18 00:01:20.123\",\"KeyId\":\"123\",\"ReceiptValue\":532,\"Version\":\"v3\",\"UserName\":\"zsl\",\"OperationType\":\"OperationType\",\"DataType\":\"datp\"}"
	//name2 := "{\"EntityId\":\"2\",\"Timestamp\":\"2021-04-18 00:02:20.213\",\"KeyId\":\"123\",\"ReceiptValue\":532,\"Version\":\"v3\",\"UserName\":\"zsl\",\"OperationType\":\"OperationType\",\"DataType\":\"datp\"}"
	//name3 := "{\"EntityId\":\"3\",\"Timestamp\":\"2021-04-18 00:03:20.321\",\"KeyId\":\"123\",\"ReceiptValue\":532,\"Version\":\"v3\",\"UserName\":\"zsl\",\"OperationType\":\"OperationType\",\"DataType\":\"datp\"}"
	//dataa:=name1+"_"+name2+"_"+name3

	//目前 a  b  c是存证
	//d   e   是交易
	//r, err := c.Dataa(context.Background(), &pb.LedgerTypea{Dataa: dataa})
	//r, err := c.Datab(context.Background(), &pb.LedgerTypeb{Datab: datab})
	//r, err := c.Datac(context.Background(), &pb.LedgerTypec{Datac: datac})
	//r, err := c.Datad(context.Background(), &pb.LedgerTyped{Datad: datad})
	//r, err := c.Datae(context.Background(), &pb.LedgerTypee{Datae: datae})



	var data pb.ServiceAccessData
	var dat1 pb.Transaction
	var dat2 pb.Transaction

	dat1.EntityId = "zsl"
	dat1.CreateTimestamp = "2021-06-02 16:19:44.456"
	dat1.TransactionId = "zsl456"

	dat2.EntityId = "ls"
	dat2.CreateTimestamp = "2021-06-02 16:19:44.456"
	dat2.TransactionId = "ls456"

	data.Transactions = append(data.Transactions, &dat1)
	data.Transactions = append(data.Transactions, &dat2)

	r, err := c.ServiceAccess(context.Background(), &pb.ServiceAccessData{Transactions: data.Transactions})
	if err != nil {
		log.Fatalf("cloud not greet: %v", err)
	}
	fmt.Println("Response.ErrCode:", r.ErrCode)
	fmt.Println("Response.ErrMsg:", r.ErrMsg)
}

func recvVideo(videoData [] pb.DataReceipt)  {

}