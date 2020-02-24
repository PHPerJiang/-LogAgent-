package config

// KafkaConf 配置
type KafkaConf struct{
	Address string `ini:"address"`
	Topic 	string `ini:"topic"`
}
// TaillogCof 配置
type TaillogCof struct{
	FilePath string `ini:"filePath"`
}

// Conf 配置
type Conf struct{
	KafkaConf	`ini:"kafka"`
	TaillogCof	`ini:"taillog"`
}
