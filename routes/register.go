package routes

import (
	"myapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
