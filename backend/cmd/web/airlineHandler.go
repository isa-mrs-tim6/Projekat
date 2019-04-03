package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
)

func (app *Application) GetAirline(w http.ResponseWriter, r *http.Request) {
	airlines, err := app.Store.GetAirlines()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(airlines)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode airlines into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (app *Application) GetAirlineProfiles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	airlineProfile, err := app.Store.GetAirlineProfile(uint(id))
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(airlineProfile)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode airline profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateAirlineProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var airlineProfile models.AirlineProfile

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&airlineProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateAirline(uint(id), airlineProfile)
	if err != nil {
		app.ErrorLog.Printf("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
