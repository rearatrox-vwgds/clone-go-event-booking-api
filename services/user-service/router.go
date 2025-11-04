package main

import (
	"os"
	"rearatrox/event-booking-api/services/user-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group(os.Getenv("API_PREFIX"))
	{
		api.GET("/users", handlers.GetUsers)

		api.POST("/users/signup", handlers.Signup)
		api.POST("/users/login", handlers.Login)
	}

}
