package utils

import (
	"hraft/dataStruct"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf dataStruct.GlobalConfig

func InitAndGetConf(ConfPath string) dataStruct.GlobalConfig {
	data, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		log.Fatal(err)
	}

	var config dataStruct.GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	Conf = config
	return config
}

// func main1(){
// 	conf:=GetConf()

// 	fmt.Println(reflect.TypeOf(conf.Consense.AsyncTask.Async_node_credible.Blockheader.Interval))
// 	fmt.Println(reflect.TypeOf(conf.Consense.AsyncTask.Async_node_credible.Blockheader.Interval))
// 	fmt.Println(reflect.TypeOf(conf.Consense.AsyncTask.Async_node_credible.Blockheader.Interval))
// 	fmt.Println(reflect.TypeOf(conf.Consense.AsyncTask.Async_node_credible.Blockheader.Interval))
// 	fmt.Println(reflect.TypeOf(conf.Consense.AsyncTask.Async_node_credible.Blockheader.Interval))
// }
