package routes

import (
    "net/http"
    "medium/server/handlers"
)

func RegisterRoutes() {
    http.HandleFunc("/api/v1/signup", handlers.CreateUser)
    http.HandleFunc("/api/v1/signin", handlers.SigninUser)
}
