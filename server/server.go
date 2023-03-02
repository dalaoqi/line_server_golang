package server

import (
	"line_server_golang/internal/utils/conf"
	"line_server_golang/internal/utils/line"
)

func Run() {
	conf.Init()
	line.Init()
}
