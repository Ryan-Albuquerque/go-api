package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ryan-Albuquerque/go-api/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	rows, err := pr.db.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) error {
	var id int
	query, err := pr.db.Prepare("INSERT INTO product" +
		"(name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	err := pr.db.QueryRow("SELECT * FROM product WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(id int, product model.Product) (*model.Product, error) {
	getProduct, err := pr.GetProductByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if getProduct == nil {
		fmt.Printf("product with id %d does not exist\n", id)
		return nil, nil
	}

	query, err := pr.db.Prepare("UPDATE product SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(product.Name, product.Price, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) (*model.Product, error) {
	product, err := pr.GetProductByID(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if product == nil {
		fmt.Printf("product with id %d does not exist\n", id)
		return nil, nil
	}

	query, err := pr.db.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()
	return product, nil
}
