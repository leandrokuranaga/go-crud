package routes

import (
	"myapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not parse id try again later"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch event try again later"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not register event try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "Registered"})
}

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
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.Id = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not cancel registration try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "Cancelled"})
}
