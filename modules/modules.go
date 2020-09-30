package modules

import (
	"github.com/vrechson/slack-golang-bot/modules/sample"
)

// Handler is a struct ta handle all modules
type Handler struct {
	Sample     *sample.Handler
}

// CreateModules Setup all modules
func CreateModules() *Handler {
	s := sample.CreateSample()

	return &Handler{s}
}
