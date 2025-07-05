package models

import (
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
