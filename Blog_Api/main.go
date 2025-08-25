package main

import (
	"blog_api/api"
	"blog_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Post{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
