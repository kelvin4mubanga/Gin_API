package api

import (
	"github.com/gin-gonic/gin"
	"travel_api/database"
	"net/http"
)

func CreateDestination(c *gin.Context) {
	var destination Destination
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&destination)
	c.JSON(http.StatusCreated, destination)
}

func GetDestinations(c *gin.Context) {
	var destinations []Destination
	database.DB.Find(&destinations)
	c.JSON(http.StatusOK, destinations)
}

func GetDestination(c *gin.Context) {
	id := c.Param("id")
	var destination Destination
	if err := database.DB.First(&destination, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	c.JSON(http.StatusOK, destination)
}

func UpdateDestination(c *gin.Context) {
	id := c.Param("id")
	var destination Destination
	if err := database.DB.First(&destination, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&destination)
	c.JSON(http.StatusOK, destination)
}

func DeleteDestination(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Destination{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Destination deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/destinations", CreateDestination)
	router.GET("/destinations", GetDestinations)
	router.GET("/destinations/:id", GetDestination)
	router.PUT("/destinations/:id", UpdateDestination)
	router.DELETE("/destinations/:id", DeleteDestination)
}
