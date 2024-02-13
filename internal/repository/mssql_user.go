package repository

import (
	"database/sql"
	"errors"
	"fmt"
	models "myapp/internal/models"
	"myapp/internal/utils"
)

type UserRepository interface {
	Save(user *models.User) error
	ValidateCredentials(user *models.User) error
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (repo *userRepositoryImpl) Save(user *models.User) error {
	query := `
    INSERT INTO users(email, password) 
    OUTPUT INSERTED.id
    VALUES (@p1, @p2)
    `
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	var id int64
	err = stmt.QueryRow(user.Email, pass).Scan(&id)
	if err != nil {
		return err
	}

	user.Id = id
	return nil
}

func (repo *userRepositoryImpl) ValidateCredentials(user *models.User) error {
	query := `SELECT id, password FROM users WHERE email = @p1`
	row := repo.db.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("No user found with the given email")
		}
		return fmt.Errorf("database scan error: %v", err)
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid password")
	}

	return nil
}
