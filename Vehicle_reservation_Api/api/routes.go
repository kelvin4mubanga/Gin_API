package api

import (
	"github.com/gin-gonic/gin"
	"vehicle_reservation_api/database"
	"net/http"
)

func CreateReservation(c *gin.Context) {
	var reservation Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&reservation)
	c.JSON(http.StatusCreated, reservation)
}

func GetReservations(c *gin.Context) {
	var reservations []Reservation
	database.DB.Find(&reservations)
	c.JSON(http.StatusOK, reservations)
}

func GetReservation(c *gin.Context) {
	id := c.Param("id")
	var reservation Reservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

func UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var reservation Reservation
	if err := database.DB.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&reservation)
	c.JSON(http.StatusOK, reservation)
}

func DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Reservation{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/reservations", CreateReservation)
	router.GET("/reservations", GetReservations)
	router.GET("/reservations/:id", GetReservation)
	router.PUT("/reservations/:id", UpdateReservation)
	router.DELETE("/reservations/:id", DeleteReservation)
}
