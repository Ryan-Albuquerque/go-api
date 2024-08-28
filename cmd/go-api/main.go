package main

import (
	"github.com/Ryan-Albuquerque/go-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{})

	ProductController := controller.NewProductController()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/products", ProductController.GetProducts)

	router.Run(":8080")
}
