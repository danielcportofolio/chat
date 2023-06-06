package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/your-package/chat-api/models"
	"github.com/your-package/chat-api/repositories"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uh.userRepo.CreateUser(&user)
	if err != nil {
		log.Printf("Failed to create a new user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
