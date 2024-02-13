package models

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}