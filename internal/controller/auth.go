package controller

import (
	"github.com/Ryan-Albuquerque/go-api/internal/model"
	"github.com/Ryan-Albuquerque/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auc *usecase.AuthUseCase
}

func NewAuthController(auc *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		auc: auc,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var user model.Login

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	response, err := ac.auc.Login(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Login success",
		"data":    response,
	})

}

func (ac *AuthController) Register(c *gin.Context) {
	var user model.Register

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	response, err := ac.auc.Register(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Register success",
		"data":    response,
	})

}
