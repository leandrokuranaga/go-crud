package register

import (
	"myapp/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	Repo repository.EventRepository
}

func NewEventHandler(repo repository.EventRepository) *EventHandler {
	return &EventHandler{Repo: repo}
}

// Register For Event godoc
// @Summary Register for an event
// @Description Registers a user for an event by their IDs
// @Tags Log
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Authorization token"
// @Param id path int64 true "Event ID"
// @Success 201 {object} object "Successfully registered"
// @Failure 500 {object} object "Error message"
// @Router /events/register/{id} [post]
func (h *EventHandler) RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not parse id try again later"})
		return
	}

	event, err := h.Repo.GetEventById(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch event try again later"})
		return
	}

	err = h.Repo.Register(userId, event.Id)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not register event try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "Registered"})
}
