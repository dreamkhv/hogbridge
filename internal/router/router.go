package router

import (
	"hog-bridge/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *di.Container) *gin.Engine {
	r := gin.Default()

	r.POST("/v3/:domain/messages", c.MailgunHandler.Handle) // Mailgun
	r.POST("/v3/mail/send", c.MailchimpHandler.Handle)      // Sendgrid

	return r
}
