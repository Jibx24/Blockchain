package entity

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	NationalID string `json:"national_id"`

	FirstName string `json:"first_name"`
	
	LastName string `json:"last_name"`

	Gender string `json:"gender"`

	Age uint8 `json:"age"`

	Appointment []Appointment `gorm:"foreignKey:PatientID"`
}
