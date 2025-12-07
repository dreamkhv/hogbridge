package factory

import (
	"hog-bridge/internal/request"

	"github.com/wneessen/go-mail"
)

type MessageFactoryInterface interface {
	New(req request.MessageInterface) (*mail.Msg, error)
}
