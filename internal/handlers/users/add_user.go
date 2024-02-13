package users

import (
	"myapp/internal/models"
	"myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	Repo repository.UserRepository
}

func NewEventHandler(repo repository.UserRepository) *EventHandler {
	return &EventHandler{Repo: repo}
}

// Signup godoc
// @Summary Register a new user
// @Description It creates a new user in the database
// @Accept json
// @Produce application/json
// @Tags Users
// @Param user body models.User true "User data for registration"
// @Success 201
// @Failure 400 {object} object "Could not parse data"
// @Failure 500 {object} object "Could not create user, try again later"
// @Router /signup [post]
func (h *EventHandler) Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	err = h.Repo.Save(&user)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not create user try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "User created", "user": user})
}
