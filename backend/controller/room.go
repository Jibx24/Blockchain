package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "blockchain.com/bc-67/config"
    "blockchain.com/bc-67/entity"
)

// CreateRoom - POST /rooms
func CreateRoom(c *gin.Context) {
    var room entity.Room

    // Bind JSON data to the room struct
    if err := c.ShouldBindJSON(&room); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    // Create a new room entry
    if err := db.Create(&room).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Room created successfully",
        "data":    room,
    })
}

// GetAllRooms - GET /rooms
func GetAllRooms(c *gin.Context) {
    var rooms []entity.Room

    db := config.DB()

    // Fetch all rooms
    results := db.Find(&rooms)

    if results.Error != nil || results.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No rooms found"})
        return
    }

    // Return all rooms
    c.JSON(http.StatusOK, gin.H{
        "message": "Rooms fetched successfully",
        "data":    rooms,
    })
}

// GetRoomByID - GET /rooms/:id
func GetRoomByID(c *gin.Context) {
    id := c.Param("id")
    var room entity.Room

    db := config.DB()

    // Fetch room by ID
    if err := db.First(&room, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
        return
    }

    // Return room data
    c.JSON(http.StatusOK, gin.H{
        "message": "Room fetched successfully",
        "data":    room,
    })
}

// UpdateRoom - PUT /rooms/:id
func UpdateRoom(c *gin.Context) {
    id := c.Param("id")
    var room entity.Room

    db := config.DB()

    // Find the room by ID
    if err := db.First(&room, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
        return
    }

    // Bind the updated data to the room struct
    if err := c.ShouldBindJSON(&room); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the updated room data
    if err := db.Save(&room).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Room updated successfully",
        "data":    room,
    })
}

// DeleteRoom - DELETE /rooms/:id
func DeleteRoom(c *gin.Context) {
    id := c.Param("id")
    db := config.DB()

    // Delete the room by ID
    if err := db.Delete(&entity.Room{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Room deleted successfully",
    })
}
