package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.Use(gin.Recovery())
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is running fine!",
			"statusCode":200,
		})
	})
	fmt.Println("Go app started successfully")

	app.Run()
}