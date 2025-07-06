package models

import (
	"errors"

	"github.com/popey17/go-rest-api/db"
	"github.com/popey17/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := ` INSERT INTO users(email, password)VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	u.ID = lastID
	return err

}

func (u *User) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	resultRow := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := resultRow.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return err
	}

	isValid := utils.ComparePassword(hashedPassword, u.Password)

	if !isValid {
		return errors.New("invalid password")
	}

	return nil

}
