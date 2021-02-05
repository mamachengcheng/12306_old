package static

import "gopkg.in/ini.v1"

const ConfFilePath = "conf/app/config.ini"

var (
	ServerAddress string
	GrpcAddress   string
	RMQAddress    string
)

func init() {

	cfg, _ := ini.Load(ConfFilePath)
	server := cfg.Section("server")
	rocketmq := cfg.Section("rocketmq")
	grpc := cfg.Section("grpc")

	ServerAddress = server.Key("http").String() + ":" + server.Key("port").String()
	GrpcAddress = grpc.Key("http").String() + ":" + grpc.Key("port").String()
	RMQAddress = rocketmq.Key("http").String() + ":" + rocketmq.Key("port").String()
}
