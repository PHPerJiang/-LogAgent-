package elasticsearch

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
)

// Init 初始化
func Init(address string) (err error) {
	client, err = elastic.NewClient(elastic.SetURL(address))
	return
}

// CreateIndex 创建索引
func CreateIndex(indexStr string) {
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties":{
				"log":{
					"type":"text"
				},
				"age":{
					"type":"integer"
				}
			}
		}
	}`
	ctx := context.Background()
	_, err := client.CreateIndex(indexStr).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Printf("create index failed %v", err)
		return
	}
	log.Println("create index success")
	defer client.Stop()
}
