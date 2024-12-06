package repository

import (
	"database/sql"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}


func (r UserRepository) GetUser() ([]model.User, error) {
	query := "SELECT * FROM user"
	rows, err := r.db.Query(query)
	if err !=nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID, 
			&user.Username,
			&user.Password,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}