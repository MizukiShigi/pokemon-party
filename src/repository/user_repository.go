package repository

import (
	"database/sql"
	"log"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUser(user *domain.User) error {
	cmd := "SELECT id, name, email FROM users WHERE id = ?"
	if err := ur.db.QueryRow(cmd, user.ID).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			myError := domain.NewMyError(domain.NotFound, "user data")
			return myError
		} else {
			return err
		}
	}
	return nil
}

// TODO: 連続で複数回インサートしたい時はdbではなくプリペアドステートメントを利用できるようにする
func (ur *UserRepository) CreateUser(user *domain.User) error {
	cmd := "INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, now(), now())"
	result, err := ur.db.Exec(cmd, user.Name, user.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.ID = int(id)
	return nil
}
