package model

type BootConfFlag struct {
	ServerPort string `env:"gm_server_port" desc:"gm系统http服务器监听端口" default:"9001"`
}
