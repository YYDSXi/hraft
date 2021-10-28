
sumall(){

  what=$1
  sum=0
  arg1=$2
  arg2=$3
  arg3=$4
  arg4=$5
  arg5=$6

  if [ -n "${arg1}" ]; then
    sum=$[$sum+$arg1]
  fi
  if [ -n "${arg2}" ]; then
    sum=$[$sum+$arg2]
  fi
  if [ -n "${arg3}" ]; then
    sum=$[$sum+$arg3]
  fi
  if [ -n "${arg4}" ]; then
    sum=$[$sum+$arg4]
  fi
  if [ -n "${arg5}" ]; then
    sum=$[$sum+$arg5]
  fi

  if [ $what = "1" ]; then
    echo "五类账本总数据条目:"
    echo $sum
  elif [ $what = "2" ]; then
    echo "五类账本总数据量:"
    echo $sum
  else
    echo "五类账本总延时数据条目:"
    echo $sum
  fi
}

#当前自动任务状态
curAutoWorkState(){
  ymdNow=$(date +'%Y-%m-%d') #2021-06-28
  hmsNow=$(date -d '0 minute ago' +'%H:%M:%S')
  ymd_hms_Now=$ymdNow" "$hmsNow
  #2021-06-21 15:20:21
  hour=${ymd_hms_Now: 11: 2}
  min=${ymd_hms_Now: 14: 2}
  second=${ymd_hms_Now: 17: 2}
  indexMin=$[$hour*60+$min]
  preIndexMin=$[indexMin-1]

  echo "------------------------------------------------------------"
  echo "video账本 第"${preIndexMin}"个分钟块打包成功,数据如下:"
  etcdctl get $ymdNow":""video:MINUTE:"$preIndexMin
  echo "video账本 第"${indexMin}"个分钟块正在打包..."

  echo "------------------------------------------------------------"
  echo "user_behaviour账本 第"${preIndexMin}"个分钟块打包成功,数据如下:"
  etcdctl get $ymdNow":""user_behaviour:MINUTE:"$preIndexMin
  echo "user_behaviour账本 第"${indexMin}"个分钟块正在打包..."

  echo "------------------------------------------------------------"
  echo "node_credible账本 第"${preIndexMin}"个分钟块打包成功,数据如下:"
  etcdctl get $ymdNow":""node_credible:MINUTE:"$preIndexMin
  echo "node_credible账本 第"${indexMin}"个分钟块正在打包..."

  echo "------------------------------------------------------------"
  echo "sensor账本 第"${preIndexMin}"个分钟块打包成功,数据如下:"
  etcdctl get $ymdNow":""sensor:MINUTE:"$preIndexMin
  echo "sensor账本 第"${indexMin}"个分钟块正在打包..."

  echo "------------------------------------------------------------"
  echo "service_access账本 第"${preIndexMin}"个分钟块打包成功,数据如下:"
  etcdctl get $ymdNow":""service_access:MINUTE:"$preIndexMin
  echo "service_access账本 第"${indexMin}"个分钟块正在打包..."
  echo "------------------------------------------------------------"
}
