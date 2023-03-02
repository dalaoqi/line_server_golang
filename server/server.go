package server

import (
	"line_server_golang/internal/utils/conf"
	"line_server_golang/internal/utils/db/mongo"
	"line_server_golang/internal/utils/line"
)

func Run() {
	conf.Init()
	line.Init()
	mongo.Init()
}
