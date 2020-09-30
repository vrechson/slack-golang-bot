package app

// Note that all storage and repository codes are empty. Will be necessary to create a new storage model

import (
	"fmt"
	"log"
	"time"

	"github.com/vrechson/slack-golang-bot/config"
	//models "github.com/vrechson/slack-golang-bot/database/models"
	repositories "github.com/vrechson/slack-golang-bot/database/repositories"
	storages "github.com/vrechson/slack-golang-bot/database/storages"
	botFactory "github.com/vrechson/slack-golang-bot/factory/bot"
	dbFactory "github.com/vrechson/slack-golang-bot/factory/database"
	modules "github.com/vrechson/slack-golang-bot/modules"
	"github.com/vrechson/slack-golang-bot/slack"
	//"github.com/vrechson/slack-golang-bot/utils"
)

// SlackBot is the main structure containing all data types and the bot itselff
type SlackBot struct {
	conf       *config.Config
	bot        *slack.Handler
	storage    *storages.Handler
	repository *repositories.Handler
	modules    *modules.Handler
}

// CreateApp setup every structure from SlackBot
func CreateApp(conf *config.Config) *SlackBot {

	// Init the bot
	bot, err := botFactory.BotFactory(conf)
	if err != nil {
		log.Fatal("Can't create slack bot: ", err)
	}

	// Init the Database
	db, err := dbFactory.DatabaseFactory(conf)
	if err != nil {
		log.Fatal("Couldn't create database: ", err)
	}

	// Init the database Storage
	storage, err := storages.CreateStorage(db)
	if err != nil {
		log.Fatal("Couldn't setup database: ", err)
	}

	// Init the database Repository
	repository, err := repositories.CreateRepository(db)
	if err != nil {
		log.Fatal("Couldn't setup database: ", err)
	}

	// Init the database Repository
	mod := modules.CreateModules()
	if err != nil {
		log.Fatal("Couldn't setup database: ", err)
	}

	fmt.Println("App            | App Initializated")
	return &SlackBot{conf, bot, storage, repository, mod}
}

// Start is the function that start the bot service
func (slack *SlackBot) Start() {

	// do something on startup

	// Event Loop. Soon will become a go routine and event loop will be in modules package
	go slack.bot.ListenAndResponse()

	// Future event loop
	slack.Handler()

}

func (slack *SlackBot) Handler() {

	for { //ever
		
		// do something

		// 10 second before restart
		time.Sleep(360000 * time.Millisecond)
	}

}
