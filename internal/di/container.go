package di

import (
	"hog-bridge/internal/config"
	"hog-bridge/internal/handler"
	"hog-bridge/internal/mailer"
)

func Initialize(cfg config.Config) (*handler.MailgunHandler, error) {
	m, err := mailer.NewMailer(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUsername,
		cfg.MailPassword,
	)

	if err != nil {
		return nil, err
	}

	return handler.NewMailgunHandler(m), nil
}
