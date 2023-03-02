package api

import (
	"context"
	"line_server_golang/internal/utils/db"
	"line_server_golang/internal/utils/line"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type broadcastBody struct {
	Message string `json:"message" binding:"required"`
}

func Webhook(c *gin.Context) {
	events, err := line.LineClient.ParseRequest(c.Request)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, event := range events {
		_, err = db.LineEvent.InsertOne(context.Background(), event)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			log.Printf("save db error: %v", err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func Broadcast(c *gin.Context) {
	body := broadcastBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Printf("read request error: %v", err.Error())
		return
	}
	_, err = line.LineClient.BroadcastMessage(linebot.NewTextMessage(body.Message)).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Printf("broadcast error: %v", err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
