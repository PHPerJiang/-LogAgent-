package taillog

import (
	"LogAgent/kafka"
	"log"
	"time"

	"github.com/hpcloud/tail"
)

var (
	tailsObj *tail.Tail
)

// TailTask 任务结构体
type TailTask struct {
	Path     string
	Topic    string
	Instance *tail.Tail
}

// NewTail 创建一个tail
func NewTail(path,topic string)(tailObj *TailTask){
	tailObj := &TailTask{
		Path: path,
		Topic: topic
	}
	tailObj.init()
	return
}

// init 初始化一个tail
func (t *TailTask)init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.Instance, err = tail.TailFile(t.Path, config)
	if err != nil{
		log.Printf("init tail failed %v", err)
		return 
	}
	go t.getTailLines()
}

//获取要收集的日志并发送到kafka
func (t *TailTask)getTailLines(){
	for {
		select{
		case line := <- t.Instance.Lines:
			kafka.SendMessag2Kafka(t.Topic, line.Text)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}