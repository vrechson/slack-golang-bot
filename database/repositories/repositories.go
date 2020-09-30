package database

import (
	"context"
	"fmt"
	"log"

	models "github.com/vrechson/slack-golang-bot/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handler is the structure containing all models to read data from database
type Handler struct {
	DB          *mongo.Client
	Program     models.ProgramModel
	WebScope    models.WebScopeModel
	MobileScope models.MobileScopeModel
	SourceScope models.SourceScopeModel
	Domain      models.DomainModel
}

// All stuff necessary to read from database
