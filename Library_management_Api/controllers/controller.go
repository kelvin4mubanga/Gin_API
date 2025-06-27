package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func DummyHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello from controller"})
}
