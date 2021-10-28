
while true
do
  read -p "请输入查询key:" inputKey
  echo "查询结果为:"
  etcdctl get $inputKey
done

read -n1
