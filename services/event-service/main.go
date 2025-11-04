package main

import (
	"rearatrox/event-booking-api/pkg/config"
	"rearatrox/event-booking-api/services/event-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	db.InitDB()

	router := gin.Default()

	RegisterRoutes(router)

	router.Run(":8080") // localhost:8080 -> Container-Port wird durch EVENTSERVICE_PORT gesetzt

}
