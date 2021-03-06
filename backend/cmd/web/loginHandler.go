package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"time"
)

func (app *Application) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.LoginUser(credentials)
	if err != nil {
		if err.Error() == "not activated" {
			app.ErrorLog.Printf("Account not activated")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.ErrorLog.Printf("Invalid login")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: credentials.Email,
		Type:  "User",
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
		Path:    "/",
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		models.Credentials
		Type string
	}
	var credentials Query

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	typeUser, err := app.Store.LoginAdmin(credentials.Credentials, credentials.Type)
	if err != nil {
		if err.Error() == "not activated" {
			app.ErrorLog.Printf("Account not activated")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.ErrorLog.Printf("Invalid login")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: credentials.Email,
		Type:  typeUser,
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
		Path:    "/",
		Expires: expirationTime,
	})

	err = json.NewEncoder(w).Encode(typeUser)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode admin type into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CheckFirstPass(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	accountType := getAccountType(r)

	changed, err := app.Store.CheckFirstPassChanged(email, accountType)
	if err != nil {
		app.ErrorLog.Printf("That email does not exist")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(changed)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode boolean into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
