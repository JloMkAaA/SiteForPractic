package repository

import "github.com/jmoiron/sqlx"

type ManipulateUser interface {
}

type Repository struct {
	ManipulateUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
