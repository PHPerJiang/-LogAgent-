package taillog

import (
	"LogAgent/etcd"
	"log"
)

var logMgr *tailLogManger

type tailLogManger struct {
	LogConf []*etcd.LogConf
}

// Init 初始化tail管理
func Init(logConf []*etcd.LogConf) {
	logMgr = &tailLogManger{
		LogConf: logConf,
	}
	//根据获取的配置创建tail
	for _, conf := range logMgr.LogConf {
		NewTail(conf.Path, conf.Topic)
		log.Printf("new tail create success path:%s , topic：%s", conf.Path, conf.Topic)
	}
}
