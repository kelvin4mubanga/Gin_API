package main

import (
	"cloud_storage_api/api"
	"cloud_storage_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.File{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
