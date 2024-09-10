package Controller

import (
	"Ollama/ollama_chatbox/Model"
	"Ollama/ollama_chatbox/util"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"net/http"
)

type Chat struct {
}

func (chat *Chat) DoChat(c *gin.Context) {
	var body Model.Chat

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	prompt := util.CreatePrompt()

	data := map[string]any{
		"text": body.Text,
	}

	msg, err := prompt.FormatMessages(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	content := []llms.MessageContent{
		llms.TextParts(msg[0].GetType(), msg[0].GetContent()),
		llms.TextParts(msg[1].GetType(), msg[1].GetContent()),
	}

	llm := util.CreateOllama(c, "qwen")

	resp, err := llm.GenerateContent(context.Background(), content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    resp.Choices[0].Content,
	})

}
