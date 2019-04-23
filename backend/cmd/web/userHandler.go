package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
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
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userProfile, err := app.Store.GetUserProfile(uint(id))
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
	vars := mux.Vars(r)
	var userProfile models.Profile

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&userProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateUser(uint(id), userProfile)
	if err != nil {
		app.ErrorLog.Printf("Could not update user profile")
		w.WriteHeader(http.StatusInternalServerError)
	}
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
