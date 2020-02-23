package taillog

import (
	"github.com/hpcloud/tail"
)

var (
	tailsObj *tail.Tail
)

// Init 初始化tail
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailsObj, err = tail.TailFile(fileName, config)
	return
}

// GetTailLines 获取tailLines
func GetTailLines() <-chan *tail.Line {
	return tailsObj.Lines
}
