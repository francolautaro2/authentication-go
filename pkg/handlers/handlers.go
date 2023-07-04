package handlers

import (
	"auth-go/pkg/token"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Simulate database of users
// WIP redis database
var users = map[string]string{
	"franco":         "admin",
	"franco2":        "1234",
	"francolautaro1": "franco123",
	"user1":          "password1",
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
	}

	getPassword, ok := users[u.Username]

	if !ok || getPassword != u.Password {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You not authorized")
		return
	}

	// Generate JWT token
	token, err := token.GenerateToken(u.Username)
	if err != nil {
		fmt.Println(err)
	}

	// Set token in a Cookie
	t := http.Cookie{
		Name:     "Token",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}
	http.SetCookie(w, &t)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)

}

// Sign up handler
func Signup(w http.ResponseWriter, r *http.Request) {

}

// Home handler example
func Home(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "welcome to home")
}

func MyProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO YOU ARE LOGGED")
}
