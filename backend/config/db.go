package config

import (
	"fmt"
	"time"

	"blockchain.com/bc-67/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	db.AutoMigrate(
		&entity.Patient{},
		&entity.Appointment{},
		&entity.Department{},
		&entity.Room{}, // เพิ่มการ migrate ตาราง Like
	)

	// ผู้ป่วย
	patients := []entity.Patient{
		{
			NationalID: "1111111111",
			FirstName: "J",
			LastName: "P",
			Gender: "หญิง",
			Age: 21,
		},
		{
			NationalID: "2222222222222",
			FirstName: "A",
			LastName: "P",
			Gender: "หญิง",
			Age: 21,
		},
	}

	for _, patient := range patients {
		db.FirstOrCreate(&patient, entity.Patient{NationalID: patient.NationalID})
	}

	// นัดหมาย
	appointments := []entity.Appointment{
		{
			Date: time.Now(),
			Time: time.Now(),
			Illness: "ปวดหัว",
			DepartmentID: 1,
			PatientID: 1,
		},
		{
			Date: time.Now(),
			Time: time.Now(),
			Illness: "ปวดตัว",
			DepartmentID: 2,
			PatientID: 2,
		},
	}

	for _, appointment := range appointments {
		db.FirstOrCreate(&appointment, entity.Appointment{Illness: appointment.Illness})
	}

	//แผนก
	departments := []entity.Department{
		{
			Name: "",
			Description: "",
			HeadOfDept: "",
		},
		{
			Name: "",
			Description: "",
			HeadOfDept: "",
		},
	}

	for _, department := range departments {
		db.FirstOrCreate(&department, entity.Department{HeadOfDept: department.HeadOfDept})
	}

	//ห้อง
	rooms := []entity.Room{
		{
			Name: "",
			DepartmentID: 1,
		},
		{
			Name: "",
			DepartmentID: 2,
		},
	}

	for _, room := range rooms {
		db.FirstOrCreate(&room, entity.Room{DepartmentID: room.DepartmentID})
	}
}
