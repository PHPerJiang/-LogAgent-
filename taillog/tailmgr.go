package taillog

import "LogAgent/etcd"

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
	for _, conf := range logMgr {
		NewTail(conf.Path, conf.Topic)
	}
}
