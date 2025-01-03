package entity

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	Date time.Time `json:"date"`

	Time time.Time `json:"time"`

	Reason string `json:"reason"`

	DepartmentID uint
	Department   Department  `json:"department" gorm:"foreignKey:DepartmentID"`

	PatientID uint
	Patient   Patient `json:"patient" gorm:"foreignKey:PatientID"`
}
