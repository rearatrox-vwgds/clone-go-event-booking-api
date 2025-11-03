package handlers

import (
	"net/http"
	"rearatrox/event-booking-api/services/event-service/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRegistrationForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id.", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully registered user for event."})

}

func DeleteRegistrationForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id.", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event."})
		return
	}

	err = event.DeleteRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete user registration for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully deleted user registration for event."})

}
