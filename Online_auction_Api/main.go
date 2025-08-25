package main

import (
	"online_auction_api/api"
	"online_auction_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&api.Auction{})

	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}
