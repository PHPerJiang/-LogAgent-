module LogAgent

go 1.13

require (
	github.com/Shopify/sarama v1.26.1
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/olivere/elastic/v7 v7.0.12
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.14.0 // indirect
	google.golang.org/grpc v1.26.0
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect

)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
