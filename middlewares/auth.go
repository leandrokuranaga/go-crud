package middlewares

import (
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"message": "Empty token"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "token not valid"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
