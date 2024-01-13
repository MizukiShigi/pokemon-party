package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type IUserRepository interface {
	GetUser(user *domain.User, id int) error
	CreateUser(user *domain.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUser(user *domain.User, id int) error {
	cmd := "SELECT id, name, email FROM users WHERE id = ?"
	if err := ur.db.QueryRow(cmd, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(user *domain.User) error {
	cmd := "INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, now(), now())"
	result, err := ur.db.Exec(cmd, user.Name, user.Email)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New(fmt.Sprintf("expected to affect 1 row, affected %d", rows))
	}
	return nil
}
