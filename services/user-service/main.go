package main

import (
	"fmt"
	"os"
	"rearatrox/event-booking-api/pkg/config"
	"rearatrox/event-booking-api/services/user-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	fmt.Println(os.Getenv("JWT_SECRET"))

	db.InitDB()

	router := gin.Default()

	RegisterRoutes(router)

	router.Run(fmt.Sprintf(":%v", os.Getenv("USERSERVICE_PORT"))) // localhost:USERSERVICE_PORT

}
