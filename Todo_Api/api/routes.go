package api

import (
	"github.com/gin-gonic/gin"
	"todo_api/database"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	var todos []Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Todo{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/todos", CreateTodo)
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodo)
	router.PUT("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)
}
