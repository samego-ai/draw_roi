package config

import "os"

var Server server

type server struct {
	Protocol string // 应用通信协议
	Address  string // 应用通信地址
}

func init() {
	Server.Protocol = os.Getenv("SG_SERVER_PROTOCOL")
	Server.Address = os.Getenv("SG_SERVER_ADDRESS")
}
