package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Type     string `json:"type"`
	jwt.StandardClaims
}

func (app *Application) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.Login(credentials)
	if err != nil {
		app.ErrorLog.Printf("Invalid login")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: credentials.Email,
		Type:     "User",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusCreated)
}
