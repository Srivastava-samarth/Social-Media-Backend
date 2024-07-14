package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() *mongo.Client{
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading the .env file: ", err)
    }
    MONGODB_URI := os.Getenv("MONGODB_URI")
    if MONGODB_URI == "" {
        log.Fatal("MONGODB_URI is not set in the environment variables")
    }
	clientOptions := options.Client().ApplyURI(MONGODB_URI);
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