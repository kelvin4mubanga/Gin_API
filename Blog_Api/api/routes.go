package api

import (
	"github.com/gin-gonic/gin"
	"blog_api/database"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	var posts []Post
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&Post{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/posts", CreatePost)
	router.GET("/posts", GetPosts)
	router.GET("/posts/:id", GetPost)
	router.PUT("/posts/:id", UpdatePost)
	router.DELETE("/posts/:id", DeletePost)
}
