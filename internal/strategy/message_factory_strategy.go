package strategy

import (
	"errors"
	"hog-bridge/internal/factory"
	"hog-bridge/internal/request"
)

type MessageFactoryStrategy struct{}

func (f *MessageFactoryStrategy) New(req request.MessageInterface) (factory.MessageFactoryInterface, error) {
	switch req.(type) {
	case request.MailgunMessage:
		return &factory.MailgunMessageFactory{}, nil
	case request.SendgridMessage:
		return &factory.SendgridMessageFactory{}, nil
	default:
		return nil, errors.New("unknown mail type")
	}
}
