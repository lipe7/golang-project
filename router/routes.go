package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api")

	v1.GET("opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "opening",
		})
	})
}
