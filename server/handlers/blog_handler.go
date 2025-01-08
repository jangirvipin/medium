package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// BlogHandle handles blog-related API requests
func BlogHandle(w http.ResponseWriter, r *http.Request) {
   subpath:=r.URL.Path[len("/api/blog"):]
   fmt.Printf("subpath: %s\n", subpath)
   fmt.Printf("I am in blog handler\n")
    switch {
    case subpath=="/create":
        CreateBlog(w, r)
    case subpath=="/update":
        UpdateBlog(w, r)
    case subpath=="/delete":
        DeleteBlog(w, r)
    default:
}
}

// CreateBlog creates a new blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Blog created successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateBlog updates an existing blog
func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Blog"))
}

// DeleteBlog deletes an existing blog
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Blog"))
}
