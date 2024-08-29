package usecase

import (
	"fmt"
	"regexp"

	"github.com/Ryan-Albuquerque/go-api/internal/model"
	"github.com/Ryan-Albuquerque/go-api/internal/repository"
	"github.com/Ryan-Albuquerque/go-api/internal/utils"
)

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

type AuthUseCase struct {
	ar *repository.AuthRepository
}

func NewAuthUseCase(ar *repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		ar: ar,
	}
}

func (au *AuthUseCase) Register(user *model.Register) (model.RegisterResponse, error) {
	if err := validateEmail(user.Email); err != nil {
		return model.RegisterResponse{}, err
	}
	return au.ar.Register(user)
}

func validateEmail(email string) error {
	re := regexp.MustCompile(emailPattern)

	if !re.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func (au *AuthUseCase) Login(user *model.Login) (model.LoginResponse, error) {
	err := au.ar.Login(user)

	if err != nil {
		return model.LoginResponse{}, err
	}

	token, err := utils.Generate(user.Username)

	if err != nil {
		return model.LoginResponse{}, err
	}
	return model.LoginResponse{
		Token: token,
	}, nil
}
