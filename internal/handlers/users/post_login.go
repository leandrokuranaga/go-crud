package users

import (
	"myapp/internal/models"
	"myapp/internal/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary User login
// @Description Authenticate a user and return a token
// @Accept json
// @Produce application/json
// @Param user body models.User true "User data for registration"
// @Tags Users
// @Success 200 {object} object "Login successful, token returned"
// @Failure 400 {object} object "Could not parse data"
// @Failure 401 {object} object "Invalid credentials"
// @Router /login [post]
func (h *EventHandler) Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	err = h.Repo.ValidateCredentials(&user)

	if err != nil {
		context.JSON(401, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(401, gin.H{"message": err.Error()})
		return
	}

	context.JSON(200, gin.H{"token": token})

}
