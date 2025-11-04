package main

import (
	"rearatrox/event-booking-api/pkg/config"
	"rearatrox/event-booking-api/services/user-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	db.InitDB()

	router := gin.Default()

	RegisterRoutes(router)

	router.Run(":8080") // localhost:8080 --> USERSERVICE_PORT mappt dann den Container

}
