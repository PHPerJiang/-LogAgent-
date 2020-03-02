package main

import (
	"LogAgent/config"
	"LogAgent/etcd"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(config.Conf)
)

//主程序
func run() {
	for {
		select {
		case line := <-taillog.GetTailLines():
			//发送信息到kafka
			kafka.SendMessag2Kafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

//logagent 入口程序
func main() {
	//  加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		log.Printf("load config failed err: %v", err)
		return
	}
	log.Println("配置文件加载成功")

	//	初始化kafka链接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		log.Printf("kafka init failed,err:%v\n", err)
		return
	}
	log.Println("连接kakfa成功")

	//  初始化etcd
	err = etcd.Init(cfg.Etcd.Address, time.Duration(cfg.Etcd.Timeout)*time.Millisecond)
	if err != nil {
		log.Printf("etcd init failed,err:%v\n", err)
		return
	}
	log.Println("etcd初始化成功")

	//测试连接
	// etcd.Put("/etcd", `[{"path":"/data/webroot/go/src/LogAgent/my.log","topic":"test01"},{"path":"/data/webroot/go/src/LogAgent/my.log","topic":"test02"}]`)

	logconf, err := etcd.Get("/etcd", time.Duration(cfg.Timeout)*time.Millisecond)
	if err != nil {
		log.Printf("get etcd data failed %v", err)
		return
	}
	//打印数据
	// for _, v := range logconf {
	// 	log.Printf("log conf %v", v)
	// }
	//	收集日志文件
	err = taillog.Init(cfg.TaillogCof.FilePath)
	if err != nil {
		log.Printf("tail init failed,err:%v\n", err)
		return
	}
	log.Println("连接tail成功")

	//启动主程序
	run()
}
