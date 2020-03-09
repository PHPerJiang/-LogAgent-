package main

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
)

// Init 初始化
func Init() (err error) {
	client, err = elastic.NewClient(elastic.SetURL("http://172.81.251.154:9200"))
	return
}

// CreateIndex 创建索引
func CreateIndex() {
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties":{
				"name":{
					"type":"keyword"
				},
				"age":{
					"type":"integer"
				}
			}
		}
	}`
	ctx := context.Background()
	_, err := client.CreateIndex("student").BodyString(mapping).Do(ctx)
	if err != nil {
		log.Printf("create index failed %v", err)
		return
	}
	log.Println("create index success")
	defer client.Stop()
}
