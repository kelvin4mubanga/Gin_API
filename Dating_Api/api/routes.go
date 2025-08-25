package api

import (
	"github.com/gin-gonic/gin"
	"dating_api/database"
	"net/http"
)

func CreateProfile(c *gin.Context) {
	var profile Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&profile)
	c.JSON(http.StatusCreated, profile)
}

func GetProfiles(c *gin.Context) {
	var profiles []Profile
	database.DB.Find(&profiles)
	c.JSON(http.StatusOK, profiles)
}

func GetProfile(c *gin.Context) {
	id := c.Param("id")
	var profile Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	var profile Profile
	if err := database.DB.First(&profile, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&profile)
	c.JSON(http.StatusOK, profile)
}

func DeleteProfile(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Profile{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/profiles", CreateProfile)
	router.GET("/profiles", GetProfiles)
	router.GET("/profiles/:id", GetProfile)
	router.PUT("/profiles/:id", UpdateProfile)
	router.DELETE("/profiles/:id", DeleteProfile)
}
