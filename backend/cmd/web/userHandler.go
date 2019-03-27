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
		app.ErrorLog.Printf("Could not add user to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}