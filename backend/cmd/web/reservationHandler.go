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