package main

import (
	"micro_lending_api/api"
	"micro_lending_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Loan{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
