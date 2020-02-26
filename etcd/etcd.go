package etcd

import (
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
)

// Init 初始化etcd
func Init(address string, timeout time.Duration) (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Printf("etcd connect failed, err: %v", err)
		return
	}
	defer cli.Close()
	return
}
