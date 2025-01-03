package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "blockchain.com/bc-67/config"
    "blockchain.com/bc-67/entity"
)

// CreatePatient - POST /patients
func CreatePatient(c *gin.Context) {
    var patient entity.Patient

    // Bind JSON data to the patient struct
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    // Create a new patient entry
    if err := db.Create(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Patient created successfully",
        "data":    patient,
    })
}

// GetAllPatients - GET /patients
func GetAllPatients(c *gin.Context) {
    var patients []entity.Patient

    db := config.DB()

    // Fetch all patients
    results := db.Find(&patients)

    if results.Error != nil || results.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No patients found"})
        return
    }

    // Return all patients
    c.JSON(http.StatusOK, gin.H{
        "message": "Patients fetched successfully",
        "data":    patients,
    })
}

// GetPatientByID - GET /patients/:id
func GetPatientByID(c *gin.Context) {
    id := c.Param("id")
    var patient entity.Patient

    db := config.DB()

    // Fetch patient by ID
    if err := db.First(&patient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }

    // Return patient data
    c.JSON(http.StatusOK, gin.H{
        "message": "Patient fetched successfully",
        "data":    patient,
    })
}

// UpdatePatient - PUT /patients/:id
func UpdatePatient(c *gin.Context) {
    id := c.Param("id")
    var patient entity.Patient

    db := config.DB()

    // Find the patient by ID
    if err := db.First(&patient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }

    // Bind the updated data to the patient struct
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the updated patient data
    if err := db.Save(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Patient updated successfully",
        "data":    patient,
    })
}

// DeletePatient - DELETE /patients/:id
func DeletePatient(c *gin.Context) {
    id := c.Param("id")
    db := config.DB()

    // Delete the patient by ID
    if err := db.Delete(&entity.Patient{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Patient deleted successfully",
    })
}
