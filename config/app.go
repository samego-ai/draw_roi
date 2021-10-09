package config

import "os"

var App app

type app struct {
	Protocol string // 应用通信协议
	Address  string // 应用通信地址
}

func init() {
	App.Protocol = os.Getenv("SG_APP_PROTOCOL")
	App.Address = os.Getenv("SG_APP_ADDRESS")
}
