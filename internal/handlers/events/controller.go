package events

import (
	"myapp/internal/middlewares"
	"myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, eventRepo repository.EventRepository) {
	eventHandler := NewEventHandler(eventRepo)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", eventHandler.CreateEvent)
	authenticated.GET("/events/:id", eventHandler.GetEvent)
	authenticated.GET("/events", eventHandler.GetEvents)
	authenticated.PUT("/events/:id", eventHandler.UpdateEvent)
	authenticated.DELETE("/events/:id", eventHandler.DeleteEvent)
}
