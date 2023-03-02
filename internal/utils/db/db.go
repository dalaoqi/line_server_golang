package db

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoClient *mongo.Client
	LineEvent   *mongo.Collection
)
