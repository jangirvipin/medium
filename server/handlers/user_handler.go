package handlers
import (

	"net/http"

	"gorm.io/gorm"
)


var db *gorm.DB

// SetDB assigns the database instance to handlers
func SetDB(database *gorm.DB) {
	db = database
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

// ListUsers lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}
