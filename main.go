package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/api/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "messaage": "Implement Google OAuth2 in Golang"})
	})

	log.Fatal(server.Run(":" + "8000"))
}
