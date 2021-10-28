source ./utils.sh
videoAllDataCountsKey="video:AllDataCountsKey"
videoAllDataSizeKey="video:AllDataSizeKey"
videoAllCurDayDelayDataKey="video:AllCurDayDelayData"

userBehaviourAllDataCountsKey="user_behaviour:AllDataCountsKey"
userBehaviourAllDataSizeKey="user_behaviour:AllDataSizeKey"
userBehaviourAllCurDayDelayDataKey="user_behaviour:AllCurDayDelayData"

nodeCredibleAllDataCountsKey="node_credible:AllDataCountsKey"
nodeCredibleAllDataSizeKey="node_credible:AllDataSizeKey"
nodeCredibleAllCurDayDelayDataKey="node_credible:AllCurDayDelayData"

sensorAllDataCountsKey="sensor:AllDataCountsKey"
sensorAllDataSizeKey="sensor:AllDataSizeKey"
sensorAllCurDayDelayDataKey="sensor:AllCurDayDelayData"

serviceAccessAllDataCountsKey="service_access:AllDataCountsKey"
serviceAccessAllDataSizeKey="service_access:AllDataSizeKey"
serviceAccessAllCurDayDelayDataKey="service_access:AllCurDayDelayData"

function PrintInfo() {
  echo "1.video账本信息"
  echo "2.user_behaviour账本统计信息"
  echo "3.node_credible账本统计信息"
  echo "4.sensor账本统计信息"
  echo "5.service_access账本统计信息"
  echo "6.五类账本统计信息"
  echo "7.本节点信息描述"
  echo "8.集群信息"
  echo "9.所有key统计记录"
  echo "10.查询自动任务状态" #加个时间点 什么时候打的包
}

while true
do
  PrintInfo
  read -p "请输入:" input
  case $input in
   "1")
          echo "---------------------------------------"
          echo "video账本总数据条目:"
          videoAllDataCounts=$(etcdctl get $videoAllDataCountsKey)
          if [ ! -n "${videoAllDataCounts: 23}" ]; then
            echo "0"
          else
            echo ${videoAllDataCounts: 23}
          fi

          echo "video账本总数据大小:"
          videoAllDataSize=$(etcdctl get $videoAllDataSizeKey)
          if [ ! -n "${videoAllDataSize: 21}" ]; then
            echo "0"
          else
            echo ${videoAllDataSize: 21}
          fi

          echo "video账本延时数据条目:"
          videoAllCurDayDelayData=$(etcdctl get $videoAllCurDayDelayDataKey)
          if [ ! -n "${videoAllCurDayDelayData: 25}" ]; then
            echo "0"
          else
            echo ${videoAllCurDayDelayData: 25}
          fi
          echo "---------------------------------------"
          ;;
   "2")
          echo "---------------------------------------"
          echo "user_behaviour账本总数据条目:"
          userBehaviourAllDataCounts=$(etcdctl get $userBehaviourAllDataCountsKey)
          if [ ! -n "${userBehaviourAllDataCounts: 32}" ]; then
            echo "0"
          else
            echo ${userBehaviourAllDataCounts: 32}
          fi

          echo "user_behaviour账本总数据大小:"
          userBehaviourAllDataSize=$(etcdctl get $userBehaviourAllDataSizeKey)
          if [ ! -n "${userBehaviourAllDataSize: 30}" ]; then
            echo "0"
          else
            echo ${userBehaviourAllDataSize: 30}
          fi

          echo "user_behaviour账本延时数据条目:"
          userBehaviourAllCurDayDelayData=$(etcdctl get $userBehaviourAllCurDayDelayDataKey)
          if [ ! -n "${userBehaviourAllCurDayDelayData: 34}" ]; then
            echo "0"
          else
            echo ${userBehaviourAllCurDayDelayData: 34}
          fi
          echo "---------------------------------------"
          ;;
   "3")
          echo "---------------------------------------"
          echo "node_credible账本总数据条目:"
          nodeCredibleAllDataCounts=$(etcdctl get $nodeCredibleAllDataCountsKey)
          if [ ! -n "${nodeCredibleAllDataCounts: 31}" ]; then
            echo "0"
          else
            echo ${nodeCredibleAllDataCounts: 31}
          fi

          echo "node_credible账本总数据大小:"
          nodeCredibleAllDataSize=$(etcdctl get $nodeCredibleAllDataSizeKey)
          if [ ! -n "${nodeCredibleAllDataSize: 29}" ]; then
            echo "0"
          else
            echo ${nodeCredibleAllDataSize: 29}
          fi
          echo "node_credible账本延时数据条目:"
          nodeCredibleAllCurDayDelayData=$(etcdctl get $nodeCredibleAllCurDayDelayDataKey)
          if [ ! -n "${nodeCredibleAllCurDayDelayData: 33}" ]; then
            echo "0"
          else
            echo ${nodeCredibleAllCurDayDelayData: 33}
          fi
          echo "---------------------------------------"
          ;;
   "4")
          echo "---------------------------------------"
          echo "sensor账本总数据条目:"
          sensorAllDataCounts=$(etcdctl get $sensorAllDataCountsKey)
          if [ ! -n "${sensorAllDataCounts: 24}" ]; then
            echo "0"
          else
            echo ${sensorAllDataCounts: 24}
          fi

          echo "sensor账本总数据大小:"
          sensorAllDataSize=$(etcdctl get $sensorAllDataSizeKey)
          if [ ! -n "${sensorAllDataSize: 22}" ]; then
            echo "0"
          else
            echo ${sensorAllDataSize: 22}
          fi
          echo "sensor账本延时数据条目:"
          sensorAllCurDayDelayData=$(etcdctl get $sensorAllCurDayDelayDataKey)
          if [ ! -n "${sensorAllCurDayDelayData: 26}" ]; then
            echo "0"
          else
            echo ${sensorAllCurDayDelayData: 26}
          fi
          echo "---------------------------------------"
          ;;
   "5")
          echo "---------------------------------------"
          echo "service_access账本总数据条目:"
          serviceAccessAllDataCounts=$(etcdctl get $serviceAccessAllDataCountsKey)
          if [ ! -n "${serviceAccessAllDataCounts: 32}" ]; then
            echo "0"
          else
            echo ${serviceAccessAllDataCounts: 32}
          fi

          echo "service_access账本总数据大小:"
          serviceAccessAllDataSize=$(etcdctl get $serviceAccessAllDataSizeKey)
          if [ ! -n "${serviceAccessAllDataSize: 30}" ]; then
            echo "0"
          else
            echo ${serviceAccessAllDataSize: 30}
          fi
          echo "service_access账本延时数据条目:"
          serviceAccessAllCurDayDelayData=$(etcdctl get $serviceAccessAllCurDayDelayDataKey)
          if [ ! -n "${serviceAccessAllCurDayDelayData: 34}" ]; then
            echo "0"
          else
            echo ${serviceAccessAllCurDayDelayData: 34}
          fi
          echo "---------------------------------------"
          ;;
   "6")
          echo "---------------------------------------"
          m1=1
          m2=2
          m3=3
          sumall $m1 ${videoAllDataCounts: 23} ${userBehaviourAllDataCounts: 32}  ${nodeCredibleAllDataCounts: 31} ${sensorAllDataCounts: 24} ${serviceAccessAllDataCounts: 32}
          sumall $m2 ${videoAllDataSize: 21} ${userBehaviourAllDataSize: 30} ${nodeCredibleAllDataSize: 29} ${sensorAllDataSize: 22}
          sumall $m3 ${videoAllCurDayDelayData: 25} ${userBehaviourAllCurDayDelayData: 34} ${nodeCredibleAllCurDayDelayData: 33} ${sensorAllCurDayDelayData: 26} ${serviceAccessAllCurDayDelayData: 34}
          echo "---------------------------------------"
          ;;
   "7")
          #当前节点状态信息
          etcdctl endpoint status --write-out="table"
          ;;
   "8")
          #集群状态信息
          etcdctl member list --write-out="table"
          ;;
   "9")
          echo "所有key统计数量:"
          etcdctl get "AllKeysCounts" --prefix
          ;;
   "10")
          echo "当前自动任务状态:"
          curAutoWorkState
          ;;
   "-1")
          exit 1
          ;;
   *)
          echo "输入错误！"
          ;;
  esac
done

read -n1
