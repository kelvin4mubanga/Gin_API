package main

import (
	"vehicle_reservation_api/api"
	"vehicle_reservation_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Reservation{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
