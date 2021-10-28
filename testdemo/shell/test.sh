source ./dataStruct.sh

function PrintInfo() {
  echo "========================================="
  # dataReceipt1_1T dataReceipt1_2T
  echo "1.多笔无异常存证数据"

  # transaction2_1T transaction2_2T
  echo "2.多笔无异常交易数据"

  # dataReceipt3_1T dataReceipt3_2T
  echo "3.多笔时间戳相同存证数据"

  # transaction4_1T transaction4_2T
  echo "4.多笔时间戳相同交易数据"

  # dataReceipt5_1T 正常 dataReceipt5_2T 正常 dataReceipt5_3F 超过10m时间
  echo "5.多笔延时数据存证交易数据(只接收近10m内数据)"

  # dataReceipt6_1F dataReceipt6_2F 两笔存证数据ID相同
  echo "6.多笔存证ID相同数据(根据存证ID/交易ID去重)"

  # dataReceipt7_1F dataReceipt7_1F 缺失时间戳 缺失ID
  echo "7.存证/交易数据时间戳或ID缺失(过滤掉)"

  # dataReceipt8_1F dataReceipt8_2F 时间戳不带毫秒 格式不正确
  echo "8.存证/交易数据时间戳不合法"

  echo "0.清理所有数据"

  echo "-1.退出"
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
          echo "---------------------数据------------------"
          echo $dataReceipt1_1T
          echo $dataReceipt1_2T
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt1_1T}"  "${dataReceipt1_2T}"
          ;;
   "2")
          echo "---------------------数据------------------"
          echo ${transaction2_1T}
          echo ${transaction2_2T}
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerTransaction "${transaction2_1T}"  "${transaction2_2T}"
          ;;
   "3")
          echo "---------------------数据------------------"
          echo $dataReceipt3_1T
          echo $dataReceipt3_2T
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt3_1T}"  "${dataReceipt3_2T}"
          ;;
   "4")
          echo "---------------------数据------------------"
          echo $transaction4_1T
          echo $transaction4_2T
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerTransaction "${transaction4_1T}"  "${transaction4_2T}"
          ;;
   "5")
          echo "---------------------数据------------------"
          echo $dataReceipt5_1T
          echo $dataReceipt5_2T
          echo $dataReceipt5_3F
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt5_1T}"  "${dataReceipt5_2T}" "${dataReceipt5_3F}"
          ;;
   "6")
          echo "---------------------数据------------------"
          echo $dataReceipt6_1F
          echo $dataReceipt6_2F
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt6_1F}"  "${dataReceipt6_2F}"
          ;;
   "7")
          echo "---------------------数据------------------"
          echo $dataReceipt7_1F
          echo $dataReceipt7_2F
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt7_1F}"  "${dataReceipt7_2F}"
          ;;
   "8")
          echo "---------------------数据------------------"
          echo $dataReceipt8_1F
          echo $dataReceipt8_2F
          echo "---------------------数据------------------"
          ./shell_redis_client $ledgerReceipt "${dataReceipt8_1F}"  "${dataReceipt8_2F}"
          ;;
   "0")
          cd ../../bin && etcdctl del "" --prefix
          ;;
   "-1")
          exit 1
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
