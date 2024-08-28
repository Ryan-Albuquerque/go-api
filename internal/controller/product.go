package controller

import (
	"net/http"

	"github.com/Ryan-Albuquerque/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	puc *usecase.ProductUseCase
}

func NewProductController(puc *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		puc: puc,
	}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.puc.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
