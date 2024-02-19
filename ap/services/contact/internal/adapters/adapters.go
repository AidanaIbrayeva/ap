package adapters

import (
	"database/sql"
	"services/contact/internal/domain"
)

type UserRepositoryPostgres struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{DB: db}
}

func (r *UserRepositoryPostgres) CreateUser(user *domain.User) (*domain.User, error) {

	_, err := r.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryPostgres) GetUserByID(userID int) (*domain.User, error) {

	var user domain.User
	err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
