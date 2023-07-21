package model

type BootConfFlag struct {
	ServerPort string `env:"user_server_port" desc:"user系统http服务器监听端口" default:"9000"`
}
