package register

import (
	"myapp/internal/middlewares"
	"myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, eventRepo repository.EventRepository) {

	eventHandler := NewEventHandler(eventRepo)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events/:id/register", eventHandler.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", eventHandler.CancelRegistration)
}
