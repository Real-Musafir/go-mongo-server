package main

import (
	"fmt"

	"github.com/Real-Musafir/bookshop/config"
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

	port := config.GetEnvProperty("port")

	app.Run(fmt.Sprintf(":%s", port))
}