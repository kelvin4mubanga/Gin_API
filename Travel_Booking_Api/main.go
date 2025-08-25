package main

import (
	"travel_booking_api/api"
	"travel_booking_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Booking{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
