package config

import (
	"hog-bridge/internal/enum"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort   string        `envconfig:"SERVER_PORT" default:":8080"`
	MailType     enum.MailType `envconfig:"MAIT_TYPE" default:"mailgun"`
	MailHost     string        `envconfig:"MAIL_HOST" default:"localhost"`
	MailPort     int           `envconfig:"MAIL_PORT" default:"1025"`
	MailUsername string        `envconfig:"MAIL_USERNAME"`
	MailPassword string        `envconfig:"MAIL_PASSWORD"`
}

func LoadConfig() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
