package main

import (
	"LogAgent/config"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(config.Conf)
)

func run() {
	for {
		select {
		case line := <-taillog.GetTailLines():
			kafka.SendMessag2Kafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

//logagent 入口程序
func main() {
	//1.加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed err: %v", err)
		return
	}

	//	2.初始化kafka链接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("kafka init failed,err:%v\n", err)
		return
	}
	fmt.Println("连接kakfa成功")
	//	3.收集日志文件
	err = taillog.Init(cfg.TaillogCof.FilePath)
	if err != nil {
		fmt.Printf("tail init failed,err:%v\n", err)
		return
	}
	fmt.Println("连接tail成功")

	//启动主程序
	run()
}
