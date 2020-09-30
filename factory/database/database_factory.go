package factory

import (
	"context"
	"fmt"
	"log"

	"github.com/vrechson/slack-golang-bot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseFactory is a function that initializes the database
func DatabaseFactory(c *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(c.MongoDB)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Factory        | Connected to MongoDB")

	return client, nil
}
