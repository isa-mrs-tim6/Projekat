package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"
)

func (app *Application) CompleteRegistration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := app.Store.CompleteRegistration(vars["email"], vars["type"]); err != nil {
		app.ErrorLog.Printf("Could not complete registration")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	type RequestJSON struct {
		models.Credentials
		AccountType string
		CompanyID   uint
	}
	var requestObject RequestJSON

	// DECODE REQUEST
	err := json.NewDecoder(r.Body).Decode(&requestObject)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// CHECK VALIDITY
	if requestObject.Email == "" || requestObject.Password == "" {
		app.ErrorLog.Println("Request fields are empty")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	emailAddress, emailError := mail.ParseAddress(requestObject.Email)
	if emailError != nil {
		app.ErrorLog.Println("Invalid email format")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	requestObject.AccountType = strings.Join(strings.Fields(requestObject.AccountType), "")
	if requestObject.AccountType != "SystemAdmin" && requestObject.CompanyID == 0 {
		app.ErrorLog.Println("Invalid company")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// ADD NEW ADMIN TO DATABASE
	err = app.Store.RegisterAdmin(requestObject.Credentials, requestObject.AccountType, requestObject.CompanyID)
	if err != nil {
		app.ErrorLog.Printf("Could not add system admin to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SEND CONFIRMATION EMAIL
	go app.SendRegistrationEmail(emailAddress.Address, requestObject.AccountType, requestObject.Credentials)

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) SendRegistrationEmail(receiver string, accountType string, credentials models.Credentials) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := "From: " + app.EmailAddress + "\n" +
		"To: " + receiver + "\n" +
		"Subject: Registration\n" +
		mime +
		`<html><head></head>
		<body><h1>ISA/MRS TIM6</h1>
		Complete your registration <a href="http://localhost:8080/confirmRegistration/` + accountType + `/` + receiver + `/">here</a>
		Your login details are: <br>
		EMAIL: ` + credentials.Email + `<br>
		PASSWORD: ` + credentials.Password + `
		</body></html>`

	_ = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
}

func (app *Application) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// DECODE REQUEST
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// CHECK VALIDITY
	if user.Email == "" || user.Password == "" ||
		user.Name == "" || user.Surname == "" ||
		user.Address == "" || user.Phone == "" {
		app.ErrorLog.Println("Request fields are empty")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	emailAddress, emailError := mail.ParseAddress(user.Email)

	if emailError != nil {
		app.ErrorLog.Println("Invalid email format")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// ADD NEW USER TO DATABASE
	err = app.Store.RegisterUser(user)
	if err != nil {
		app.ErrorLog.Printf("Could not add user to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SEND EMAIL
	go app.SendRegistrationEmail(emailAddress.Address, "User", user.Credentials)
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) ResendEmail(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		models.Credentials
		Type string
	}
	var query Query

	// DECODE REQUEST
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch query.Type {
	case "SystemAdmin":
		admin, err := app.Store.GetSystemAdmin(query.Email)
		if err != nil {
			app.ErrorLog.Println("No such system admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.SendRegistrationEmail(query.Email, "SystemAdmin", admin.Credentials)
	case "AirlineAdmin":
		admin, err := app.Store.GetAirlineAdmin(query.Email)
		if err != nil {
			app.ErrorLog.Println("No such airline admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.SendRegistrationEmail(query.Email, "AirlineAdmin", admin.Credentials)
	case "HotelAdmin":
		admin, err := app.Store.GetHotelAdmin(query.Email)
		if err != nil {
			app.ErrorLog.Println("No such hotel admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.SendRegistrationEmail(query.Email, "HotelAdmin", admin.Credentials)
	case "Rent-A-CarAdmin":
		admin, err := app.Store.GetRACAdmin(query.Email)
		if err != nil {
			app.ErrorLog.Println("No such rent-a-car admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.SendRegistrationEmail(query.Email, "Rent-A-CarAdmin", admin.Credentials)
	case "User":
		admin, err := app.Store.GetUser(query.Email)
		if err != nil {
			app.ErrorLog.Println("No such user")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		app.SendRegistrationEmail(query.Email, "User", admin.Credentials)
	}
}
