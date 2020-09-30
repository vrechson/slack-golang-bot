package database

import (
	//"context"
	"fmt"
	//"log"

	//models "github.com/vrechson/slack-golang-bot/database/models"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

// Handler is the structure containing all models to read data from database
type Handler struct {
	DB          *mongo.Client
	// add here all models
}

// CreateRepository is a function to setup repositories
func CreateRepository(c *mongo.Client) (*Handler, error) {
	fmt.Println("Database       | [Repositories] Instatiated")
	return &Handler{
		DB:          c,
	}, nil
}
// All stuff necessary to read from database
