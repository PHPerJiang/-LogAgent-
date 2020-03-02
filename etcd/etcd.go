package etcd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
)

// LogConf 从etcd中取出的数据的结构体
type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// Init 初始化etcd
func Init(address string, timeout time.Duration) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{address},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Printf("etcd connect failed, err: %v", err)
		return
	}
	return
}

//Put 向etcd中推数据
func Put(etcdK string, etcdV string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	_, err := client.KV.Put(ctx, etcdK, etcdV)
	cancel()
	if err != nil {
		log.Printf("Put etcd error %v", err)
		return
	}
}

// Get 获取etcd中的配置
func Get(etcK string, timeout time.Duration) (logconf []*LogConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := client.Get(ctx, etcK)
	cancel()
	if err != nil {
		log.Printf("Get etcd data failed %v", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logconf)
		if err != nil {
			log.Printf("unmarshal failed %c", err)
			return
		}
	}
	return
}
