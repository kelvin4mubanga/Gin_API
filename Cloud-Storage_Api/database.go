package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "cloud_storage_api/models"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("cloudstorageapi.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.Model{})
}
