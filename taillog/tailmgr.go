package taillog

import (
	"LogAgent/etcd"
	"fmt"
	"log"
	"time"
)

var (
	logMgr *tailLogManger
)

type tailLogManger struct {
	LogConf   []*etcd.LogConf
	taskMap   map[string]*TailTask
	newConfCh chan []*etcd.LogConf
}

// Init 初始化tail管理
func Init(logConf []*etcd.LogConf) {
	logMgr = &tailLogManger{
		LogConf:   logConf,
		taskMap:   make(map[string]*TailTask, 16),
		newConfCh: make(chan []*etcd.LogConf),
	}
	//根据获取的配置创建tail
	for _, conf := range logMgr.LogConf {
		tailObj := NewTail(conf.Path, conf.Topic)
		mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
		logMgr.taskMap[mk] = tailObj
		log.Printf("new tail create success path:%s , topic：%s", conf.Path, conf.Topic)
	}
	//读配置通道，有新配置进来则打印
	go logMgr.handleNewConf()
}

//处理配置通道的数据
func (t *tailLogManger) handleNewConf() {
	for {
		select {
		case newConf := <-logMgr.newConfCh:
			log.Printf("program get new conf %v", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 向外暴露一个只能写入新配置的通道
func NewConfChan() chan<- []*etcd.LogConf {
	return logMgr.newConfCh
}
