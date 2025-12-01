package mailer

import (
	"hog-bridge/internal/factory"
	"hog-bridge/internal/request"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client  *mail.Client
	factory *factory.MessageFactory
}

func NewMailer(host string, port int, username, password string) (*Mailer, error) {
	client, err := mail.NewClient(
		host,
		mail.WithPort(port),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTLSPolicy(mail.TLSOpportunistic),
	)
	if err != nil {
		return nil, err
	}

	return &Mailer{
		client:  client,
		factory: &factory.MessageFactory{},
	}, nil
}

func (m *Mailer) Send(req request.MailgunRequest) error {
	msg, err := m.factory.NewMailgun(req)
	if err != nil {
		return err
	}

	if err := m.client.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
