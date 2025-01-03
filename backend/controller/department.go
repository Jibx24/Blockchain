package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "blockchain.com/bc-67/config"
    "blockchain.com/bc-67/entity"
)

// CreateDepartment - POST /departments
func CreateDepartment(c *gin.Context) {
    var department entity.Department

    // Bind JSON data to the department struct
    if err := c.ShouldBindJSON(&department); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    // Create a new department entry
    if err := db.Create(&department).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create department"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Department created successfully",
        "data":    department,
    })
}

// GetAllDepartments - GET /departments
func GetAllDepartments(c *gin.Context) {
    var departments []entity.Department

    db := config.DB()

    // Fetch all departments
    results := db.Find(&departments)

    if results.Error != nil || results.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No departments found"})
        return
    }

    // Return all departments
    c.JSON(http.StatusOK, gin.H{
        "message": "Departments fetched successfully",
        "data":    departments,
    })
}

// GetDepartmentByID - GET /departments/:id
func GetDepartmentByID(c *gin.Context) {
    id := c.Param("id")
    var department entity.Department

    db := config.DB()

    // Fetch department by ID
    if err := db.First(&department, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
        return
    }

    // Return department data
    c.JSON(http.StatusOK, gin.H{
        "message": "Department fetched successfully",
        "data":    department,
    })
}

// UpdateDepartment - PUT /departments/:id
func UpdateDepartment(c *gin.Context) {
    id := c.Param("id")
    var department entity.Department

    db := config.DB()

    // Find the department by ID
    if err := db.First(&department, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
        return
    }

    // Bind the updated data to the department struct
    if err := c.ShouldBindJSON(&department); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the updated department data
    if err := db.Save(&department).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Department updated successfully",
        "data":    department,
    })
}

// DeleteDepartment - DELETE /departments/:id
func DeleteDepartment(c *gin.Context) {
    id := c.Param("id")
    db := config.DB()

    // Delete the department by ID
    if err := db.Delete(&entity.Department{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete department"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Department deleted successfully",
    })
}
