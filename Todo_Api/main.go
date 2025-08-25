package main

import (
	"todo_api/api"
	"todo_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Todo{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
