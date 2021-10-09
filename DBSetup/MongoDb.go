package DBSetup

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://naveen:jI5jrhnXHI8ibyQw@cluster1.ezz33.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	return client
}
