package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
)

func (app *Application) GetHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := app.Store.GetHotels()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(hotels)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotels into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel

	err := json.NewDecoder(r.Body).Decode(&hotel.HotelProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.CreateHotel(&hotel)
	if err != nil {
		app.ErrorLog.Printf("Could not add hotel to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) GetHotelProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hotelProfile, err := app.Store.GetHotelProfile(uint(id))
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(hotelProfile)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotel profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateHotelProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var hotelProfile models.HotelProfile

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&hotelProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateHotel(uint(id), hotelProfile)
	if err != nil {
		app.ErrorLog.Printf("Could not add hotel to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
