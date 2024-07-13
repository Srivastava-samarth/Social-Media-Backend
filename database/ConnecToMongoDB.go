package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() *mongo.Client{
	mongoUri := os.Getenv("MONGO_URI");
	clientOptions := options.Client().ApplyURI(mongoUri);
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel();
	client,err := mongo.Connect(ctx,clientOptions);
	if err!=nil{
		log.Fatal(err)
	}
	return client;
}

func MongoCollection(colltionName string) *mongo.Collection{
	return client.Database("social-media-backend").Collection(colltionName);
}