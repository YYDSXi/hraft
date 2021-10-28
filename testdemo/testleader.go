package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

// etcd lease

func main1() {

	client, errNewCli := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if errNewCli != nil {
		println(errNewCli)
	}

	contextLeader := clientv3.WithRequireLeader(context.Background())

	// 生成一个30s超时的上下文
	timeout, cancelFunc := context.WithTimeout(contextLeader, 30*time.Second)
	defer cancelFunc()
	// 获取租约
	response, e := client.Grant(contextLeader, int64(10))
	if e != nil {
		log.Fatal(e.Error())
	}
	// 通过租约创建session
	session, e := concurrency.NewSession(client, concurrency.WithLease(response.ID))
	if e != nil {
		log.Fatal(e.Error())
	}
	//defer session.Close()
	// 通过session和锁前缀
	mutex := concurrency.NewMutex(session, "t12")
	e = mutex.Lock(timeout)
	if e != nil {
		log.Fatal(e.Error())
	}

	// 业务逻辑
	// put
	response1, err := client.Put(timeout, "t12", "t12", clientv3.WithLease(response.ID))
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println(response1)

	// 释放锁
	//defer mutex.Unlock(timeout)
	//mutex.Unlock(contextLeader)
}
