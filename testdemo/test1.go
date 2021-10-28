package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/client/v3/concurrency"

	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd lease

func C1(client *clientv3.Client) {
	// 生成一个30s超时的上下文
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	// 获取租约
	response, e := client.Grant(timeout, 30)
	if e != nil {
		log.Fatal(e.Error())
	}
	// 通过租约创建session
	session, e := concurrency.NewSession(client, concurrency.WithLease(response.ID))
	if e != nil {
		log.Fatal(e.Error())
	}
	defer session.Close()
	// 通过session和锁前缀
	mutex := concurrency.NewMutex(session, "/lock")
	e = mutex.Lock(timeout)
	if e != nil {
		log.Fatal(e.Error())
	}

	// 业务逻辑
	fmt.Println("acquired lock for s1")
	fmt.Println("s1 ...")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("released lock for s1")

	// 释放锁
	defer mutex.Unlock(timeout)
}
