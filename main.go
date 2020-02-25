package main

import (
	"LogAgent/config"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
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
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("etcd connect failed, err: %v", err)
		return
	}
	defer cli.Close()
	fmt.Println("etcd connection success!")
	os.Exit(1)

	//1.加载配置文件
	err = ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		log.Printf("load config failed err: %v", err)
		return
	}

	//	2.初始化kafka链接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		log.Printf("kafka init failed,err:%v\n", err)
		return
	}

	log.Println("连接kakfa成功")
	//	3.收集日志文件
	err = taillog.Init(cfg.TaillogCof.FilePath)
	if err != nil {
		log.Printf("tail init failed,err:%v\n", err)
		return
	}
	log.Println("连接tail成功")

	//启动主程序
	run()
}
