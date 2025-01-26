package user

import (
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &userRepository{db}
}

func (r UserRepository) GetUser() ([]User, error) {
	query := "SELECT * FROM user"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
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
