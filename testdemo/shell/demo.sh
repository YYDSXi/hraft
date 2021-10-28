
while true
do
  read -p "请输入:" input
  case $input in
   "1")
          echo "1..."
          ;;
   "2")
          echo "2 ..."
          ;;
   "3")
          echo "3..."
          ;;
   *)
          echo "Usage: $name [start|stop|reload]"
          ;;
  esac
done
#while true
#do
#  read -p "请输入:" input
#  if [ $input -eq 1  ];then
#    echo $input
#  elif [ $input -eq 2  ];then
#    echo $input
#  else
#     echo $input
#     break
#  fi
#done
read -n1
