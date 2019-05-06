package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"time"
)

func (app *Application) GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := app.Store.GetUsers()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode users into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)

	userProfile, err := app.Store.GetUserProfile(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userProfile)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode user profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	var params models.ProfileParams

	oldEmail := getEmail(r)

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateUser(oldEmail, params)
	if err != nil {
		app.ErrorLog.Printf("Could not update user profile")
		w.WriteHeader(http.StatusInternalServerError)
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: params.Email,
		Type:  "User",
		StandardClaims: jwt.StandardClaims{
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
}

func (app *Application) GetAdminProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	accountType := getAccountType(r)

	switch accountType {
	case "AirlineAdmin":
		admin, err := app.Store.GetAirlineAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive airline admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode user profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "HotelAdmin":
		admin, err := app.Store.GetHotelAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive hotel admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode user profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "Rent-A-CarAdmin":
		admin, err := app.Store.GetRACAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive hotel admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode user profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "SystemAdmin":
		admin, err := app.Store.GetSystemAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive hotel admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode user profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		app.ErrorLog.Printf("Invalid user type")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
}

func (app *Application) UpdateAdminProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	accountType := getAccountType(r)

	var profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch accountType {
	case "AirlineAdmin":
		// TODO
	case "HotelAdmin":
		admin, err := app.Store.GetHotelAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive hotel admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = app.Store.UpdateHotelAdmin(admin.ID, profile)
		if err != nil {
			app.ErrorLog.Printf("Could not update hotel admin profile")
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode hotel admin profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "Rent-A-CarAdmin":
		admin, err := app.Store.GetRACAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive hotel admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = app.Store.UpdateRACAdmin(admin.ID, profile)
		if err != nil {
			app.ErrorLog.Printf("Could not update rac admin profile")
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode hotel admin profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "SystemAdmin":
		admin, err := app.Store.GetSystemAdmin(email)
		if err != nil {
			app.ErrorLog.Printf("Could not retrive system admin")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = app.Store.UpdateSystemAdmin(admin.ID, profile)
		if err != nil {
			app.ErrorLog.Printf("Could not update system profile")
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(admin.Profile)
		if err != nil {
			app.ErrorLog.Printf("Cannot encode hotel admin profile into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		app.ErrorLog.Printf("Invalid user type")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
}
