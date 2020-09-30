package slack

import (
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/vrechson/slack-golang-bot/slack/commands"
)

// Handler struct handles all bot actions
type Handler struct {
	Client    *slack.Client
	BotID     string
	ChannelID string
	Debug     bool
	Version   string
}

// ListenAndResponse listen and handles slack events
func (H *Handler) ListenAndResponse() {
	rtm := H.Client.NewRTM()

	// Start listening events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := H.handleMessageEvent(ev); err != nil {
				log.Printf("Slack | Failed to handle message: %s", err)
			}
		}
	}
}

func (H *Handler) handleMessageEvent(ev *slack.MessageEvent) error {

	// Ignore what isn't comming from it's channel
	if ev.Channel != H.ChannelID {
		return nil
	}

	// Only answer mention to bot. May be changed
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", H.BotID)) {
		return nil
	}

	// Parse message >>>> NEED TO CHANGE, ANY UNSPECTED MESSAGE AND THE BOT BREAK
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	command := m[0]
	if len(m) == 0 {
		return fmt.Errorf("Invalid Message.")
	}

	switch command {

	// Showing help menu
	case commands.Help:
		H.sendMessage("Read The Fucking Manual")

	// Showing help menu
	case commands.Ping:
		H.sendMessage("pong")

	// Showing verion to debug
	case commands.Version:
		H.sendMessage("Actually running version " + H.Version)

	}

	return nil

}

func (H *Handler) sendMessage(msg string) {
	params := slack.NewPostMessageParameters()
	params.LinkNames = 1

	msgOption := slack.MsgOptionCompose(
		slack.MsgOptionAsUser(true),
		slack.MsgOptionText(msg, false),
		slack.MsgOptionPostMessageParameters(params),
	)

	if _, _, err := H.Client.PostMessage(H.ChannelID, msgOption); err != nil {
		fmt.Errorf("failed to post message: %s", err)
	}

}
