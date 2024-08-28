package usecase

import "github.com/Ryan-Albuquerque/go-api/internal/model"

type ProductUseCase struct {
}

func NewProductUseCase() *ProductUseCase {
	return &ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
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
	return products, nil
}
