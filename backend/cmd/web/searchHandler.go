package main

import (
	"encoding/json"
	"errors"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *Application) OneWaySearch(w http.ResponseWriter, r *http.Request) {
	flights,_, err := app.FlightSearch(r)
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

func (app *Application) MultiSearch(w http.ResponseWriter, r *http.Request) {
	flights, layoversString, err := app.FlightSearch(r)
	if err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	layovers := strings.Split(layoversString, ";")
	if len(flights) == 0{}
	var retFlights []models.Flight = make([]models.Flight,0)
	var requiredLayoverFlag bool
	var flightFlag bool

	for _, flight := range flights{
		flightFlag = true
		for _, placeToVisit := range layovers{
			requiredLayoverFlag = false
			for _, layover := range flight.Layovers{
				if strings.Contains(strings.ToLower(placeToVisit), strings.ToLower(layover.Address.Address)){
					requiredLayoverFlag = true
				}
			}
			if requiredLayoverFlag == false{
				flightFlag = false
				break
			}
		}
		if flightFlag == true{
			retFlights = append(retFlights, flight)
		}
	}

	err = json.NewEncoder(w).Encode(retFlights)
	if err != nil {
		app.ErrorLog.Println("Cannot encode flights into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) FlightSearch(r *http.Request) ([]models.Flight, string, error){
	var searchQueryDto models.OneWayQueryDto
	var searchQuery models.OneWayQuery
	var flights []models.Flight

	err := json.NewDecoder(r.Body).Decode(&searchQueryDto)
	if err != nil {
		err = errors.New("Could not decode JSON")
		return nil,"", err
	}

	passengers, err:= strconv.ParseUint(searchQueryDto.Passengers, 10, 32)
	if err != nil {
		err = errors.New(searchQueryDto.Passengers + " is not a valid number")
		return nil,"", err
	}

	dateUnixInt, err := strconv.ParseInt(searchQueryDto.Date, 10, 64)
	if err != nil {
		err = errors.New(searchQueryDto.Date + " is not a valid date")
		return nil,"", err
	}
	var dateUnix time.Time = time.Unix(0, dateUnixInt*int64(time.Millisecond))
	searchQuery.From = strings.ToLower(searchQueryDto.From)
	searchQuery.To = strings.ToLower(searchQueryDto.To)
	searchQuery.Passengers = uint(passengers)
	searchQuery.Date = string([]rune(dateUnix.String())[0:10])
	searchQuery.SeatClass = strings.ToUpper(searchQueryDto.SeatClass)
	flights, err = app.Store.OneWaySearch(searchQuery)
	return flights,searchQueryDto.Layovers, nil
}