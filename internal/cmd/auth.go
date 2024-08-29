package cmd

import (
	"database/sql"

	"github.com/Ryan-Albuquerque/go-api/internal/controller"
	"github.com/Ryan-Albuquerque/go-api/internal/repository"
	"github.com/Ryan-Albuquerque/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func AuthHandle(db *sql.DB, router *gin.Engine) {
	AuthRepository := repository.NewAuthRepository(db)

	AuthUseCase := usecase.NewAuthUseCase(AuthRepository)

	AuthController := controller.NewAuthController(AuthUseCase)

	aRouter := router.Group("/auth")
	{
		aRouter.POST("/login", AuthController.Login)
		aRouter.POST("/register", AuthController.Register)
	}
}
