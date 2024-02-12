package routes

import (
	"myapp/models"
	"strconv"

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
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "could not fetch events try again later"})
		return
	}
	context.JSON(200, events)
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

// Get User   godoc
// @Summary    Get User
// @Description It fetches and return by id user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id [get]
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

// Update User   godoc
// @Summary    Update User
// @Description It updates user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id  [put]
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

// Delete User   godoc
// @Summary    Delete User
// @Description It deletes user stored in database
// @Param Authorization header string true "Authorization token"
// @Param id path int true "Event ID"
// @Produce application/json
// @Tags Events
// @Success 200
// @Router /events/:id [delete]
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
