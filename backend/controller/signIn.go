package controller

import (
	"net/http"

	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"blockchain.com/bc-67/services"
	"github.com/gin-gonic/gin"
)

type (
	Authen struct {
		NationalID string `json:"national_id"`
	}
)

func SignIn(c *gin.Context) {
	var payload Authen
	var patient entity.Patient
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search for the patient by NationalID
	if err := config.DB().Raw("SELECT * FROM patients WHERE national_id = ?", payload.NationalID).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the NationalID matches
	if patient.NationalID != payload.NationalID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NationalID is incorrect"})
		return
	}

	// Generate JWT token
	jwtWrapper := services.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(patient.NationalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": patient.ID})
}
