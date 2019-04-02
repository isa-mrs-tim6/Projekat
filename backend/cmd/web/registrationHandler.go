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

	// ADD NEW SYSTEM ADMIN TO DATABASE
	err = app.Store.RegisterAdmin(requestObject.Credentials, requestObject.AccountType)
	if err != nil {
		app.ErrorLog.Printf("Could not add system admin to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SEND CONFIRMATION EMAIL
	if err := app.SendRegistrationEmail(emailAddress.Address, requestObject.AccountType, requestObject.Credentials); err != nil {
		app.ErrorLog.Printf("Could not add send confirmation email")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) SendRegistrationEmail(receiver string, accountType string, credentials models.Credentials) error {
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

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
	return err
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

	// SEND CONFIRMATION EMAIL
	if err := app.SendRegistrationEmail(emailAddress.Address, "User", user.Credentials); err != nil {
		app.ErrorLog.Printf("Could not add send confirmation email")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
