package main

import (
	"job_listing_api/api"
	"job_listing_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Job{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
