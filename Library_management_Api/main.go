package main

import (
	"library_management_api/api"
	"library_management_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Book{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
