package database

import (
	"context"
	"fmt"

	models "github.com/vrechson/slack-golang-bot/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// All stuff necessary to write to database