package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func signup_hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signup")
}

func signin_hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signin")
}

func Post_Blog_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Blog")
}

func Put_Blog_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Put Blog")
}

func Get_Blog_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Blog")
}

func main() {
	port := ":8080"

	http.HandleFunc("/signup", signup_hanlder)
	http.HandleFunc("/signin", signin_hanlder)
	http.HandleFunc("/api/v1/blog", Post_Blog_Handler)
	http.HandleFunc("/api/v1/blog", Put_Blog_Handler)
	http.HandleFunc("/api/v1/blog/:id", Get_Blog_Handler)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting server")
		return
	}

}
