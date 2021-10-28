package serviceregister

import (
	"context"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"hraft/domain"
	"hraft/utils"
)

var(
	HRAFT_NAME = "hraft"
)
//ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli     *clientv3.Client //etcd client
	leaseID clientv3.LeaseID //租约ID
	//租约keepalieve相应chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string //key
	val           string //value
}

//NewServiceRegister 新建注册服务
func NewServiceRegister(cli *clientv3.Client,endpoints []string, key, val string, lease int64) (*ServiceRegister, error) {
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints:   endpoints,
	//	DialTimeout: 5 * time.Second,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	ser := &ServiceRegister{
		cli: cli,
		key: key,
		val: val,
	}

	//申请租约设置时间keepalive
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return ser, nil
}

//设置租约
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	//设置租约时间
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	//注册服务并绑定租约
	_, err = s.cli.Put(context.Background(), s.key, s.val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	//设置续租 定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)

	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	//log.Info(s.leaseID)
	s.keepAliveChan = leaseRespChan
	//log.Info("Put key:%s  val:%s  success!", s.key, s.val)
	return nil
}

//ListenLeaseRespChan 监听 续租情况
func (s *ServiceRegister) ListenLeaseRespChan() {
	//for leaseKeepResp := range s.keepAliveChan {
	//	log.Info("续约成功", leaseKeepResp)
	//}

	for _ = range s.keepAliveChan {
		//log.Info("续约成功", leaseKeepResp)
	}

	log.Info("关闭续租")
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
	//撤销租约
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Info("撤销租约")
	return s.cli.Close()
}

func RegisterService(cli *clientv3.Client,globalLeaderName string) {

	var endpoints = []string{"localhost:2379"}
	var chainName = utils.Conf.ChainName
	if chainName == ""{
		chainName = "scope"
	}
	for i:=0;i<len(domain.GlobalLedgerArray);i++{
		ipPortLeader:=utils.Conf.Consensus.EtcdGroup[globalLeaderName].HraftAddress
		//serLeader, _ := NewServiceRegister(cli,endpoints, domain.GlobalLedgerArray[i]+"/"+HRAFT_NAME+"/"+globalLeaderName, ipPortLeader, 5)
		serLeader, _ := NewServiceRegister(cli,endpoints, chainName+"/"+HRAFT_NAME+"/"+domain.GlobalLedgerArray[i]+"/", ipPortLeader, 5)

		log.Info("已注册服务key=",chainName+"/"+HRAFT_NAME+"/"+domain.GlobalLedgerArray[i]+"/")
		log.Info("已注册服务val=",ipPortLeader)
		//监听续租相应chan
		go serLeader.ListenLeaseRespChan()

		followerArray := utils.Conf.Common.LedgerName[domain.GlobalLedgerArray[i]].Follower
		if followerArray == nil || len(followerArray)==0{
			continue
		}
		/*
		for j:=0;j<len(followerArray);j++{
			ipPortFollower:=utils.Conf.Consensus.EtcdGroup[globalLeaderName].HraftAddress
			serFollower, _ := NewServiceRegister(cli,endpoints, domain.GlobalLedgerArray[i]+"/"+HRAFT_NAME+"/"+"follower"+strconv.Itoa(j+1), ipPortFollower, 5)

			log.Info("已注册服务key=",domain.GlobalLedgerArray[i]+"/"+HRAFT_NAME+"/"+"follower"+strconv.Itoa(j+1))
			log.Info("已注册服务val=",serFollower)
			//监听续租相应chan
			go serLeader.ListenLeaseRespChan()
		}
		*/

	}


	select {
	// case <-time.After(20 * time.Second):
	// 	ser.Close()
	}
}
