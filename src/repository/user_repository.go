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

func (ur *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	user := domain.User{}
	cmd := "SELECT id, email, password FROM users WHERE email = ?"
	if err := ur.db.QueryRow(cmd, email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			myError := domain.NewMyError(domain.NotFound, "user data")
			return user, myError
		} else {
			return user, err
		}
	}
	return user, nil
}

// TODO: 連続で複数回インサートしたい時はdbではなくプリペアドステートメントを利用したい
func (ur *UserRepository) CreateUser(user *domain.User) error {
	cmd := "INSERT INTO users (email, password, created_at, updated_at) VALUES (?, ?, now(), now())"
	result, err := ur.db.Exec(cmd, user.Email, user.Password)
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

func (ur *UserRepository) CheckDuplicateEmail(email string) (bool, error) {
	cmd := "SELECT count(email) FROM users where email = ?"
	var count int
	if err := ur.db.QueryRow(cmd, email).Scan(&count); err != nil {
		return false, err
	}

	return count != 0, nil
}