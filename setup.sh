#!/bin/bash

API_DIRS=(
  Blog_Api
  Cloud-Storage_Api
  Dating_Api
  Digital-wallet_Api
  E-commerce_Api
  Job-listing_Api
  Library_management_Api
  Micro-lending_Api
  Online-auction_Api
  Student_Management_Api
  Subscription_Api
  Todo-Api
  Travel_Api
  Travel_Booking_Api
  Vehicle-reservation_Api
)

for dir in "${API_DIRS[@]}"; do
  echo "Setting up $dir..."
  cd "$dir"
  go mod init github.com/kelvin4mubanga/$dir
  go get github.com/gin-gonic/gin
  go get gorm.io/gorm
  go get gorm.io/driver/sqlite
  cd ..
done

echo "ll APIs are set up with Gin, GORM, and SQLite3!"
