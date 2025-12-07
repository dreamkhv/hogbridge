package handler

import (
	"hog-bridge/internal/mailer"
	"hog-bridge/internal/request"
	"hog-bridge/internal/strategy"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MailHandler[T request.MessageInterface] struct {
	mailer    *mailer.Mailer
	strategy  *strategy.MessageFactoryStrategy
	validator *validator.Validate
}

func NewMailgunHandler(mailer *mailer.Mailer) *MailHandler[request.MailgunMessage] {
	return &MailHandler[request.MailgunMessage]{
		mailer:    mailer,
		strategy:  &strategy.MessageFactoryStrategy{},
		validator: validator.New(),
	}
}

func NewSendgridHandler(mailer *mailer.Mailer) *MailHandler[request.SendgridMessage] {
	return &MailHandler[request.SendgridMessage]{
		mailer:    mailer,
		strategy:  &strategy.MessageFactoryStrategy{},
		validator: validator.New(),
	}
}

func (h *MailHandler[T]) Handle(c *gin.Context) {
	var req T

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	factory, err := h.strategy.New(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg, err := factory.New(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.mailer.Send(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
