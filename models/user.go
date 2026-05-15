package models

import (
	"github.com/shaneoh10/events-backend/db"
	"github.com/shaneoh10/events-backend/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password) 
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = int64(id)

	return err
}
