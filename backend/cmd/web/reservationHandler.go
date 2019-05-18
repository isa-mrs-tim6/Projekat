package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"io/ioutil"
	"net/http"
)

func (app *Application) GetRewards(w http.ResponseWriter, r *http.Request) {
	rewards, err := app.Store.GetReservationRewards()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation rewards")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rewards)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation rewards into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateRewards(w http.ResponseWriter, r *http.Request) {
	var rewards []models.ReservationReward
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &rewards)
	} else {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.Store.UpdateReservationRewards(rewards)
}

func (app *Application) ReserveVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleReservationParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.ReserveVehicle(params)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive vehicles and their ratings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetUserReservations(w http.ResponseWriter, r *http.Request) {
	var reservations []models.ReservationDAO

	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
	}

	reservations, err = app.Store.GetUserReservations(user.ID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservations")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reservations)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
