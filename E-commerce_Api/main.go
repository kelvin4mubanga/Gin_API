package main

import (
	"ecommerce_api/api"
	"ecommerce_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Product{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
