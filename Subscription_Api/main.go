package main

import (
	"subscription_api/api"
	"subscription_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Subscription{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
