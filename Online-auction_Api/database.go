package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "online_auction_api/models"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("onlineauctionapi.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.Model{})
}
