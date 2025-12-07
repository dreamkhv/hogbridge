package mailer

import (
	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client *mail.Client
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

	return &Mailer{client: client}, nil
}

func (m *Mailer) Send(msg *mail.Msg) error {
	if err := m.client.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
