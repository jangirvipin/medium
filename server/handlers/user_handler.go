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

type userLogin struct{
    Email string `json:"email"`
    Password string `json:"password"`
}
// SetDB assigns the database instance to handlers
func SetDB(database *gorm.DB) {
	db = database
}

func UserHandle(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/api/user/signup":
        CreateUser(w, r)
    case "/api/user/signin":
        SigninUser(w, r)
    default:
        http.Error(w, "Not Found", http.StatusNotFound)
    }
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
    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseUser)
}


func SigninUser(w http.ResponseWriter, r *http.Request) {
    if r.Method!="POST"{
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var cred userLogin

    err:=json.NewDecoder(r.Body).Decode(&cred)
    if err!=nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    res := db.Where("email = ?", cred.Email).First(&user)
    if res.Error!=nil{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"msg": "User not found"})
        return
    }

    //Unhashing the password
    err=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

    if err!=nil{
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"msg": "Invalid password"})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"msg": "User signed in successfully"})
}
