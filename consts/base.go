package consts

import "time"

const (
	// 服务信息
	ReadTimeout   = 15 * time.Second
	WriteTimeout  = 15 * time.Second
	IdleTimeout   = 30 * time.Second
	ServerLogFile = "server_log.log"

	Salt = "~@log@~"
)
