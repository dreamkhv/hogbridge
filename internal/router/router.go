package router

import (
	"hog-bridge/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.MailgunHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/v3/:domain/messages", h.Handle)

	return r
}
