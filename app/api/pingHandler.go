package api

import "github.com/gin-gonic/gin"

func (h *BaseHandler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
