package models

import (
	"database/sql"
	"errors"
	"fmt"
	"myapp/db"
	"myapp/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) 
	OUTPUT INSERTED.id
	VALUES (@p1, @p2)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	pass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	var id int64
	err = stmt.QueryRow(u.Email, pass).Scan(&id)

	if err != nil {
		return err
	}

	u.Id = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = @p1`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("No user found with the given email")
		}
		return fmt.Errorf("database scan error: %v", err)
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid password")
	}

	return nil
}
