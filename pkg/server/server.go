package server

import (
	"encoding/json"
	"github.com/barbibrussa/tiro-federal/pkg/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Server struct {
	db *gorm.DB
}

func (s *Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var member models.Member

	err = json.Unmarshal(body, &member)
	if err != nil {
		http.Error(w, "Failed to unmarshal request", http.StatusInternalServerError)
		return
	}

	err = s.db.Model(&models.Member{}).Save(&member).Error
	if err != nil {
		http.Error(w, "Failed to create member in database", http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(member)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(payload)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) ListMembers(w http.ResponseWriter, r *http.Request) {

	var list []models.Member

	err := s.db.Model(&models.Member{}).Find(&list).Error
	if err != nil {
		http.Error(w, "Failed to list members from database", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
