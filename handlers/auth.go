package handlers

import (
    "go-auth-app/models"
    "html/template"
    "net/http"
)

var templates = template.Must(template.ParseFiles("templates/register.html", "templates/login.html"))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        templates.ExecuteTemplate(w, "register.html", nil)
        return
    }
    username := r.FormValue("username")
    password := r.FormValue("password")
    if err := models.Register(username, password); err != nil {
        http.Error(w, "Username already taken", http.StatusConflict)
        return
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        templates.ExecuteTemplate(w, "login.html", nil)
        return
    }
    username := r.FormValue("username")
    password := r.FormValue("password")
    if models.Authenticate(username, password) {
        http.Redirect(w, r, "/welcome", http.StatusSeeOther)
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}
