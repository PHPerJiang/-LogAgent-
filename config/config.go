package config

// KafkaConf 配置
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
	MaxChan int    `ini:"maxchan"`
}

// TaillogCof 配置
type TaillogCof struct {
	FilePath string `ini:"filePath"`
}

// Etcd 配置
type Etcd struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"key"`
}

// ElasticSearch es配置
type ElasticSearch struct {
	Address string `ini:"address"`
}

// Conf 配置
type Conf struct {
	KafkaConf     `ini:"kafka"`
	TaillogCof    `ini:"taillog"`
	Etcd          `ini:"etcd"`
	ElasticSearch `ini:"elasticsearch"`
}
