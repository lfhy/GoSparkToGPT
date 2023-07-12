package handler

import (
	"net/http"
	"sparktogpt/api"
	"sparktogpt/types"
	"sparktogpt/utils"

	"github.com/gin-gonic/gin"
)

func Completions(c *gin.Context) {
	// 校验请求
	var request types.CompletionsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "error": err})
		return
	}

	var chatdata string
	for _, cm := range request.Messages {
		chatdata += cm.Content
	}

	// 从星火大模型获取消息并回复
	resq := api.NewCompletionsResponse(api.ChatToSpark(request.Messages), utils.GetToken(chatdata))
	c.String(http.StatusOK, resq.String())
}
