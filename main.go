package main

import (
	"LogAgent/kafka"
	"log"
)

//logagent 入口程序
func main() {
	//	1.初始化kafka链接
	//	2.收集日志文件
	err := kafka.Init()
	if err != nil {
		log.Printf("kafka init failed,err:%v", err)
		return
	}
}
