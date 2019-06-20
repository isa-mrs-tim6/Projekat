package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
)

func (app *Application) GetAirlines(w http.ResponseWriter, r *http.Request) {
	airlines, err := app.Store.GetAirlines()
	if err != nil {
		app.ErrorLog.Println("Could not retrieve airlines")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(airlines)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airlines into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlinesProfiles(w http.ResponseWriter, r *http.Request) {
	airlines, err := app.Store.GetAirlinesProfiles()
	if err != nil {
		app.ErrorLog.Println("Could not retrieve airlines profiles")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(airlines)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airlines into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineProfileID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		app.ErrorLog.Println("Id is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile, err := app.Store.GetAirlineProfile(uint(id))
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airline profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineRatingID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		app.ErrorLog.Println("Id is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rating, err := app.Store.GetAirlineRating(uint(id))
	err = json.NewEncoder(w).Encode(rating)
	if err != nil {
		app.ErrorLog.Println("Cannot encode rating into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineQuickReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		app.ErrorLog.Println("Id is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reservations, err := app.Store.GetCompanyQuickReservations(uint(id))
	var flight models.Flight
	for index, element := range reservations {
		flight, _ = app.Store.GetFlight(uint(element.FlightID))
		reservations[index].DestName = flight.Destination.Name
		reservations[index].OriginName = flight.Origin.Name
	}
	err = json.NewEncoder(w).Encode(reservations)
	if err != nil {
		app.ErrorLog.Println("Cannot encode reservation into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CreateAirline(w http.ResponseWriter, r *http.Request) {
	var airline models.Airline

	err := json.NewDecoder(r.Body).Decode(&airline.AirlineProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.CreateAirline(&airline)
	if err != nil {
		app.ErrorLog.Println("Could not add airline to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) GetAirlineProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	airlineProfile, err := app.Store.GetAirlineProfile(user.AirlineID)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(airlineProfile)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airline profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateAirlineProfile(w http.ResponseWriter, r *http.Request) {
	var airlineProfile models.AirlineProfile
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	err = json.NewDecoder(r.Body).Decode(&airlineProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateAirline(user.AirlineID, airlineProfile)
	if err != nil {
		app.ErrorLog.Println("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) GetFlightRatings(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	flights, err := app.Store.GetFlightRatings(user.AirlineID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive flights and their ratings")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(flights)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode flights and their ratings into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineReservations(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	reservations, err := app.Store.GetAirlineReservations(user.AirlineID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline reservations")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(reservations)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode airline reservations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
