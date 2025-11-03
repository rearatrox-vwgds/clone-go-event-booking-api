package main

import (
	"rearatrox/event-booking-api/services/user-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
}
