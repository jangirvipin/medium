package handlers

import (
	"encoding/json"
	"medium/server/models"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type userBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SetDB assigns the database instance to handlers
func SetDB(database *gorm.DB) {
	db = database
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user userBody
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userCreation := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword), // Hash the password for security
	}

	result := db.Create(&userCreation)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"msg": "Email already exists"})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"msg": result.Error.Error()})
		}
		return
	}

	responseUser := map[string]interface{}{
		"id":    userCreation.ID,
		"name":  userCreation.Name,
		"email": userCreation.Email,
		
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUser)
}

// ListUsers lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}
