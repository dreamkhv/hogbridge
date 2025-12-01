package factory

import (
	"errors"
	"hog-bridge/internal/enum"
	"hog-bridge/internal/request"

	"github.com/wneessen/go-mail"
)

type MessageFactory struct{}

func (f *MessageFactory) New(mailType enum.MailType, req request.MessageRequest) (*mail.Msg, error) {
	switch mailType {
	case enum.Mailgun:
		r, _ := req.(request.MailgunMessageRequest)
		return f.NewMailgun(r)
	default:
		return nil, errors.New("unknown mail type")
	}
}

func (f *MessageFactory) NewMailgun(req request.MailgunMessageRequest) (*mail.Msg, error) {
	msg := mail.NewMsg()

	if err := msg.From(req.From); err != nil {
		return nil, err
	}

	if err := msg.To(req.To); err != nil {
		return nil, err
	}

	msg.Subject(req.Subject)

	if req.Text != "" {
		msg.SetBodyString(mail.TypeTextPlain, req.Text)
	} else if req.Html != "" {
		msg.SetBodyString(mail.TypeTextHTML, req.Html)
	}

	if req.Attachment != nil {
		file, err := req.Attachment.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = msg.AttachReader(req.Attachment.Filename, file)
		if err != nil {
			return nil, err
		}
	}

	return msg, nil
}
