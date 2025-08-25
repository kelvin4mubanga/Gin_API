package api

import (
	"github.com/gin-gonic/gin"
	"online_auction_api/database"
	"net/http"
)

func CreateAuction(c *gin.Context) {
	var auction Auction
	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&auction)
	c.JSON(http.StatusCreated, auction)
}

func GetAuctions(c *gin.Context) {
	var auctions []Auction
	database.DB.Find(&auctions)
	c.JSON(http.StatusOK, auctions)
}

func GetAuction(c *gin.Context) {
	id := c.Param("id")
	var auction Auction
	if err := database.DB.First(&auction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}
	c.JSON(http.StatusOK, auction)
}

func UpdateAuction(c *gin.Context) {
	id := c.Param("id")
	var auction Auction
	if err := database.DB.First(&auction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}
	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&auction)
	c.JSON(http.StatusOK, auction)
}

func DeleteAuction(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Auction{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Auction deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/auctions", CreateAuction)
	router.GET("/auctions", GetAuctions)
	router.GET("/auctions/:id", GetAuction)
	router.PUT("/auctions/:id", UpdateAuction)
	router.DELETE("/auctions/:id", DeleteAuction)
}
