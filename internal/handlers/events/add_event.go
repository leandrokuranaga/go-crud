package events

import (
	models "myapp/internal/models"
	"myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	Repo repository.EventRepository
}

func NewEventHandler(repo repository.EventRepository) *EventHandler {
	return &EventHandler{Repo: repo}
}

// Create Users   godoc
// @Summary    Create Users
// @Description It creates and return all users stored in database
// @Param Authorization header string true "Authorization token"
// @Accept json
// @Produce application/json
// @Param events body []models.Event true "Array of Event objects to be created"
// @Tags Events
// @Success 201
// @Router /events [Post]
func (h *EventHandler) CreateEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	if err := h.Repo.Save(&event); err != nil {
		context.JSON(500, gin.H{"message": "could not create events try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "Event created", "event": event})
}
