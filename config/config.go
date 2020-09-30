package config

import "github.com/kelseyhightower/envconfig"

// Config is a structure to control environment variables
type Config struct {
	Token     string `envconfig:"TOKEN" required:"true"`
	BotID     string `envconfig:"BOTID" required:"true"`
	ChannelID string `envconfig:"CHANNELID" required:"true"`
	MongoDB   string `envconfig:"MONGODB" required:"true"`
	Debug     bool   `envconfig:"DEBUG"`
	Version   string
}

// Setup will load Config into environment variables
func Setup() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return &c, err
	}

	c.Version = "0.1.5"
	return &c, nil
}
