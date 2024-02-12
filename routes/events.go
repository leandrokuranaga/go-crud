package routes

import (
	"myapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch events try again later"})
		return
	}
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"message": "could not parse data"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(500, gin.H{"message": "could not create events try again later"})
		return
	}

	context.JSON(201, gin.H{"message": "Event created", "event": event})
}

func getEvent(context *gin.Context) {
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

	context.JSON(200, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(500, gin.H{"message": "could not parse id try again later"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

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
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch event try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "Event updated succesfully"})
}

func deleteEvent(context *gin.Context) {
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

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(401, gin.H{"message": "Not Authorized"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(500, gin.H{"message": "could not delete event try again later"})
		return
	}

	context.JSON(200, gin.H{"message": "event deleted successfully"})
}
