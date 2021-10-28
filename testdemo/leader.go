package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	DialTimeout = 5 * time.Second
	Endpoints   = []string{"127.0.0.1:2379", "127.0.0.2:2379"}
	ConfPath    = "./config.yaml"
)

func main4() {
	cliTemp, errNewCli := clientv3.New(clientv3.Config{
		Endpoints:   Endpoints,
		DialTimeout: DialTimeout,
	})
	if errNewCli != nil {
		println(errNewCli)
	}
	cli := cliTemp

	defer cli.Close()

	for _, ep := range Endpoints {
		resp, err := cli.Status(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("endpoint: %s / Leader: %v\n", ep, resp.Header.MemberId == resp.Leader)
	}

	// for _, ep := range Endpoints {
	// 	cli, err := clientv3.New(clientv3.Config{
	// 		Endpoints:   []string{ep},
	// 		DialTimeout: 5 * time.Second,
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer cli.Close()

	// 	resp, err := cli.Status(context.Background(), ep)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("endpoint: %s / Leader: %v\n", ep, resp.Header.MemberId == resp.Leader)
	// }
}
