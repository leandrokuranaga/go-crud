package register

import (
	"myapp/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Cancel Event Registration godoc
// @Summary Cancel event registration
// @Description Cancels a user's registration for an event by their IDs
// @Tags Log
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Authorization token"
// @Param id path int64 true "Event ID"
// @Success 200 {object} object "Successfully cancelled"
// @Failure 500 {object} object "Error message"
// @Router /events/cancel/{id} [post]
func (h *EventHandler) CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.Id = eventId

	err = h.Repo.CancelRegistration(userId, eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not cancel registration try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "Cancelled"})
}
