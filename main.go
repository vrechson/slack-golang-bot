// Bot is the slack golang bot
// hi bot

package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	app "github.com/vrechson/slack-golang-bot/app"
	"github.com/vrechson/slack-golang-bot/config"
)

func main() {
	c, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	a := app.CreateApp(c)
	a.Start()

}
