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

// TailTask tail任务
type TailTask struct {
	Path     string     //要收集的日志路径
	Topic    string     //要发给kafka哪个topic
	Instance *tail.Tail //tail实例
}

// NewTail 创建一个tail实例
func NewTail(path, topic string) (tailObj *TailTask) {
	tailObj := &TailTask{
		Path:  path,
		Topic: topic,
	}
	tailObj.init()
	return
}

// Init 初始化tail
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.Instance, err = tail.TailFile(t.Path, config)
	if err != nil {
		log.Printf("init tail failed %v", err)
	}
	t.run()
}

// 开始读取日志发送到kafka
func (t *TailTask) run() {
	for {
		select {
		case line := <-t.Instance.Lines:
			kafka.SendMessag2Kafka(t.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
