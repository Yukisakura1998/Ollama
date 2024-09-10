package main

import (
	"Ollama/ollama_chatbox/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cores())
	r.Use(middleware.Router(r))

	err := r.Run()
	if err != nil {
		return
	}
}
