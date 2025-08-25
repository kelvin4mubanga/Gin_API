package api

import (
	"github.com/gin-gonic/gin"
	"micro_lending_api/database"
	"net/http"
)

func CreateLoan(c *gin.Context) {
	var loan Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&loan)
	c.JSON(http.StatusCreated, loan)
}

func GetLoans(c *gin.Context) {
	var loans []Loan
	database.DB.Find(&loans)
	c.JSON(http.StatusOK, loans)
}

func GetLoan(c *gin.Context) {
	id := c.Param("id")
	var loan Loan
	if err := database.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	c.JSON(http.StatusOK, loan)
}

func UpdateLoan(c *gin.Context) {
	id := c.Param("id")
	var loan Loan
	if err := database.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&loan)
	c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Loan{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/loans", CreateLoan)
	router.GET("/loans", GetLoans)
	router.GET("/loans/:id", GetLoan)
	router.PUT("/loans/:id", UpdateLoan)
	router.DELETE("/loans/:id", DeleteLoan)
}
