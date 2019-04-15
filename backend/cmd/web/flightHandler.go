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

func (app *Application) GetCompanysAirplanes(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline admin")
	}
	airplanes, err := app.Store.GetCompanysAirplanes(user.AirlineID)

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
	smallSuitcase, _:= strconv.ParseFloat(flightDto.SmallSuitcase, 64)
	bigSuitcase, _:= strconv.ParseFloat(flightDto.BigSuitcase, 64)
	duration, _ := strconv.ParseInt(flightDto.Duration, 10, 64)
	distance, _ := strconv.ParseUint(flightDto.Distance, 10, 64)
	departure, _ := strconv.ParseInt(flightDto.Departure, 10, 64)
	landing, _ := strconv.ParseInt(flightDto.Landing, 10, 64)

	var layovers []models.Layovers = make([]models.Layovers, 0)
	var newLayover models.Layovers
	for _, layover := range flightDto.Layovers {
		newLayover = models.Layovers{
			Address: models.Address{
				Address: layover.Address,
			},
		}
		layovers = append(layovers, newLayover)
	}

	flight := models.Flight{
		PriceList: models.PriceList{
			PriceECONOMY:    priceECONOMY, PriceBUSINESS:   priceBUSINESS, PriceFIRSTCLASS: priceFIRSTCLASS,
			SmallSuitcase: smallSuitcase,  BigSuitcase: bigSuitcase,
		},
		Duration:        time.Minute * time.Duration(duration),
		Distance:        uint(distance),
		Departure:       time.Unix(0, departure*int64(time.Millisecond)),
		Landing:         time.Unix(0, landing*int64(time.Millisecond)),
		Layovers:        layovers,
	}
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline admin")
	}
	airplane, err := app.Store.GetAirplane(user.AirlineID, flightDto.Airplane)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airplane")
	}

	DeepCopy(airplane, &flight.Airplane)
	flight.AirlineID = airplane.ID

	origin, err := app.Store.FindDestination(user.AirlineID, flightDto.Origin)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive destination")
	}
	flight.Origin = &origin
	flight.OriginID = flight.Origin.ID

	destination, err := app.Store.FindDestination(user.AirlineID, flightDto.Destination)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive destination")
	}
	flight.Destination = &destination
	flight.DestinationID = flight.Destination.ID
	flight.AirlineID = user.AirlineID
	err = app.Store.CreateFlight(&flight)
	if err != nil {
		app.ErrorLog.Printf("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
