package main

import (
	"travel_api/api"
	"travel_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Destination{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
