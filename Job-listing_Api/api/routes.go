package api

import (
	"github.com/gin-gonic/gin"
	"job_listing_api/database"
	"net/http"
)

func CreateJob(c *gin.Context) {
	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&job)
	c.JSON(http.StatusCreated, job)
}

func GetJobs(c *gin.Context) {
	var jobs []Job
	database.DB.Find(&jobs)
	c.JSON(http.StatusOK, jobs)
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	var job Job
	if err := database.DB.First(&job, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job Job
	if err := database.DB.First(&job, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&job)
	c.JSON(http.StatusOK, job)
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Job{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/jobs", CreateJob)
	router.GET("/jobs", GetJobs)
	router.GET("/jobs/:id", GetJob)
	router.PUT("/jobs/:id", UpdateJob)
	router.DELETE("/jobs/:id", DeleteJob)
}
