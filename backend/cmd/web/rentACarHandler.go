package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"strconv"
)

func (app *Application) GetRentACarCompanies(w http.ResponseWriter, r *http.Request) {
	racCompanies, err := app.Store.GetRentACarCompanies()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive rent-a-car companies")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(racCompanies)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode rent-a-car companies into JSON object")
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
		app.ErrorLog.Printf("Could not add hotel to database")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) GetRentACarCompanyProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel admin")
		w.WriteHeader(http.StatusInternalServerError)
	}

	racCompanyProfile, err := app.Store.GetRentACarCompanyProfile(user.RentACarCompanyID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrieve rent-a-car company profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(racCompanyProfile)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode rent-a-car company profile into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateRentACarCompanyProfile(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive rent-a-car admin")
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
		app.ErrorLog.Printf("Could not add rent-a-car company to database")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) FindVehicles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var params models.FindVehicleParams
	var vehicles []models.Vehicle

	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		app.ErrorLog.Printf(vars["id"] + "is not a valid ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vehicles, err = app.Store.FindVehicles(uint(id), params)
	if err != nil {
		app.ErrorLog.Printf("error while searching for vehicles")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(vehicles)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode hotels into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
