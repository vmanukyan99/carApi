package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	credential := options.Credential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB_NAME")),
	}
}
