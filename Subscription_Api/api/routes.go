package api

import (
	"github.com/gin-gonic/gin"
	"subscription_api/database"
	"net/http"
)

func CreateSubscription(c *gin.Context) {
	var subscription Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&subscription)
	c.JSON(http.StatusCreated, subscription)
}

func GetSubscriptions(c *gin.Context) {
	var subscriptions []Subscription
	database.DB.Find(&subscriptions)
	c.JSON(http.StatusOK, subscriptions)
}

func GetSubscription(c *gin.Context) {
	id := c.Param("id")
	var subscription Subscription
	if err := database.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	c.JSON(http.StatusOK, subscription)
}

func UpdateSubscription(c *gin.Context) {
	id := c.Param("id")
	var subscription Subscription
	if err := database.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&subscription)
	c.JSON(http.StatusOK, subscription)
}

func DeleteSubscription(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Subscription{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/subscriptions", CreateSubscription)
	router.GET("/subscriptions", GetSubscriptions)
	router.GET("/subscriptions/:id", GetSubscription)
	router.PUT("/subscriptions/:id", UpdateSubscription)
	router.DELETE("/subscriptions/:id", DeleteSubscription)
}
