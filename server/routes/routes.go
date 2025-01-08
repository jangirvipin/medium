package routes

import (

	"medium/server/handlers"
	"net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/api/user/", handlers.UserHandle)
    http.HandleFunc("/api/blog/", handlers.BlogHandle)
	
}
