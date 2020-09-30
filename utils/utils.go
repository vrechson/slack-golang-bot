package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	models "github.com/vrechson/slack-golang-bot/database/models"
	repositories "github.com/vrechson/slack-golang-bot/database/repositories"
	storages "github.com/vrechson/slack-golang-bot/database/storages"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/yaml.v2"
)

// Auxiliar functions to be used anywhere
