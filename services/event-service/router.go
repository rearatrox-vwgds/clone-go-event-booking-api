package main

import (
	middleware "rearatrox/event-booking-api/pkg/middleware/auth"
	"rearatrox/event-booking-api/services/event-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/events", handlers.GetEvents)
	router.GET("/events/:id", handlers.GetEvent)

	authenticated := router.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", handlers.DeleteEvent)

	authenticated.POST("/events/:id/register", handlers.AddRegistrationForEvent)
	authenticated.DELETE("/events/:id/delete", handlers.DeleteRegistrationForEvent)
}
