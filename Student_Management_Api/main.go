package main

import (
	"student_management_api/api"
	"student_management_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Student{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
