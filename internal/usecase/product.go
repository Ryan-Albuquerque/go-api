package usecase

import (
	"github.com/Ryan-Albuquerque/go-api/internal/model"
	"github.com/Ryan-Albuquerque/go-api/internal/repository"
)

type ProductUseCase struct {
	pr *repository.ProductRepository
}

func NewProductUseCase(pr *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		pr: pr,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.pr.GetProducts()
}
