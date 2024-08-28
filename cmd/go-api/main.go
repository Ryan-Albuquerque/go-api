package main

import (
	"github.com/Ryan-Albuquerque/go-api/internal/controller"
	"github.com/Ryan-Albuquerque/go-api/internal/infra/db"
	"github.com/Ryan-Albuquerque/go-api/internal/repository"
	"github.com/Ryan-Albuquerque/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{})

	connected, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(connected)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/products", ProductController.GetProducts)

	router.Run(":8080")
}
