package main

import (
	"github.com/Ryan-Albuquerque/go-api/internal/cmd"
	"github.com/Ryan-Albuquerque/go-api/internal/infra/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{})

	connected, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	cmd.ProductHandle(connected, router)

	router.Run(":8000")
}
