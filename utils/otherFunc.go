package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetStringMD5(origin string) string {
	m := md5.New()
	m.Write([]byte(origin))
	return hex.EncodeToString(m.Sum(nil))
}

func GetIndexMinInt() int {
	timeFormatString := time.Now().Format("2006-01-02 15:04:05")
	//var naosecond = time.Now().Nanosecond()/1e6
	dayTimeArray := strings.Split(timeFormatString, " ")
	minTimeArray := strings.Split(dayTimeArray[1], ":")
	hourInt, _ := strconv.Atoi(minTimeArray[0])
	minInt, _ := strconv.Atoi(minTimeArray[1])
	indexMinInt := hourInt*60 + minInt
	//indexMinString := strconv.Itoa(indexMinInt)
	//timeCorrect:=fmt.Sprintf("%s.%d",timeFormatString,naosecond)
	return indexMinInt
}
func GetNowSecondInt() int {
	timeFormatString := time.Now().Format("2006-01-02 15:04:05")
	//var naosecond = time.Now().Nanosecond()/1e6
	dayTimeArray := strings.Split(timeFormatString, " ")
	secondTimeArray := strings.Split(dayTimeArray[1], ":")
	secondint, _ := strconv.Atoi(secondTimeArray[2])

	return secondint
}

func GetPreOneMinTimeStamp() string {
	var naosecond = time.Now().Nanosecond() / 1e3
	timeFormatString := time.Now().Add(-time.Minute * 1).Format("2006-01-02 15:04:05")
	timeCorrect := fmt.Sprintf("%s.%d", timeFormatString, naosecond)
	return timeCorrect
}

func GetNowOneMinTimeStamp() string {
	time.Sleep(10 * time.Millisecond)
	var naosecond = time.Now().Nanosecond() / 1e3
	timeFormatString := time.Now().Format("2006-01-02 15:04:05")
	timeCorrect := fmt.Sprintf("%s.%d", timeFormatString, naosecond)
	return timeCorrect
}

func GetMinIntByTimeStamp(timeStamp string) (string, int) {
	dayTimeArray := strings.Split(timeStamp, " ")
	minTimeArrayPoint := strings.Split(dayTimeArray[1], ".")
	minTimeArray := strings.Split(minTimeArrayPoint[0], ":")
	hourInt, _ := strconv.Atoi(minTimeArray[0])
	minInt, _ := strconv.Atoi(minTimeArray[1])
	indexMinInt := hourInt*60 + minInt
	return dayTimeArray[0], indexMinInt
}

//时间戳转时间
func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

//时间转时间戳
func StrToUnix(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}

//构建根据时间戳  构建同一时间内第几个数据
func DevTimestampGenerateIndex(cli *clientv3.Client, createTimestamp string, requestTimeout time.Duration) string {
	//获取以该时间戳为前缀的元数据个数
	resp := GetDataPrefix(cli, createTimestamp, requestTimeout)
	curIndex := len(resp.Kvs) + 1
	ms := strings.Split(createTimestamp, ".")[1]
	strFormat := "%0" + strconv.Itoa(6-len(ms)) + "d"
	curIndexStr := fmt.Sprintf(strFormat, curIndex)
	curCreateTimestamp := createTimestamp + curIndexStr
	return curCreateTimestamp
}
