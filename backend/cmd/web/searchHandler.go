package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *Application) OneWaySearch(w http.ResponseWriter, r *http.Request) {
	flights, _, err := app.FlightSearch(r)
	if err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(flights)
	if err != nil {
		app.ErrorLog.Println("Cannot encode flights into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) UserSearch(w http.ResponseWriter, r *http.Request) {
	var query models.UserQueryDto
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		err = errors.New("Could not decode JSON")
		return
	}
	var retUsers = make([]models.UserResultDTO, 0)
	retUsers, err = app.Store.UserSearch(user.ID, query)

	err = json.NewEncoder(w).Encode(retUsers)
	if err != nil {
		app.ErrorLog.Println("Cannot encode users into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) MultiSearch(w http.ResponseWriter, r *http.Request) {
	flights, layoversString, err := app.FlightSearch(r)
	if err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	layovers := strings.Split(layoversString, ";")
	if len(flights) == 0 {
	}
	var retFlights []models.Flight = make([]models.Flight, 0)
	var requiredLayoverFlag bool
	var flightFlag bool

	for _, flight := range flights {
		flightFlag = true
		for _, placeToVisit := range layovers {
			requiredLayoverFlag = false
			for _, layover := range flight.Layovers {
				if strings.Contains(strings.ToLower(placeToVisit), strings.ToLower(layover.Address.Address)) {
					requiredLayoverFlag = true
				}
			}
			if requiredLayoverFlag == false {
				flightFlag = false
				break
			}
		}
		if flightFlag == true {
			retFlights = append(retFlights, flight)
		}
	}

	err = json.NewEncoder(w).Encode(retFlights)
	if err != nil {
		app.ErrorLog.Println("Cannot encode flights into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) FlightSearch(r *http.Request) ([]models.Flight, string, error) {
	var searchQueryDto models.OneWayQueryDto
	var searchQuery models.OneWayQuery
	var flights []models.Flight

	err := json.NewDecoder(r.Body).Decode(&searchQueryDto)
	if err != nil {
		err = errors.New("Could not decode JSON")
		return nil, "", err
	}

	passengers, err := strconv.ParseUint(searchQueryDto.Passengers, 10, 32)
	if err != nil {
		err = errors.New(searchQueryDto.Passengers + " is not a valid number")
		return nil, "", err
	}

	dateUnixInt, err := strconv.ParseInt(searchQueryDto.Date, 10, 64)
	if err != nil {
		err = errors.New(searchQueryDto.Date + " is not a valid date")
		return nil, "", err
	}
	var dateUnix time.Time = time.Unix(0, dateUnixInt*int64(time.Millisecond))
	searchQuery.From = strings.ToLower(searchQueryDto.From)
	searchQuery.To = strings.ToLower(searchQueryDto.To)
	searchQuery.Passengers = uint(passengers)
	searchQuery.Date = string([]rune(dateUnix.String())[0:10])
	searchQuery.SeatClass = strings.ToUpper(searchQueryDto.SeatClass)
	flights, err = app.Store.OneWaySearch(searchQuery)
	return flights, searchQueryDto.Layovers, nil
}

func (app *Application) HotelSearch(w http.ResponseWriter, r *http.Request) {
	var searchQuery models.HotelQueryDTO
	var query models.HotelQuery

	if err := json.NewDecoder(r.Body).Decode(&searchQuery); err != nil {
		app.ErrorLog.Println("Cannot decode JSON object")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dateFromInt, err := strconv.ParseInt(searchQuery.From, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid from date")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateFrom := time.Unix(0, dateFromInt*int64(time.Millisecond))

	dateToInt, err := strconv.ParseInt(searchQuery.To, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid to date")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateTo := time.Unix(0, dateToInt*int64(time.Millisecond))

	query.Name = searchQuery.Name
	query.Address = searchQuery.Address
	query.From = dateFrom
	query.To = dateTo

	if hotels, err := app.Store.HotelSearch(query); err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(hotels); err != nil {
			app.ErrorLog.Println("Cannot encode hotels into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (app *Application) RoomSearch(w http.ResponseWriter, r *http.Request) {
	var searchQuery models.RoomQueryDTO
	var query models.RoomQuery

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// GET HOTEL ID
	vars := mux.Vars(r)
	hotelID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get hotel ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&searchQuery); err != nil {
		app.ErrorLog.Println("Cannot decode JSON object")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dateFromInt, err := strconv.ParseInt(searchQuery.From, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid from date")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateFrom := time.Unix(0, dateFromInt*int64(time.Millisecond))

	dateToInt, err := strconv.ParseInt(searchQuery.To, 10, 64)
	if err != nil {
		app.ErrorLog.Println("Invalid to date")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dateTo := time.Unix(0, dateToInt*int64(time.Millisecond))

	query.HotelID = uint(hotelID)
	query.Capacities = searchQuery.Capacities
	query.From = dateFrom
	query.To = dateTo

	if rooms, err := app.Store.RoomSearch(query, user.ID, false); err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(rooms); err != nil {
			app.ErrorLog.Println("Cannot encode rooms into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (app *Application) QuickReserveRoomSearch(w http.ResponseWriter, r *http.Request) {
	var searchQuery models.RoomQueryDTO
	var query models.RoomQuery

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// GET HOTEL ID
	vars := mux.Vars(r)
	hotelID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get hotel ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&searchQuery); err != nil {
		app.ErrorLog.Println("Cannot decode JSON object")
		w.WriteHeader(http.StatusBadRequest)
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

	query.HotelID = uint(hotelID)
	query.Capacities = searchQuery.Capacities
	query.From = dateFrom
	query.To = dateTo

	if rooms, err := app.Store.QuickReservationRoomSearch(query, user.ID, false); err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(rooms); err != nil {
			app.ErrorLog.Println("Cannot encode rooms into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
