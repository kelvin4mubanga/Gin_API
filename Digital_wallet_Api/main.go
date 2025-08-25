package main

import (
	"digital_wallet_api/api"
	"digital_wallet_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Wallet{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
