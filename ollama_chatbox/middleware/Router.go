package middleware

import (
	"Ollama/ollama_chatbox/Controller"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) gin.HandlerFunc {

	chat := Controller.Chat{}
	mainGroup := r.Group("/")
	return func(c *gin.Context) {
		mainGroup.POST("/chat", chat.DoChat)
	}
}
