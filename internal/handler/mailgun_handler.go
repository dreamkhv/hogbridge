package handler

import (
	"hog-bridge/internal/mailer"
	"hog-bridge/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MailgunHandler struct {
	mailer    *mailer.Mailer
	validator *validator.Validate
}

func NewMailgunHandler(mailer *mailer.Mailer) *MailgunHandler {
	return &MailgunHandler{
		mailer:    mailer,
		validator: validator.New(),
	}
}

func (h *MailgunHandler) Handle(c *gin.Context) {
	var req request.MailgunRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.mailer.Send(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      "<mock@mg.example.com>",
		"message": "Queued. Thank you.",
	})
}
