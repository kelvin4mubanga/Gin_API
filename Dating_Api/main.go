package main

import (
	"dating_api/api"
	"dating_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Profile{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
