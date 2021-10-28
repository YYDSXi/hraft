package utils
//
//import (
//	"fmt"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"os"
//	"os/exec"
//	"regexp"
//	"strconv"
//	"strings"
//	"time"
//)
//
//var (
//	videoAllDataCountsKey="video:AllDataCountsKey"
//	videoAllDataSizeKey="video:AllDataSizeKey"
//	videoAllCurDayDelayDataKey="video:AllCurDayDelayData"
//
//	userBehaviourAllDataCountsKey="user_behaviour:AllDataCountsKey"
//	userBehaviourAllDataSizeKey="user_behaviour:AllDataSizeKey"
//	userBehaviourAllCurDayDelayDataKey="user_behaviour:AllCurDayDelayData"
//
//	nodeCredibleAllDataCountsKey="node_credible:AllDataCountsKey"
//	nodeCredibleAllDataSizeKey="node_credible:AllDataSizeKey"
//	nodeCredibleAllCurDayDelayDataKey="node_credible:AllCurDayDelayData"
//
//	sensorAllDataCountsKey="sensor:AllDataCountsKey"
//	sensorAllDataSizeKey="sensor:AllDataSizeKey"
//	sensorAllCurDayDelayDataKey="sensor:AllCurDayDelayData"
//
//	serviceAccessAllDataCountsKey="service_access:AllDataCountsKey"
//	serviceAccessAllDataSizeKey="service_access:AllDataSizeKey"
//	serviceAccessAllCurDayDelayDataKey="service_access:AllCurDayDelayData"
//)
////调用命令行查数据
//func ExecQueryLedger(op string,key string,ip string,des string){
//	cmd := exec.Command("etcdctl", op, key,"--endpoints="+ip+":2379")
//	buf, _ := cmd.Output()
//	ans := "0"
//	fmt.Print("节点["+ip+"],")
//	if string(buf) != ""{
//		resultArr := strings.Split(string(buf),"\n")
//		if resultArr[1] != ""{
//			ans = resultArr[1]
//		}
//	}
//	fmt.Print(des)
//	fmt.Print(ans)
//	fmt.Println()
//}
//func ExecGetDataInt(op string,key string,ip string)int{
//	cmd := exec.Command("etcdctl", op, key,"--endpoints="+ip+":2379")
//	buf, _ := cmd.Output()
//	resultInt := 0
//	if string(buf) != ""{
//		resultArr := strings.Split(string(buf),"\n")
//		if resultArr[1] != ""{
//			resultInt,_=strconv.Atoi(resultArr[1])
//		}
//	}
//	return resultInt
//}
//func HraftStatus()  {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		fmt.Printf("connect to etcd failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("connect to etcd success")
//	defer cli.Close()
//
//	var op string
//	var ip string
//	for true {
//		PrintInfo()
//		fmt.Println("请输入操作:")
//		fmt.Scan(&op)
//		if op =="1" || op == "2" || op == "3" || op == "4"|| op == "5"|| op == "6" || op == "9" || op == "10"{
//			for true{
//				fmt.Println("请输入节点IP,先退出程序再根据【8】获取节点IP")
//				fmt.Scan(&ip)
//				if !IsIp(ip){
//					fmt.Println("输入IP不合法,请重新输入,")
//				}else{
//					break
//				}
//			}
//		}
//		if ip == ""{
//			ip = "127.0.0.1"
//		}
//		switch {
//		case op == "1" :
//			ExecQueryLedger("get",videoAllDataCountsKey,ip,"video账本总数据条目:")
//			ExecQueryLedger("get",videoAllDataSizeKey,ip,"video账本总数据大小:")
//			ExecQueryLedger("get",videoAllCurDayDelayDataKey,ip,"video账本延时数据条目:")
//			//PrintStatisticalInfo(cli,videoAllDataCountsKey,"video账本总数据条目:")
//			//PrintStatisticalInfo(cli,videoAllDataSizeKey,"video账本总数据大小:")
//			//PrintStatisticalInfo(cli,videoAllCurDayDelayDataKey,"video账本延时数据条目:")
//		case op == "2":
//			ExecQueryLedger("get",userBehaviourAllDataCountsKey,ip,"user_behaviour账本总数据条:")
//			ExecQueryLedger("get",userBehaviourAllDataSizeKey,ip,"user_behaviour账本总数据大小:")
//			ExecQueryLedger("get",userBehaviourAllCurDayDelayDataKey,ip,"user_behaviour账本延时数据条目:")
//			//PrintStatisticalInfo(cli,userBehaviourAllDataCountsKey,"user_behaviour账本总数据条:")
//			//PrintStatisticalInfo(cli,userBehaviourAllDataSizeKey,"user_behaviour账本总数据大小:")
//			//PrintStatisticalInfo(cli,userBehaviourAllCurDayDelayDataKey,"user_behaviour账本延时数据条目:")
//		case op == "3" :
//			ExecQueryLedger("get",nodeCredibleAllDataCountsKey,ip,"node_credible账本总数据条目:")
//			ExecQueryLedger("get",nodeCredibleAllDataSizeKey,ip,"node_credible账本总数据大小:")
//			ExecQueryLedger("get",nodeCredibleAllCurDayDelayDataKey,ip,"node_credible账本延时数据条目:")
//			//PrintStatisticalInfo(cli,nodeCredibleAllDataCountsKey,"node_credible账本总数据条目:")
//			//PrintStatisticalInfo(cli,nodeCredibleAllDataSizeKey,"node_credible账本总数据大小:")
//			//PrintStatisticalInfo(cli,nodeCredibleAllCurDayDelayDataKey,"node_credible账本延时数据条目:")
//		case op == "4":
//			ExecQueryLedger("get",sensorAllDataCountsKey,ip,"sensor账本总数据条目:")
//			ExecQueryLedger("get",sensorAllDataSizeKey,ip,"sensor账本总数据大小:")
//			ExecQueryLedger("get",sensorAllCurDayDelayDataKey,ip,"sensor账本延时数据条目:")
//			//PrintStatisticalInfo(cli,sensorAllDataCountsKey,"sensor账本总数据条目:")
//			//PrintStatisticalInfo(cli,sensorAllDataSizeKey,"sensor账本总数据大小:")
//			//PrintStatisticalInfo(cli,sensorAllCurDayDelayDataKey,"sensor账本延时数据条目:")
//		case op == "5":
//			ExecQueryLedger("get",serviceAccessAllDataCountsKey,ip,"service_access账本总数据条目:")
//			ExecQueryLedger("get",serviceAccessAllDataSizeKey,ip,"service_access账本总数据大小:")
//			ExecQueryLedger("get",serviceAccessAllCurDayDelayDataKey,ip,"service_access账本延时数据条目:")
//			//PrintStatisticalInfo(cli,serviceAccessAllDataCountsKey,"service_access账本总数据条目:")
//			//PrintStatisticalInfo(cli,serviceAccessAllDataSizeKey,"service_access账本总数据大小:")
//			//PrintStatisticalInfo(cli,serviceAccessAllCurDayDelayDataKey,"service_access账本延时数据条目:")
//		case op == "6":
//			//a1:=GetDataInt(cli,videoAllDataCountsKey)
//			a1:=ExecGetDataInt("get",videoAllDataCountsKey,ip)
//			a2:=ExecGetDataInt("get",userBehaviourAllDataCountsKey,ip)
//			a3:=ExecGetDataInt("get",nodeCredibleAllDataCountsKey,ip)
//			a4:=ExecGetDataInt("get",sensorAllDataCountsKey,ip)
//			a5:=ExecGetDataInt("get",serviceAccessAllDataCountsKey,ip)
//			suma := a1+a2+a3+a4+a5
//			fmt.Print("节点["+ip+"],")
//			fmt.Println("五类账本总数据条目:",suma)
//			b1:=ExecGetDataInt("get",videoAllDataSizeKey,ip)
//			b2:=ExecGetDataInt("get",userBehaviourAllDataSizeKey,ip)
//			b3:=ExecGetDataInt("get",nodeCredibleAllDataSizeKey,ip)
//			b4:=ExecGetDataInt("get",sensorAllDataSizeKey,ip)
//			b5:=ExecGetDataInt("get",serviceAccessAllDataSizeKey,ip)
//			sumb := b1+b2+b3+b4+b5
//			fmt.Print("节点["+ip+"],")
//			fmt.Println("五类账本总数据量:",sumb)
//			c1:=ExecGetDataInt("get",videoAllCurDayDelayDataKey,ip)
//			c2:=ExecGetDataInt("get",userBehaviourAllCurDayDelayDataKey,ip)
//			c3:=ExecGetDataInt("get",nodeCredibleAllCurDayDelayDataKey,ip)
//			c4:=ExecGetDataInt("get",sensorAllCurDayDelayDataKey,ip)
//			c5:=ExecGetDataInt("get",serviceAccessAllCurDayDelayDataKey,ip)
//			sumc := c1+c2+c3+c4+c5
//			fmt.Print("节点["+ip+"],")
//			fmt.Println("五类账本总延时数据条目:",sumc)
//		case op == "7":
//			cmd := exec.Command("etcdctl","endpoint","status","--endpoints="+ip+":2379","--write-out=table")
//			buf, _ := cmd.Output()
//			fmt.Printf("当前节点状态信息: \n %s\n",buf)
//		case op == "8":
//			cmd := exec.Command("etcdctl", "member", "list", "--write-out=table")
//			buf, _ := cmd.Output()
//			fmt.Printf("集群状态信息: \n %s\n",buf)
//		case op == "9":
//			cmd := exec.Command("etcdctl", "get", "AllKeysCounts","--endpoints="+ip+":2379")
//			buf, _ := cmd.Output()
//			ans := "0"
//			if string(buf) != ""{
//				ans = string(buf)
//			}
//			fmt.Printf("所有key统计数量: \n %s\n",ans)
//			//allKeyCounts := GetDataStr(cli,"AllKeysCounts")
//			//fmt.Printf("所有key统计数量: \n%s\n",allKeyCounts)
//		case op == "10":
//			PrintAutoWorkInfo(cli,"video",ip)
//			PrintAutoWorkInfo(cli,"user_behaviour",ip)
//			PrintAutoWorkInfo(cli,"node_credible",ip)
//			PrintAutoWorkInfo(cli,"sensor",ip)
//			PrintAutoWorkInfo(cli,"service_access",ip)
//		case op == "0":
//			os.Exit(1)
//		default:
//			fmt.Println("输入不合法！")
//		}
//	}
//
//}
//
////封装打印信息
//func PrintStatisticalInfo(cli *clientv3.Client,key string,info string){
//	getResponse := GetData(cli, key, 5 * time.Second)
//	for _, ev := range getResponse.Kvs {
//		ans := string(ev.Value)
//		fmt.Print(info)
//		if ans == ""{
//			ans = "0"
//		}
//		fmt.Print(ans)
//		fmt.Println()
//	}
//}
//
//func GetDataInt(cli *clientv3.Client,key string) int{
//	getResponse := GetData(cli, key, 5 * time.Second)
//	for _, ev := range getResponse.Kvs {
//		ansStr := string(ev.Value)
//		ansInt, _ := strconv.Atoi(ansStr)
//		return ansInt
//	}
//	return 0
//}
//func PrintAutoWorkInfo(cli *clientv3.Client,ledgerType string,ip string){
//	indexMinInt := GetIndexMinInt()
//	yearMonthDay := time.Now().Format("2006-01-02")
//	dataString := GetDataStr(cli,yearMonthDay+":"+ledgerType+":MINUTE:"+strconv.Itoa(indexMinInt-1),ip)
//
//	preMin := time.Now().Add(-time.Minute * 1).Format("2006-01-02 15:04")
//	fmt.Print(ledgerType)
//	fmt.Print("账本 第",indexMinInt-1)
//	fmt.Print("个分钟块打包成功,打包时间:",preMin+":30,")
//	fmt.Println(",数据如下:",dataString)
//
//
//	fmt.Print(ledgerType)
//	fmt.Print("账本 第",indexMinInt)
//	fmt.Println("个分钟块正在打包...")
//}
//func GetDataStr(cli *clientv3.Client,key string,ip string) string{
//	cmd := exec.Command("etcdctl", "get", key,"--endpoints="+ip+":2379")
//	buf, _ := cmd.Output()
//	resultArr := strings.Split(string(buf),"\n")
//	return resultArr[1]
//}
//func PrintInfo() {
//	fmt.Println("1.video账本统计信息")
//	fmt.Println("2.user_behaviour账本统计信息")
//	fmt.Println("3.node_credible账本统计信息")
//	fmt.Println("4.sensor账本统计信息")
//	fmt.Println("5.service_access账本统计信息")
//	fmt.Println("6.五类账本统计信息")
//	fmt.Println("7.本节点信息描述")
//	fmt.Println("8.集群信息")
//	fmt.Println("9.所有key统计记录")
//	fmt.Println("10.查询自动任务状态")
//	fmt.Println("0.退出")
//}
//func IsIp(ip string) (b bool) {
//	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
//		return false
//	}
//	return true
//}