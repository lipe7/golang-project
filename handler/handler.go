package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "POST opening",
	})
}

func ShowOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "opening",
	})
}

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "PUT opening",
	})
}

func ListOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "openings",
	})
}

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "DELETE opening",
	})
}
