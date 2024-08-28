package controller

import (
	"net/http"
	"strconv"

	"github.com/Ryan-Albuquerque/go-api/internal/model"
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

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product model.Product
	err := c.BindJSON(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = pc.puc.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}
	product, err := pc.puc.GetProductByID(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
