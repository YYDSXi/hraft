package utils

import (
	"strings"
)

func TimeTrue(timeStamp string)  bool {
	if len(timeStamp) != 23{
		return false
	}
	if !strings.Contains(timeStamp," "){
		return false
	}
	timeArr := strings.Split(timeStamp," ")
	if len(timeArr) != 2 {
		return false
	}

	if !strings.Contains(timeArr[0],"-"){
		return false
	}
	ymdArr := strings.Split(timeArr[0],"-")
	if len(ymdArr) != 3 {
		return false
	}
	if len(ymdArr[0]) != 4 || len(ymdArr[1]) != 2 || len(ymdArr[2]) != 2 {
		return false
	}
	if !strings.Contains(timeArr[1],".") || !strings.Contains(timeArr[1],":"){
		return false
	}
	hmsArr := strings.Split(timeArr[1],".")
	if len(hmsArr) != 2{
		return false
	}
	if len(hmsArr[0]) != 8 || len(hmsArr[1]) != 3 {
		return false
	}
	return true
}
