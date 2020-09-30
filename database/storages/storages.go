package database

import (
	//"context"
	"fmt"

	//models "github.com/vrechson/slack-golang-bot/database/models"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// All stuff necessary to write to database
type Handler struct {
	DB          *mongo.Client
	// add here all models
}

// CreateStorage is a function that setup database storages
func CreateStorage(c *mongo.Client) (*Handler, error) {
	fmt.Println("Database       | [Storages] Instatiated")
	return &Handler{
		DB:          c,
	}, nil
}