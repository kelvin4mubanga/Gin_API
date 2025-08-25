package api

import (
	"github.com/gin-gonic/gin"
	"cloud_storage_api/database"
	"net/http"
)

func CreateFile(c *gin.Context) {
	var file File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&file)
	c.JSON(http.StatusCreated, file)
}

func GetFiles(c *gin.Context) {
	var files []File
	database.DB.Find(&files)
	c.JSON(http.StatusOK, files)
}

func GetFile(c *gin.Context) {
	id := c.Param("id")
	var file File
	if err := database.DB.First(&file, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.JSON(http.StatusOK, file)
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	var file File
	if err := database.DB.First(&file, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&file)
	c.JSON(http.StatusOK, file)
}

func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&File{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/files", CreateFile)
	router.GET("/files", GetFiles)
	router.GET("/files/:id", GetFile)
	router.PUT("/files/:id", UpdateFile)
	router.DELETE("/files/:id", DeleteFile)
}
