package main

import (
	"sparktogpt/conf"
	"sparktogpt/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/v1/chat/completions", handler.Completions)
	r.Run(":" + conf.Port)
}
