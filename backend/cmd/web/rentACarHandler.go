package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) GetRentACarCompanies(w http.ResponseWriter, r *http.Request) {
	racCompanies, err := app.Store.GetRentACarCompanies()
	if err != nil {
		app.ErrorLog.Println("Could not retrive rent-a-car companies")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(racCompanies)
	if err != nil {
		app.ErrorLog.Println("Cannot encode rent-a-car companies into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CreateRentACarCompany(w http.ResponseWriter, r *http.Request) {
	var rac models.RentACarCompany

	err := json.NewDecoder(r.Body).Decode(&rac.RentACarCompanyProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.CreateRentACarCompany(&rac)
	if err != nil {
		app.ErrorLog.Println("Could not add rent-a-car company to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) GetRentACarCompanyProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive rent-a-car admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	racCompanyProfile, err := app.Store.GetRentACarCompanyProfile(user.RentACarCompanyID)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve rent-a-car company profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(racCompanyProfile)
	if err != nil {
		app.ErrorLog.Println("Cannot encode rent-a-car company profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateRentACarCompanyProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive rent-a-car admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var racCompanyProfile models.RentACarCompanyProfile

	err = json.NewDecoder(r.Body).Decode(&racCompanyProfile)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.UpdateRentACarCompanyProfile(user.RentACarCompanyID, racCompanyProfile)
	if err != nil {
		app.ErrorLog.Println("Could not add rent-a-car company to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) RacSearch(w http.ResponseWriter, r *http.Request) {
	var searchQuery models.RacQueryDTO
	var query models.RacQuery

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

	query.Name = searchQuery.Name
	query.Address = searchQuery.Address
	query.From = dateFrom
	query.To = dateTo

	if hotels, err := app.Store.RacSearch(query); err != nil {
		app.ErrorLog.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(hotels); err != nil {
			app.ErrorLog.Println("Cannot encode hotels into JSON object")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (app *Application) FindVehicles(w http.ResponseWriter, r *http.Request) {
	var params models.FindVehicleParams
	var vehicles []models.VehicleRatingsDAO

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vehicles, err = app.Store.FindVehicles(params)
	if err != nil {
		app.ErrorLog.Println("error while searching for vehicles")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vehicles)
	if err != nil {
		app.ErrorLog.Println("Cannot encode vehicles into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) GetCompanyVehicles(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve rent-a-car admin")
	}
	vehicles, err := app.Store.GetCompanyVehicles(user.RentACarCompanyID)

	err = json.NewEncoder(w).Encode(vehicles)
	if err != nil {
		app.ErrorLog.Println("Cannot encode vehicles into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetVehicleRatings(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive rac admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	v, err := app.Store.GetVehicleRatings(user.RentACarCompanyID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive vehicles and their ratings")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(v)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode vehicles and their ratings into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetRACReservations(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive rac admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	rooms, err := app.Store.GetRACReservations(user.RentACarCompanyID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive rac reservations")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode rac reservations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var newVehicle models.Vehicle
	var params models.VehicleParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newVehicle.Name = params.Name
	newVehicle.Capacity = params.Capacity
	newVehicle.Type = params.Type
	newVehicle.PricePerDay = params.PricePerDay
	newVehicle.Discount = params.Discount

	err = app.Store.UpdateVehicle(params.ID, newVehicle)
	if err != nil {
		app.ErrorLog.Println("Could not update vehicle in database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.DeleteVehicle(params.ID)
	if err != nil {
		app.ErrorLog.Println("Could not delete vehicle from database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) AddVehicle(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive rac admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var params models.VehicleParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newVehicle models.Vehicle

	newVehicle.Name = params.Name
	newVehicle.Capacity = params.Capacity
	newVehicle.Type = params.Type
	newVehicle.PricePerDay = params.PricePerDay
	newVehicle.Discount = params.Discount
	newVehicle.RentACarCompanyID = user.RentACarCompanyID

	err = app.Store.AddVehicle(newVehicle)
	if err != nil {
		app.ErrorLog.Println("Could not add new vehicle to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetCompanyLocations(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive rent-a-car admin")
	}
	locations, err := app.Store.GetCompanyLocations(user.RentACarCompanyID)

	err = json.NewEncoder(w).Encode(locations)
	if err != nil {
		app.ErrorLog.Println("Cannot encode locations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var newOffice models.Location
	var params models.LocationParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newOffice.Address = params.Address
	newOffice.Longitude = params.Longitude
	newOffice.Latitude = params.Latitude

	err = app.Store.UpdateLocation(params.ID, newOffice)
	if err != nil {
		app.ErrorLog.Println("Could not update location in database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	var params models.LocationParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.DeleteLocation(params.ID)
	if err != nil {
		app.ErrorLog.Println("Could not delete location from database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) AddLocation(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive rac admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var params models.LocationParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var newLocation models.Location

	newLocation.Address = params.Address
	newLocation.Longitude = params.Longitude
	newLocation.Latitude = params.Latitude
	newLocation.RentACarCompanyID = user.RentACarCompanyID

	err = app.Store.AddLocation(newLocation)
	if err != nil {
		app.ErrorLog.Println("Could not add new location to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
