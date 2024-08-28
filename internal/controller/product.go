package controller

import (
	"net/http"

	"github.com/Ryan-Albuquerque/go-api/internal/model"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Product 1",
			Price: 100.00,
		},
		{
			ID:    2,
			Name:  "Product 2",
			Price: 200.00,
		},
	}
	c.JSON(http.StatusOK, products)
}
