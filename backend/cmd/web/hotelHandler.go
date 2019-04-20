package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"io/ioutil"
	"net/http"
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
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	hotelProfile, err := app.Store.GetHotelProfile(user.HotelID)
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
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}
	var hotelProfile models.HotelProfile

	err = json.NewDecoder(r.Body).Decode(&hotelProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateHotel(user.HotelID, hotelProfile)
	if err != nil {
		app.ErrorLog.Printf("Could not add hotel to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) GetRooms(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	rooms, err := app.Store.GetRooms(user.HotelID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel rooms")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotel rooms into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetRoomRatings(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	rooms, err := app.Store.GetRoomRatings(user.HotelID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel rooms and their ratings")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotel rooms and their ratings into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetHotelReservations(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	rooms, err := app.Store.GetHotelReservations(user.HotelID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel reservations")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotel reservations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateRooms(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var rooms []models.Room
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &rooms)
	} else {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateRooms(user.HotelID, rooms)
	if err != nil {
		app.ErrorLog.Println("Could not add new rooms to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) DeleteRooms(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var rooms []models.Room
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &rooms)
	} else {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.DeleteRooms(user.HotelID, rooms)
	if err != nil {
		app.ErrorLog.Printf(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
	}
}
