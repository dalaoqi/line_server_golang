package server

import (
	"line_server_golang/internal/utils/conf"
	"line_server_golang/internal/utils/db/mongo"
	"line_server_golang/internal/utils/line"
	"line_server_golang/router"
)

func Run() {
	conf.Init()
	line.Init()
	mongo.Init()
	router.Init()
}
