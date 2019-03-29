package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"net/http"
	"net/mail"
	"net/smtp"
)

func (app *Application) CompleteRegistrationSysAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := app.Store.CompleteRegistrationSystemAdmin(vars["email"]); err != nil {
		app.ErrorLog.Printf("Could not complete registration")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) RegisterSystemAdmin(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	// DECODE REQUEST
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// CHECK VALIDITY
	if credentials.Email == "" || credentials.Password == "" {
		app.ErrorLog.Println("Request fields are empty")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	emailAdress, emailError := mail.ParseAddress(credentials.Email)
	if emailError != nil {
		app.ErrorLog.Println("Invalid email format")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// ADD NEW SYSTEM ADMIN TO DATABASE
	err = app.Store.RegisterSystemAdmin(credentials)
	if err != nil {
		app.ErrorLog.Printf("Could not add system admin to database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SEND CONFIRMATION EMAIL
	if err := app.SendRegistrationEmail(emailAdress.Address, credentials); err != nil {
		app.ErrorLog.Printf("Could not add send confirmation email")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) SendRegistrationEmail(receiver string, credentials models.Credentials) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := "From: " + app.EmailAddress + "\n" +
		"To: " + receiver + "\n" +
		"Subject: Registration\n" +
		mime +
		`<html><head></head>
		<body><h1>ISA/MRS TIM6</h1>
		Complete your registration <a href="http://localhost:8080/confirmRegistration/` + receiver + `/">here</a>
		Your login details are: <br>
		EMAIL: ` + credentials.Email + `<br>
		PASSWORD: ` + credentials.Password + `
		</body></html>`

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
	return err
}
