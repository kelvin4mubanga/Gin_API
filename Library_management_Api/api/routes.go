package api

import (
	"github.com/gin-gonic/gin"
	"library_management_api/database"
	"net/http"
)

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
	var books []Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Book{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/books", CreateBook)
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)
}
