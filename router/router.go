package router

import (
	"fmt"
	"line_server_golang/router/api"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Router = gin.Default()

func Init() {
	r := gin.Default()
	r.POST("/callback", api.Webhook)
	port := fmt.Sprintf(":%d", viper.GetInt64("app.port"))
	err := r.Run(port)

	if err != nil {
		log.Panicf("router init failed: %v", err.Error())
	}
}
