package routes

import (
    "net/http"
    "medium/server/handlers"
)

func RegisterRoutes() {
    http.HandleFunc("/users/create", handlers.CreateUser)
    http.HandleFunc("/users", handlers.ListUsers)
}
