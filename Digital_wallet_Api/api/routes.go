package api

import (
	"github.com/gin-gonic/gin"
	"digital_wallet_api/database"
	"net/http"
)

func CreateWallet(c *gin.Context) {
	var wallet Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&wallet)
	c.JSON(http.StatusCreated, wallet)
}

func GetWallets(c *gin.Context) {
	var wallets []Wallet
	database.DB.Find(&wallets)
	c.JSON(http.StatusOK, wallets)
}

func GetWallet(c *gin.Context) {
	id := c.Param("id")
	var wallet Wallet
	if err := database.DB.First(&wallet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(http.StatusOK, wallet)
}

func UpdateWallet(c *gin.Context) {
	id := c.Param("id")
	var wallet Wallet
	if err := database.DB.First(&wallet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&wallet)
	c.JSON(http.StatusOK, wallet)
}

func DeleteWallet(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Wallet{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Wallet deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/wallets", CreateWallet)
	router.GET("/wallets", GetWallets)
	router.GET("/wallets/:id", GetWallet)
	router.PUT("/wallets/:id", UpdateWallet)
	router.DELETE("/wallets/:id", DeleteWallet)
}
