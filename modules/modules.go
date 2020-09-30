package modules

import (
	"github.com/vrechson/slack-golang-bot/modules/amass"
)

// Handler is a struct ta handle all modules
type Handler struct {
	Sample     *sample.Handler
}

// CreateModules Setup all modules
func CreateModules() *Handler {
	s, _ := sample.CreateSample()

	return &Handler{s}
}
