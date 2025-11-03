package handlers

import (
	"net/http"
	"rearatrox/event-booking-api/services/event-service/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Event-Handlers, die beim Aufruf von Routen /events aufgerufen werden (Verarbeitung der Requests)

func GetEvents(context *gin.Context) {

	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events.", "error": err.Error()})
		return
	}
	//Response in JSON
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	var event *models.Event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id.", "error": err.Error()})
		return
	}

	event, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event.", "error": err.Error()})
		return
	}

	//Response in JSON
	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.CreatorID = userId

	err = event.SaveEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func UpdateEvent(context *gin.Context) {
	var event *models.Event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id to update event.", "error": err.Error()})
		return
	}

	event, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event.", "error": err.Error()})
		return
	}

	userId, _ := context.Get("userId")

	if event.CreatorID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update that event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event.", "error": err.Error()})
		return
	}
	//Response in JSON
	context.JSON(http.StatusOK, gin.H{"message": "updated event successfully", "updatedEvent": updatedEvent})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id to update event.", "error": err.Error()})
		return
	}

	deleteEvent, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event.", "error": err.Error()})
		return
	}

	userId, _ := context.Get("userId")

	if deleteEvent.CreatorID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete that event"})
		return
	}

	err = deleteEvent.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event.", "error": err.Error()})
		return
	}
	//Response in JSON
	context.JSON(http.StatusOK, gin.H{"message": "deleted event successfully", "deletedEvent": deleteEvent})
}
