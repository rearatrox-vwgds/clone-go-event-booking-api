package handlers

import (
	"net/http"
	"rearatrox/event-booking-api/services/user-service/models"

	"github.com/gin-gonic/gin"
)

// User-Handlers, die beim Aufruf von Routen /users aufgerufen werden (Verarbeitung der Requests)

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}

	err = user.SaveUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
