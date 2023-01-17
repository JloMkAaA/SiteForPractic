package repository

import (
	"fmt"
	"strings"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserById(userIdAuth, userId int) (SiteForPractic.User, error) {
	var user SiteForPractic.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND id=$2", userTable)

	err := r.db.Get(&user, query, userIdAuth, userId)

	return user, err
}

func (r *UserPostgres) Delete(userIdAuth, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND id=$2", userTable)
	_, err := r.db.Exec(query, userIdAuth, userId)

	return err
}

func (r *UserPostgres) Update(userIdAuth, userId int, input SiteForPractic.UpdateUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Phone_number != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number = $%d", argId))
		args = append(args, *input.Phone_number)
		argId++
	}
	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password = $%d", argId))
		args = append(args, *input.Password)
		argId++
	}
	if input.Position != nil {
		setValues = append(setValues, fmt.Sprintf("position = $%d", argId))
		args = append(args, *input.Position)
		argId++
	}
	if input.Expirience != nil {
		setValues = append(setValues, fmt.Sprintf("expirience = $%d", argId))
		args = append(args, *input.Expirience)
		argId++
	}
	if input.Education_level != nil {
		setValues = append(setValues, fmt.Sprintf("education_level = $%d", argId))
		args = append(args, *input.Education_level)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	// UPDATE users SET phone_number = '123123123', password = '321321', description = 'ВаВавывау' WHERE ...

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d AND id=%d", userTable, setQuery, userIdAuth, userId)

	logrus.Debugf("Update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
