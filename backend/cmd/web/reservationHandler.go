package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

func (app *Application) GetReservation(w http.ResponseWriter, r *http.Request) {
	// GET RESERVATION ID
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get flight ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// FIND RESERVATION AND CHECK IF IT BELONGS TO USER REQUESTING IT
	reservation, err := app.Store.GetReservation(uint(reservationID))
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if reservation.UserID != user.ID {
		app.ErrorLog.Printf("Cannot access this reservation")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(reservation)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CompleteQuickResVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.CompleteQuickResVehParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.CompleteQuickResVehicle(params)

	if err != nil {
		app.ErrorLog.Printf("Cannot complete quick reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetCompanyQuickVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleQuickResParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := app.Store.GetCompanyQuickVehicle(params)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetQuickVehRes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get vehicle ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := app.Store.GetQuickVehRes(uint(vehicleID))

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineGraphData(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	data, err := app.Store.GetReservationGraphData(user.ID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) ReserveVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleReservationParams

	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
	}

	// GET RESERVATION ID
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get reservation ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := app.Store.ReserveVehicle(uint(reservationID), params, user.ID); err != nil {
		app.ErrorLog.Printf("Could not complete reservation")
		w.WriteHeader(http.StatusInternalServerError)
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

func (app *Application) ReserveFlight(w http.ResponseWriter, r *http.Request) {
	// GET FLIGHT ID
	vars := mux.Vars(r)
	flightID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get flight ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DECODE QUERY
	var query models.FlightReservationParams
	err = json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// SET MASTER USER ID
	query.Users[0].ID = user.ID

	// RESERVE
	reservationID, err := app.Store.ReserveFlight(flightID, query)
	if err != nil {
		app.ErrorLog.Println("Could not complete reservation")
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode(reservationID); err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) ReserveHotel(w http.ResponseWriter, r *http.Request) {
	// GET MASTER REF ID, HOTEL ID
	vars := mux.Vars(r)
	masterID, err := strconv.ParseUint(vars["master_id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get master reservation ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hotelID, err := strconv.ParseUint(vars["hotel_id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get hotel ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DECODE QUERY
	var searchQuery models.HotelReservationParamsDTO
	err = json.NewDecoder(r.Body).Decode(&searchQuery)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dateFromInt, err := strconv.ParseInt(searchQuery.From, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid from date")
		w.WriteHeader(http.StatusBadRequest)
	}
	dateFrom := time.Unix(0, dateFromInt*int64(time.Millisecond))

	dateToInt, err := strconv.ParseInt(searchQuery.To, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid to date")
		w.WriteHeader(http.StatusBadRequest)
	}
	dateTo := time.Unix(0, dateToInt*int64(time.Millisecond))

	// CLEANUP QUERY
	var query models.HotelReservationParams
	query.Rooms = searchQuery.Rooms
	query.From = dateFrom
	query.To = dateTo
	query.IsQuickReserve = searchQuery.IsQuickReserve
	for k := range searchQuery.Features {
		query.Features = append(query.Features, &searchQuery.Features[k])
	}

	// RESERVE
	reservationID, err := app.Store.ReserveHotel(uint(masterID), uint(hotelID), user.ID, query)
	if err != nil {
		app.ErrorLog.Println("Could not complete reservation")
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode(reservationID); err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CancelFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelFlight(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) CancelHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelHotel(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) CancelVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelVehicle(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}
