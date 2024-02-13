package users

import (
	"myapp/internal/repository"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(server *gin.Engine, userRepo repository.UserRepository) {
	eventHandler := NewEventHandler(userRepo)

	server.POST("/signup", eventHandler.Signup)
	server.POST("/login", eventHandler.Login)
}
