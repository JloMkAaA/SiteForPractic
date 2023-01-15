package repository

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user SiteForPractic.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
