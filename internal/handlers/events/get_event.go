package events

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get User   godoc
// @Summary    Get User
// @Description It fetches and return by id user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id [get]
func (h *EventHandler) GetEvent(context *gin.Context) {
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

	context.JSON(200, event)
}
