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
		app.ErrorLog.Println("Could not retrive list of airplanes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airplanes into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdatePriceList(w http.ResponseWriter, r *http.Request) {
	var flightDTO models.FlightDto
	var flight models.Flight

	err := json.NewDecoder(r.Body).Decode(&flightDTO)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	priceECONOMY, err := strconv.ParseFloat(flightDTO.PriceECONOMY, 64)
	if err != nil {
		app.ErrorLog.Println("Economy class price is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceBUSINESS, err := strconv.ParseFloat(flightDTO.PriceBUSINESS, 64)
	if err != nil {
		app.ErrorLog.Println("Business class price is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceFIRSTCLASS, err := strconv.ParseFloat(flightDTO.PriceFIRSTCLASS, 64)
	if err != nil {
		app.ErrorLog.Println("First class price is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	smallSuitcase, err:= strconv.ParseFloat(flightDTO.SmallSuitcase, 64)
	if err != nil {
		app.ErrorLog.Println("Small suitacase price is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bigSuitcase, err:= strconv.ParseFloat(flightDTO.BigSuitcase, 64)
	if err != nil {
		app.ErrorLog.Println("Big suitacase price is not a valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err:= strconv.ParseUint(flightDTO.FlightID, 10, 32)
	if err != nil {
		app.ErrorLog.Println(flightDTO.FlightID + " is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flight, err = app.Store.GetFlight(uint(id))
	if err != nil {
		app.ErrorLog.Println("Could not retrive flight")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	flight.PriceBUSINESS = priceBUSINESS
	flight.PriceECONOMY = priceECONOMY
	flight.PriceFIRSTCLASS = priceFIRSTCLASS
	flight.BigSuitcase = bigSuitcase
	flight.SmallSuitcase = smallSuitcase

	err = app.Store.UpdateFlight(uint(id),flight)
	if err != nil {
		app.ErrorLog.Println("Could not update flight in database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) UpdateSeats(w http.ResponseWriter, r *http.Request) {
	var flightDTO models.FlightDto;
	var flight models.Flight;
	err := json.NewDecoder(r.Body).Decode(&flightDTO)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err:= strconv.ParseUint(flightDTO.FlightID, 10, 32)
	if err != nil {
		app.ErrorLog.Printf(flightDTO.FlightID + " is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flight, err = app.Store.GetFlight(uint(id))
	if err != nil {
		app.ErrorLog.Println("Could not retrive flight")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flight.Airplane.Seats = flightDTO.AirplaneObject.Seats;
	err = app.Store.UpdateFlight(uint(flight.ID),flight)
	if err != nil {
		app.ErrorLog.Println("Could not update flight in database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) GetCompanyAirplanes(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	airplanes, err := app.Store.GetCompanyAirplanes(user.AirlineID)

	err = json.NewEncoder(w).Encode(airplanes)
	if err != nil {
		app.ErrorLog.Println("Cannot encode airplanes into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func(app *Application) GetCompanyFlights(w http.ResponseWriter, r *http.Request){
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil{
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	flights, err := app.Store.GetCompanyFlights(user.AirlineID)
	if err != nil{
		app.ErrorLog.Println("Could not retrive flights")
	}

	err = json.NewEncoder(w).Encode(flights)
	if err != nil{
		app.ErrorLog.Println("Cannot encode flights into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
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
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	airplane, err := app.Store.GetAirplane(user.AirlineID, flightDto.Airplane)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airplane")
	}

	DeepCopy(airplane, &flight.Airplane)
	flight.AirlineID = airplane.ID

	origin, err := app.Store.FindDestination(user.AirlineID, flightDto.Origin)
	if err != nil {
		app.ErrorLog.Println("Could not retrive destination")
	}
	flight.Origin = &origin
	flight.OriginID = flight.Origin.ID

	destination, err := app.Store.FindDestination(user.AirlineID, flightDto.Destination)
	if err != nil {
		app.ErrorLog.Println("Could not retrive destination")
	}
	flight.Destination = &destination
	flight.DestinationID = flight.Destination.ID
	flight.AirlineID = user.AirlineID
	err = app.Store.CreateFlight(&flight)
	if err != nil {
		app.ErrorLog.Println("Could not add destination to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
