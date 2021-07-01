package main

import (
	"github.com/barbibrussa/tiro-federal/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("members.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
	}

	err = db.AutoMigrate(&models.Member{})
	if err != nil {
		log.Fatal("Error while auto-migrating models: ", err)
	}
}
