package main

import (
	"encoding/json"
	"fmt"
	pb "hraft/rpc"
)

func main() {
	str := "{\"CreateTimestamp\":\"\",\"EntityId\":\"fg79v9fg7rr1g\",\"TransactionId\":\"dnsjhja\",\"Initiator\":\"g52\",\"Receipt\":\"from\",\"TxAmount\":10.0,\"DataType\":\"测试123001\",\"ServiceType\":\"e218hen\",\"Remark\":\"me\",\"BlockIdentify\":\"219e\"}"
	var transaction pb.Transaction
	fmt.Println("str[i]=", str)
	err := json.Unmarshal([]byte(str), &transaction)
	fmt.Println("transaction=", transaction)
	if err != nil {
		fmt.Println(err)
	}
}
