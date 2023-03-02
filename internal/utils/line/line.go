package line

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

var LineClient *linebot.Client

func Init() {
	var err error
	LineClient, err = linebot.New(viper.GetString("app.lineSecret"), viper.GetString("app.lineAccessToken"))
	if err != nil {
		log.Panic(err)
	}
}
