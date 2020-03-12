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
	logStr := "log"
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
	//判断索引是否存在
	resp, err := elasticsearch.IndexExists(logStr)
	if err != nil {
		log.Printf("get index exists failed %v", err)
		return
	}
	//不存在则创建
	if !resp {
		err = elasticsearch.CreateIndex(logStr)
		if err != nil {
			log.Printf("create index failed%v", err)
			return
		}
		log.Printf("索引 %s 不存在, 已自动创建成功", logStr)
	}
	//kafka数据入es
	err = kafka.ConsumeMessage(cfg.KafkaConf.Address, "test01")
	if err != nil {
		log.Printf("consmeMessage failed err :%err", err)
		return
	}
}
