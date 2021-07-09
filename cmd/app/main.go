package main

import (
	"github.com/barbibrussa/tiro-federal/pkg/models"
	"github.com/barbibrussa/tiro-federal/pkg/server"
	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	r := chi.NewRouter()

	s := server.NewServer(db)

	r.Post("/members", s.CreateMember)
	r.Get("/members", s.ListMembers)
	r.Delete("/members/{id}", s.DeleteMember)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error serving HTPP on port :8080 ", err)
	}
}
