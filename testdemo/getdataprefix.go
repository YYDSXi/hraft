package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 写入读取
func mainget() {
	// 配置 etcd ,创建客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	// 存储
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/sample_key/name/1", "sample_value_name")
	defer cancel()
	if err != nil {
		fmt.Println("put name failed, err:", err)
		return
	}
	// 存储
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/sample_key/name/2", "sample_value_name_laixhe")
	defer cancel()
	if err != nil {
		fmt.Println("put name failed, err:", err)
		return
	}

	// 获取
	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// resp, err := cli.Get(ctx, "/sample_key/name/")
	// defer cancel()
	// if err != nil {
	// 	fmt.Println("get failed, err:", err)
	// 	return
	// }
	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("get (/sample_key/name/) %s : %s\n", ev.Key, ev.Value)
	// }
	// 获取前缀的 /sample_key/name/ 都返回
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/sample_key/name/", clientv3.WithPrefix())
	defer cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("get all (/sample_key/name/) - %s : %s\n", ev.Key, ev.Value)
	}
	fmt.Println(len(resp.Kvs))
	time.Sleep(time.Duration(20) * time.Second)
}
