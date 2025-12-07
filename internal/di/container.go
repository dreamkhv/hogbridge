package di

import (
	"hog-bridge/internal/config"
	"hog-bridge/internal/handler"
	"hog-bridge/internal/mailer"
	"hog-bridge/internal/request"
)

type Container struct {
	MailgunHandler   *handler.MailHandler[request.MailgunMessage]
	MailchimpHandler *handler.MailHandler[request.SendgridMessage]
}

func Initialize(cfg config.Config) (*Container, error) {
	m, err := mailer.NewMailer(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUsername,
		cfg.MailPassword,
	)

	if err != nil {
		return nil, err
	}

	return &Container{
		MailgunHandler:   handler.NewMailgunHandler(m),
		MailchimpHandler: handler.NewSendgridHandler(m),
	}, nil
}
