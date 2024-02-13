package events

import (
	"github.com/gin-gonic/gin"
)

// Get Events godoc
// @Summary    Get Events
// @Description It fetches and return all events stored in database
// @Param Authorization header string true "Authorization token"
// @Produce application/json
// @Tags Events
// @Success 200 {array} models.Event "An array of events"
// @Router /events [get]
func (h *EventHandler) GetEvents(context *gin.Context) {
	events, err := h.Repo.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch events try again later"})
		return
	}
	context.JSON(200, events)
}
