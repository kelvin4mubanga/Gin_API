package api

import (
	"github.com/gin-gonic/gin"
	"travel_booking_api/database"
	"net/http"
)

func CreateBooking(c *gin.Context) {
	var booking Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&booking)
	c.JSON(http.StatusCreated, booking)
}

func GetBookings(c *gin.Context) {
	var bookings []Booking
	database.DB.Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}

func GetBooking(c *gin.Context) {
	id := c.Param("id")
	var booking Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var booking Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&booking)
	c.JSON(http.StatusOK, booking)
}

func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Booking{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/bookings", CreateBooking)
	router.GET("/bookings", GetBookings)
	router.GET("/bookings/:id", GetBooking)
	router.PUT("/bookings/:id", UpdateBooking)
	router.DELETE("/bookings/:id", DeleteBooking)
}
