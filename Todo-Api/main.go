// Todo-Api/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "net/http"
)

type Item struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
}

func main() {
    db, err := gorm.Open(sqlite.Open("Todo-Api.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Item{})

    r := gin.Default()

    r.GET("/items", func(c *gin.Context) {
        var items []Item
        db.Find(&items)
        c.JSON(http.StatusOK, items)
    })

    r.POST("/items", func(c *gin.Context) {
        var item Item
        if err := c.ShouldBindJSON(&item); err == nil {
            db.Create(&item)
            c.JSON(http.StatusCreated, item)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    r.PUT("/items/:id", func(c *gin.Context) {
        var item Item
        if err := db.First(&item, c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
            return
        }
        if err := c.ShouldBindJSON(&item); err == nil {
            db.Save(&item)
            c.JSON(http.StatusOK, item)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    r.DELETE("/items/:id", func(c *gin.Context) {
        db.Delete(&Item{}, c.Param("id"))
        c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
    })

    r.Run(":8080")
}
