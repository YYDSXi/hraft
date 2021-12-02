source ./dataStruct.sh

function PrintInfo() {
  echo "========================================="
  # dataReceipt1_1T dataReceipt1_2T
  echo "1.1000笔无异常存证数据"

  # transaction2_1T transaction2_2T
  echo "2.5000笔无异常存证数据"

  # dataReceipt3_1T dataReceipt3_2T
  echo "3.10000笔无异常存证数据"

  # transaction4_1T transaction4_2T
  echo "4.15000笔无异常存证数据"

  # dataReceipt5_1T 正常 dataReceipt5_2T 正常 dataReceipt5_3F 超过10m时间
  echo "5.20000笔无异常存证数据"

  # dataReceipt6_1F dataReceipt6_2F 两笔存证数据ID相同
  echo "6.25000笔无异常存证数据"

  # dataReceipt7_1F dataReceipt7_1F 缺失时间戳 缺失ID
  echo "7.40000笔无异常存证数据"

  # dataReceipt8_1F dataReceipt8_2F 时间戳不带毫秒 格式不正确
  echo "8.50000笔无异常存证数据"

  echo "9.60000笔无异常存证数据"

  echo "10.70000笔无异常存证数据"
}
#暂定存证类型数据 存到video账本
ledgerReceipt="video"
#暂定交易类型数据 存到node_credible账本
ledgerTransaction="node_credible"

while true
do
  PrintInfo
  read -p "请输入:" input
  ReflushTime
  case $input in
   "1")
         # echo "---------------------数据------------------"
         # echo $dataReceipt1_1T
          #echo $dataReceipt1_2T
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "1000"
          ;;
   "2")
       
         # echo ${transaction2_1T}
         # echo ${transaction2_2T}
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "5000"  
          ;;
   "3")
         
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "10000"
          ;;
   "4")
         
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "15000"
          ;;
   "5")
          
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "20000"
          ;;
   "6")
         
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "25000"
          ;;
   "7")
         
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "40000"
          ;;
   "8")
          
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "50000"
          ;;
   "9")
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "60000"
          ;;
   "10")
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "70000"
          ;;
   *)
          echo "输入错误！"
          ;;
  esac
done
#./shell_redis_client "video" "${dataReceipt123001}"  "${dataReceipt123002}"

#hraftPath="/home/zsl/data/gopath/src/hraft"
#cd  ${hraftPath}
#echo "切换到项目目录：$(pwd)"
#
#nohup ./bin/etcd & > nohup.out
#echo "开启etcd服务端！"
#
#./bin/etcdctl del "" --prefix
#echo "删除etcd中原始数据成功！"
#
#go run main.go $1 &
#echo "运行主程序！"
#
#
#go run ./testdemo/redis_client.go &
#echo "运行测试数据！"



read -n1
