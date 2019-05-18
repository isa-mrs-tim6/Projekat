package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
)

func (app *Application) GetFriends(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	var friends []models.User

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	friends, err = app.Store.GetFriends(user.ID)

	err = json.NewEncoder(w).Encode(friends)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode friends into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
