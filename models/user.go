package models

import (
	"errors"

	"example.com/go_rest_api/db"
	"example.com/go_rest_api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "insert into users (email, password) values ($1, $2) returning id"

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	var returningId int64
	err = db.DB.QueryRow(
		query,
		u.Email,
		hashedPassword,
	).Scan(&returningId)

	u.Id = returningId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `select id, password from users where email=$1`

	r := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := r.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
