package api

import (
	"context"
	"fmt"
	"line_server_golang/dto"
	"line_server_golang/internal/utils/db"
	"line_server_golang/internal/utils/line"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/bson"
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
		if event.Type == linebot.EventTypeFollow {
			_, err = line.LineClient.BroadcastMessage(linebot.NewTextMessage(fmt.Sprintf("User ID: %v", event.Source.UserID))).Do()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				log.Printf("broadcast(push ID) error: %v", err.Error())
				return
			}
		}
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

func GetUserMessages(c *gin.Context) {
	var lineEvent dto.LineEvent
	var res = make([]dto.LineEvent, 0)
	userId := c.Param("userId")

	cur, err := db.LineEvent.Find(context.Background(), bson.M{"source.userid": userId, "type": "message"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Printf("get users' messages error: %v", err.Error())
		return
	}

	for cur.Next(context.Background()) {
		if err := cur.Decode(&lineEvent); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			log.Printf("get users' messages error: %v", err.Error())
			return
		}
		res = append(res, lineEvent)
	}
	c.JSON(http.StatusOK, res)
}
