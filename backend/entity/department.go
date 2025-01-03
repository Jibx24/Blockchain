package entity

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	HeadOfDept  string `json:"head_of_department"`
}
