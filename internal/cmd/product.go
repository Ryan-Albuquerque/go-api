package cmd

import (
	"database/sql"

	"github.com/Ryan-Albuquerque/go-api/internal/controller"
	"github.com/Ryan-Albuquerque/go-api/internal/middleware"
	"github.com/Ryan-Albuquerque/go-api/internal/repository"
	"github.com/Ryan-Albuquerque/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func ProductHandle(db *sql.DB, router *gin.Engine) {
	ProductRepository := repository.NewProductRepository(db)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	pRouter := router.Group("/products")
	pRouter.Use(middleware.AuthMiddleware())
	{
		pRouter.GET("/", ProductController.GetProducts)
		pRouter.POST("/", ProductController.CreateProduct)
		pRouter.GET("/:id", ProductController.GetProductByID)
		pRouter.PUT("/:id", ProductController.UpdateProduct)
		pRouter.DELETE("/:id", ProductController.DeleteProduct)
	}
}
