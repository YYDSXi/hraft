package utils
//
//import (
//	"fmt"
//	"hraft/domain"
//	pb "hraft/rpc"
//
//	log "github.com/sirupsen/logrus"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//)
//
//var (
//	address = ":8080"
//)
//
//func PutGenesisBlockToTdengine(genesisBlock pb.GenesisBlock, LEDGER_TYPE string) string {
//
//	var err error
//	switch {
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[0]:
//		client := GetTdengineClientVideoLedger()
//		_, err = client.AddGenesisBlock(context.Background(), &genesisBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[1]:
//		client := GetTdengineClientUserLedger()
//		_, err = client.AddGenesisBlock(context.Background(), &genesisBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[2]:
//		client := GetTdengineClientNodeLedger()
//		_, err = client.AddGenesisBlock(context.Background(), &genesisBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[3]:
//		client := GetTdengineClientSensorLedger()
//		_, err = client.AddGenesisBlock(context.Background(), &genesisBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[4]:
//		client := GetTdengineClientAccessLedger()
//		_, err = client.AddGenesisBlock(context.Background(), &genesisBlock)
//	}
//	if err != nil {
//		fmt.Println("PutGenesisBlockToTdengine.err=", err)
//		return "PutGenesisBlockToTdengine错误"
//	} else {
//		return "PutGenesisBlockToTdengine成功"
//	}
//}
//func checkError() {
//	log.Error("当前数据类型与节点账本接口不匹配")
//}
//func PutMinBlockToTdengine(minBlock pb.MinuteTxBlock, dataBlock pb.MinuteDataBlock, LEDGER_TYPE string) string {
//	fmt.Println("PutMinBlockToTdengine========minBlock", minBlock, dataBlock)
//	var err error
//
//	switch {
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[0]:
//
//		client := GetTdengineClientVideoLedger()
//		_, err = client.AddMinuteBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[1]:
//		client := GetTdengineClientUserLedger()
//		_, err = client.AddMinuteBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[2]:
//		client := GetTdengineClientNodeLedger()
//		_, err = client.AddMinuteBlock(context.Background(), &minBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[3]:
//		client := GetTdengineClientSensorLedger()
//		_, err = client.AddMinuteBlock(context.Background(), &minBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[4]:
//		client := GetTdengineClientAccessLedger()
//		_, err = client.AddMinuteBlock(context.Background(), &minBlock)
//	}
//	if err != nil {
//		fmt.Println("err = ", err)
//		return "PutMinBlockToTdengine错误"
//	} else {
//		return "PutMinBlockToTdengine成功"
//	}
//}
//func PutTenMinBlockToTdengine(tenMinuteBlock pb.TenMinuteTxBlock, dataBlock pb.TenMinuteDataBlock, LEDGER_TYPE string) string {
//	var err error
//	switch {
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[0]:
//		client := GetTdengineClientVideoLedger()
//		_, err = client.AddTenMinuteBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[1]:
//		client := GetTdengineClientUserLedger()
//		_, err = client.AddTenMinuteBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[2]:
//		client := GetTdengineClientNodeLedger()
//		_, err = client.AddTenMinuteBlock(context.Background(), &tenMinuteBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[3]:
//		client := GetTdengineClientSensorLedger()
//		_, err = client.AddTenMinuteBlock(context.Background(), &tenMinuteBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[4]:
//		client := GetTdengineClientAccessLedger()
//		_, err = client.AddTenMinuteBlock(context.Background(), &tenMinuteBlock)
//	}
//
//	if err != nil {
//		return "PutMinBlockToTdengine错误"
//	} else {
//		return "PutMinBlockToTdengine成功"
//	}
//}
//func PutDailyBlockToTdengine(dailyBlock pb.DailyTxBlock, dataBlock pb.DailyDataBlock, LEDGER_TYPE string) string {
//	var err error
//	switch {
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[0]:
//		client := GetTdengineClientVideoLedger()
//		_, err = client.AddDailyBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[1]:
//		client := GetTdengineClientUserLedger()
//		_, err = client.AddDailyBlock(context.Background(), &dataBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[2]:
//		client := GetTdengineClientNodeLedger()
//		_, err = client.AddDailyBlock(context.Background(), &dailyBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[3]:
//		client := GetTdengineClientSensorLedger()
//		_, err = client.AddDailyBlock(context.Background(), &dailyBlock)
//	case LEDGER_TYPE == domain.ALL_LEDGER_TYPE_ARRAY[4]:
//		client := GetTdengineClientAccessLedger()
//		_, err = client.AddDailyBlock(context.Background(), &dailyBlock)
//	}
//	if err != nil {
//		return "PutDailyBlockToTdengine错误"
//	} else {
//		return "PutDailyBlockToTdengine成功"
//	}
//}
//
//func GetTdengineClientAccessLedger() pb.AccessLedgerServiceClient {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		panic("connect error")
//	}
//	Client := pb.NewAccessLedgerServiceClient(conn)
//	return Client
//}
//
//func GetTdengineClientNodeLedger() pb.NodeLedgerServiceClient {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		panic("connect error")
//	}
//	Client := pb.NewNodeLedgerServiceClient(conn)
//	return Client
//}
//func GetTdengineClientSensorLedger() pb.SensorLedgerServiceClient {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		panic("connect error")
//	}
//	Client := pb.NewSensorLedgerServiceClient(conn)
//	return Client
//}
//func GetTdengineClientUserLedger() pb.UserLedgerServiceClient {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		panic("connect error")
//	}
//	Client := pb.NewUserLedgerServiceClient(conn)
//	return Client
//}
//func GetTdengineClientVideoLedger() pb.VideoLedgerServiceClient {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		panic("connect error")
//	}
//	Client := pb.NewVideoLedgerServiceClient(conn)
//	return Client
//}
