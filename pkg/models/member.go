package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DNI       string `json:"dni"`
}
