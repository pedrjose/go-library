package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/pedrjose/go-library/db"
	"github.com/pedrjose/go-library/models"
	"github.com/pedrjose/go-library/services"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go library. How can we assist you?",
		})
	})

	router.POST("register-owner", func(context *gin.Context) {
		var newOwner models.Owner

		if err := context.ShouldBindJSON(&newOwner); err != nil {
			fmt.Println("Server Request Error!")
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Request.Context()
		err := services.SetOwner(ctx, newOwner)

		if err != nil {
			fmt.Println("Server Internal Settings Error!")
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Owner registered successfully!"})
	})

	router.GET("/find-owner", func(context *gin.Context) {
		var ownerEmail models.Owner

		ctx := context.Request.Context()
		owner := services.GetOwner(ctx, ownerEmail)

		fmt.Println(owner)
	})

	err := router.Run(":3000")
	if err != nil {
		return
	}
}
