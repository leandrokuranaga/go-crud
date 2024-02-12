package routes

import (
	"myapp/models"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

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
func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(500, gin.H{"message": "could not create user try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "User created", "user": user})
}

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
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	err = user.ValidateCredentials()

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
