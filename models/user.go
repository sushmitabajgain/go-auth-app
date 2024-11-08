package models

import "errors"

type User struct {
    Username string
    Password string
}

var users = map[string]User{}

// Register a new user
func Register(username, password string) error {
    if _, exists := users[username]; exists {
        return errors.New("username already exists")
    }
    users[username] = User{Username: username, Password: password}
    return nil
}

// Authenticate a user
func Authenticate(username, password string) bool {
    user, exists := users[username]
    return exists && user.Password == password
}
