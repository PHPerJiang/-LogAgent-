package elasticsearch

import (
	"context"

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
func CreateIndex(indexStr string) (err error) {
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties":{
				"log":{
					"type":"text"
				}
			}
		}
	}`
	// defer client.Stop()
	ctx := context.Background()
	_, err = client.CreateIndex(indexStr).BodyString(mapping).Do(ctx)
	return
}

// IndexExists 确认索引是否存在
func IndexExists(indexStr string) (resp bool, err error) {
	ctx := context.Background()
	// defer client.Stop()
	resp, err = client.IndexExists(indexStr).Do(ctx)
	return
}
