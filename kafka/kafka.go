package kafka

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
	// Logchan log信息
	Logchan chan *LogData
)

// LogData 日志信息结构体
type LogData struct {
	Topic string
	Data  string
}

// Init 初始化kafka
func Init(addrs []string, maxChan int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer(addrs, config)
	//初始化日志通道
	Logchan = make(chan *LogData, maxChan)
	//启动一个协程从通道中取日志并发送给kafka
	go getLogByChan()
	return
}

// 从通道中取出log并发送到kafka
func getLogByChan() {
	for {
		select {
		case logInfo := <-Logchan:
			SendMessag2Kafka(logInfo.Topic, logInfo.Data)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// SendMessag2Kafka 发送信息到kafka
func SendMessag2Kafka(topic string, message string) {
	msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder(message)}
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		log.Printf("send message failed err : %v", err)
		return
	}
	log.Printf("pid :%v , offset :%v\n", pid, offset)
}
