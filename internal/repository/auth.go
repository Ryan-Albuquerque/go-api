package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ryan-Albuquerque/go-api/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) Register(user *model.Register) (model.RegisterResponse, error) {
	query, err := ar.db.Prepare("SELECT id FROM users WHERE username = $1 OR email = $2")

	if err != nil {
		return model.RegisterResponse{}, err
	}

	var id int

	err = query.QueryRow(user.Username, user.Email).Scan(&id)

	if err == nil {
		return model.RegisterResponse{}, fmt.Errorf("username or email already exists")
	}

	query, err = ar.db.Prepare("INSERT INTO users" +
		"(userName, email, password)" +
		" VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return model.RegisterResponse{}, err
	}

	err = query.QueryRow(user.Username, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return model.RegisterResponse{}, err
	}

	query.Close()

	return model.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (ar *AuthRepository) Login(user *model.Login) error {

	query, err := ar.db.Prepare("SELECT id FROM users WHERE username = $1 AND password = $2")

	if err != nil {
		return err
	}

	var id int

	err = query.QueryRow(user.Username, user.Password).Scan(&id)

	if err != nil {
		return fmt.Errorf("Invalid username or password")
	}

	query.Close()

	return nil
}
