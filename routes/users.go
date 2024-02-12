package routes

import (
	"myapp/models"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

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
