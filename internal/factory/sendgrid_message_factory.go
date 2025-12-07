package factory

import (
	"hog-bridge/internal/request"

	"github.com/wneessen/go-mail"
)

type SendgridMessageFactory struct{}

func (f *SendgridMessageFactory) New(r request.MessageInterface) (*mail.Msg, error) {
	req := r.(request.SendgridMessage)
	msg := mail.NewMsg()

	if err := msg.To(req.Personalizations[0].To[0].Email); err != nil {
		return nil, err
	}

	if err := msg.From(req.From.Email); err != nil {
		return nil, err
	}

	msg.Subject(req.Subject)
	msg.SetBodyString(mail.ContentType(req.Content[0].Type), req.Content[0].Type)

	return msg, nil
}
