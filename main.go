package main

import (
    "go-auth-app/handlers"
    "net/http"
)

func main() {
    // Serve static files (optional, for CSS/JS)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Set up routes
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)

    // Start the web server
    http.ListenAndServe(":8080", nil)
}
