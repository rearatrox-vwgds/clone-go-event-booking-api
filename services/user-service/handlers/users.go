package handlers

import (
	"net/http"
	"rearatrox/event-booking-api/services/user-service/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	users, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events.", "error": err.Error()})
		return
	}
	//Response in JSON
	context.JSON(http.StatusOK, users)
}
