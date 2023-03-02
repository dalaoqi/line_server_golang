package mongo

import (
	"context"
	"line_server_golang/internal/utils/db"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	var err error
	db.MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("db.mongo.url")))
	if err != nil {
		log.Fatal(err)
	}
	err = db.MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	dbName := viper.GetString("db.mongo.name")
	collectionName := viper.GetString("db.mongo.lineEvent.name")
	db.LineEvent = db.MongoClient.Database(dbName).Collection(collectionName)
}
