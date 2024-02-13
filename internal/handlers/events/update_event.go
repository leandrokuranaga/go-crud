package events

import (
	"myapp/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Update User   godoc
// @Summary    Update User
// @Description It updates user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id  [put]
func (h *EventHandler) UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not parse id try again later"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := h.Repo.GetEventById(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch event try again later"})
		return
	}

	if event.UserId != userId {
		context.JSON(401, gin.H{"message": "Not Authorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}

	updatedEvent.Id = eventId
	err = h.Repo.Update(updatedEvent)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch event try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "Event updated succesfully"})
}
