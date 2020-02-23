package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

// Init 初始化kafka
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer(addrs, config)
	return
}

// SendMessag2Kafka 发送信息到kafka
func SendMessag2Kafka(topic string, message string) {
	msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder(message)}
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send message failed err : %v", err)
		return
	}
	fmt.Printf("pid :%v , offset :%v\n", pid, offset)
}
