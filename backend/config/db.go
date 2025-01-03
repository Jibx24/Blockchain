package config

import (
	"fmt"

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

}