package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) GetAirplanes(w http.ResponseWriter, r *http.Request) {
	airplanes, err := app.Store.GetAirplanes()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive list of airplanes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode airplanes into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CreateFlight(w http.ResponseWriter, r *http.Request) {
	var flightDto models.FlightDto

	err := json.NewDecoder(r.Body).Decode(&flightDto)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceECONOMY, _ := strconv.ParseFloat(flightDto.PriceECONOMY, 64)
	priceBUSINESS, _ := strconv.ParseFloat(flightDto.PriceBUSINESS, 64)
	priceFIRSTCLASS, _ := strconv.ParseFloat(flightDto.PriceFIRSTCLASS, 64)
	duration, _ := strconv.ParseInt(flightDto.Duration, 10, 64)
	distance, _ := strconv.ParseUint(flightDto.Distance, 10, 64)
	departure, _ := strconv.ParseInt(flightDto.Departure, 10, 64)
	landing, _ := strconv.ParseInt(flightDto.Landing, 10, 64)

	var layovers []models.Layovers
	var newLayover models.Layovers
	for index, layover := range flightDto.Layovers {
		newLayover = models.Layovers{
			Address: models.Address{
				Address: layover.Address,
			},
		}
		layovers[index] = newLayover
	}

	flight := models.Flight{
		PriceECONOMY:    priceECONOMY,
		PriceBUSINESS:   priceBUSINESS,
		PriceFIRSTCLASS: priceFIRSTCLASS,
		Duration:        time.Minute * time.Duration(duration),
		Distance:        uint(distance),
		Departure:       time.Unix(0, departure*int64(time.Millisecond)),
		Landing:         time.Unix(0, landing*int64(time.Millisecond)),
	}

	err = app.Store.CreateFlight(&flight)
	if err != nil {
		app.ErrorLog.Printf("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
