package repository

import (
	"fmt"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user SiteForPractic.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (phone_number, password, position, expirience, education_level, description) values ($1, $2, $3, $4, $5, $6) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Phone_number, user.Password, user.Position, user.Expirience, user.Education_level, user.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
