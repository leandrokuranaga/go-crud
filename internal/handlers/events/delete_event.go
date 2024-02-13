package events

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete User   godoc
// @Summary    Delete User
// @Description It deletes user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id [delete]
func (h *EventHandler) DeleteEvent(context *gin.Context) {
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

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(401, gin.H{"message": "Not Authorized"})
		return
	}

	err = h.Repo.Delete(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not delete event try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "event deleted successfully"})
}
