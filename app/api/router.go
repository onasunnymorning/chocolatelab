package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, h *BaseHandler) {
	r.GET("/ping", h.PingHandler)
}
