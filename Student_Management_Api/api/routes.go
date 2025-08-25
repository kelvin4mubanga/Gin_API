package api

import (
	"github.com/gin-gonic/gin"
	"student_management_api/database"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func GetStudents(c *gin.Context) {
	var students []Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Student{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/students", CreateStudent)
	router.GET("/students", GetStudents)
	router.GET("/students/:id", GetStudent)
	router.PUT("/students/:id", UpdateStudent)
	router.DELETE("/students/:id", DeleteStudent)
}
