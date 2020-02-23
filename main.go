package main

import (
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"time"
)

func run() {
	for {
		select {
		case line := <-taillog.GetTailLines():
			kafka.SendMessag2Kafka("LogAgentLog", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

//logagent 入口程序
func main() {
	//	1.初始化kafka链接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("kafka init failed,err:%v\n", err)
		return
	}
	fmt.Println("连接kakfa成功")
	//	2.收集日志文件
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("tail init failed,err:%v\n", err)
		return
	}
	fmt.Println("连接tail成功")

	//启动主程序
	run()
}
