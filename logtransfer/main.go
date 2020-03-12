package main

import (
	"LogAgent/config"
	"LogAgent/elasticsearch"
	"LogAgent/kafka"
	"log"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(config.Conf)
)

func main() {
	//  加载配置文件
	err := ini.MapTo(cfg, "../config/config.ini")
	if err != nil {
		log.Printf("load config failed err: %v", err)
		return
	}
	log.Println("load config success")
	//初始化es链接
	err = elasticsearch.Init(cfg.ElasticSearch.Address)
	if err != nil {
		log.Printf("init elastic failed ,%v", err)
		return
	}
	log.Println("init elastic success")
	//判断索引是否存在，不存在则创建索引
	//kafka数据入es
	err = kafka.ConsumeMessage(cfg.KafkaConf.Address, "test01")
	if err != nil {
		log.Printf("consmeMessage failed err :%err", err)
		return
	}
}
