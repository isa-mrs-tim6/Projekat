package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
)

func (app *Application) GetDestinations(w http.ResponseWriter, r *http.Request) {
	destinations, err := app.Store.GetDestinations()
	if err != nil {
		app.ErrorLog.Println("Could not retrive destinations")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(destinations)
	if err != nil {
		app.ErrorLog.Println("Cannot encode destination into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetCompanyDestinations(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	destinations, err := app.Store.GetCompanyDestinations(user.AirlineID)

	err = json.NewEncoder(w).Encode(destinations)
	if err != nil {
		app.ErrorLog.Println("Cannot encode destinations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	destination, err := app.Store.GetDestination(uint(id))
	if err != nil {
		app.ErrorLog.Println("Could not retrive destination")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(destination)
	if err != nil {
		app.ErrorLog.Println("Cannot encode destination into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var destination models.Destination

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&destination)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateDestination(uint(id), destination)
	if err != nil {
		app.ErrorLog.Println("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) CreateDestination(w http.ResponseWriter, r *http.Request) {
	var destination models.Destination
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	err = json.NewDecoder(r.Body).Decode(&destination)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	destination.AirlineID = user.AirlineID
	err = app.Store.CreateDestination(&destination)
	if err != nil {
		app.ErrorLog.Println("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
