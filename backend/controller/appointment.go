package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "blockchain.com/bc-67/config"
    "blockchain.com/bc-67/entity"
)

// CreateAppointment - POST /appointments
func CreateAppointment(c *gin.Context) {
    var appointment entity.Appointment

    // Bind JSON data to appointment struct
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    // Create a new appointment entry
    if err := db.Create(&appointment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Appointment created successfully",
        "data":    appointment,
    })
}

// GetAllAppointments - GET /appointments
func GetAllAppointments(c *gin.Context) {
    var appointments []entity.Appointment

    db := config.DB()

    // Fetch all appointments and preload the associated patient and department data
    results := db.Preload("Patient").Preload("Department").Find(&appointments)

    if results.Error != nil || results.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No appointments found"})
        return
    }

    // Return all appointments with preloaded patient and department data
    c.JSON(http.StatusOK, gin.H{
        "message": "Appointments fetched successfully",
        "data":    appointments,
    })
}

// GetAppointmentByID - GET /appointments/:id
func GetAppointmentByID(c *gin.Context) {
    id := c.Param("id")
    var appointment entity.Appointment

    db := config.DB()

    // Fetch appointment by ID and preload patient and department data
    if err := db.Preload("Patient").Preload("Department").First(&appointment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
        return
    }

    // Return appointment data with preloaded patient and department data
    c.JSON(http.StatusOK, gin.H{
        "message": "Appointment fetched successfully",
        "data":    appointment,
    })
}

// UpdateAppointment - PUT /appointments/:id
func UpdateAppointment(c *gin.Context) {
    id := c.Param("id")
    var appointment entity.Appointment

    db := config.DB()

    // Find the appointment by ID
    if err := db.First(&appointment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
        return
    }

    // Bind the updated data to the appointment struct
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the updated appointment data
    if err := db.Save(&appointment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Appointment updated successfully",
        "data":    appointment,
    })
}

// DeleteAppointment - DELETE /appointments/:id
func DeleteAppointment(c *gin.Context) {
    id := c.Param("id")
    db := config.DB()

    // Delete the appointment by ID
    if err := db.Delete(&entity.Appointment{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Appointment deleted successfully",
    })
}
