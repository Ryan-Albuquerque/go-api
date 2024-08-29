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

func (pu *ProductUseCase) CreateProduct(product model.Product) error {
	return pu.pr.CreateProduct(product)
}

func (pu *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	return pu.pr.GetProductByID(id)
}

func (pu *ProductUseCase) UpdateProduct(id int, product model.Product) (*model.Product, error) {
	return pu.pr.UpdateProduct(id, product)
}

func (pu *ProductUseCase) DeleteProduct(id int) (*model.Product, error) {
	return pu.pr.DeleteProduct(id)
}
