package elasticsearch

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
	// ElasticCh 写入es的通道
	ElasticCh chan *LogInfo
	indexStr  string
	typeStr   string
)

// LogInfo 入es的数据
type LogInfo struct {
	log  string `json:"log"`
	time string `json:"time"`
}

// Init 初始化
func Init(address, indexStrParam, typeStrParam string, size int) (err error) {
	client, err = elastic.NewClient(elastic.SetURL(address))
	indexStr = indexStrParam
	typeStr = typeStrParam
	ElasticCh = make(chan *LogInfo, size)
	isExists, err := indexExists(indexStr)
	if err != nil {
		log.Printf("indexexists failed %v", err)
		return
	}
	if !isExists {
		err = createIndex(indexStr)
		if err != nil {
			log.Printf("create index failed %v", err)
			return
		}
	}
	return
}

//创建索引
func createIndex(indexStr string) (err error) {
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
				"time":{
					"type": "date",
          			"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
				}
			}
		}
	}`
	// defer client.Stop()
	ctx := context.Background()
	_, err = client.CreateIndex(indexStr).BodyString(mapping).Do(ctx)
	return
}

//确认索引是否存在
func indexExists(indexStr string) (resp bool, err error) {
	ctx := context.Background()
	// defer client.Stop()
	resp, err = client.IndexExists(indexStr).Do(ctx)
	return
}

// SendMessag2Elastic 发送数据到es
func SendMessag2Elastic() {
	for {
		select {
		case logitem := <-ElasticCh:
			id := time.Now().UnixNano()
			_, err := client.Index().Index(indexStr).Type(typeStr).Id(strconv.Itoa(int(id))).BodyJson(logitem).Do(context.Background())
			if err != nil {
				log.Printf("insert log failed %v err", err)
				return
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
