package factory

import (
	"hog-bridge/internal/request"

	"github.com/wneessen/go-mail"
)

type MailgunMessageFactory struct{}

func (f *MailgunMessageFactory) New(r request.MessageInterface) (*mail.Msg, error) {
	req := r.(request.MailgunMessage)
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
		defer func() { _ = file.Close() }()

		err = msg.AttachReader(req.Attachment.Filename, file)
		if err != nil {
			return nil, err
		}
	}

	return msg, nil
}
