package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

func (app *Application) GetRewards(w http.ResponseWriter, r *http.Request) {
	rewards, err := app.Store.GetReservationRewards()
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation rewards")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rewards)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation rewards into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateRewards(w http.ResponseWriter, r *http.Request) {
	var rewards []models.ReservationReward
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &rewards)
	} else {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.Store.UpdateReservationRewards(rewards)
}

func (app *Application) GetReservation(w http.ResponseWriter, r *http.Request) {
	// GET RESERVATION ID
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get flight ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// FIND RESERVATION AND CHECK IF IT BELONGS TO USER REQUESTING IT
	reservation, err := app.Store.GetReservation(uint(reservationID))
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if reservation.UserID != user.ID {
		app.ErrorLog.Printf("Cannot access this reservation")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(reservation)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CompleteQuickResVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.CompleteQuickResVehParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = app.Store.CompleteQuickResVehicle(params)

	if err != nil {
		app.ErrorLog.Printf("Cannot complete quick reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CompleteQuickResFlight(w http.ResponseWriter, r *http.Request) {
	var reservationDTO models.QuickFlightReservationGDTOV2
	err := json.NewDecoder(r.Body).Decode(&reservationDTO)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flightRes, err := app.Store.GetFlightReservation(reservationDTO.ID)
	var reservation = models.Reservation{
		MasterRef: 0,
		Passenger: models.Passenger{
			user.ID,
			user.UserInfo,
		},
		ReservationFlight: flightRes,
		ReservationFlightID: flightRes.ID,
	}
	err = app.Store.CreateMaterQuickReservation(&reservation)

	if err != nil {
		app.ErrorLog.Printf("Cannot complete quick reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetCompanyQuickVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleQuickResParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := app.Store.GetCompanyQuickVehicle(params)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetQuickVehRes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get vehicle ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := app.Store.GetQuickVehRes(uint(vehicleID))

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetAirlineGraphData(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrive airline admin")
	}
	data, err := app.Store.GetReservationGraphData(user.ID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservation data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) ReserveVehicle(w http.ResponseWriter, r *http.Request) {
	var params models.VehicleReservationParams

	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
	}

	// GET RESERVATION ID
	vars := mux.Vars(r)
	reservationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get reservation ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := app.Store.ReserveVehicle(uint(reservationID), params, user.ID); err != nil {
		app.ErrorLog.Printf("Could not complete reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) GetUserReservations(w http.ResponseWriter, r *http.Request) {
	var reservations []models.ReservationDAO

	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
	}

	reservations, err = app.Store.GetUserReservations(user.ID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive reservations")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(reservations)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode reservations into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) ReserveFlight(w http.ResponseWriter, r *http.Request) {
	// GET FLIGHT ID
	vars := mux.Vars(r)
	flightID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get flight ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DECODE QUERY
	var query models.FlightReservationParams
	err = json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// SET MASTER USER ID
	query.Users[0].ID = user.ID

	// RESERVE
	reservation, err := app.Store.ReserveFlight(flightID, query)
	if err != nil {
		app.ErrorLog.Println("Could not complete reservation")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(reservation.ID); err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flight, _ := app.Store.GetFlight(reservation.ReservationFlight.FlightID)
	AirlineName := flight.Airline.Name
	price := fmt.Sprintf("%.2f", reservation.ReservationFlight.Price)
	name := reservation.Name + string(" ") +  reservation.Surname
	seatClass := reservation.ReservationFlight.Seat.Class
	seatNumber := fmt.Sprint(reservation.ReservationFlight.Seat.Number)
	from := flight.Origin.Name
	to := flight.Destination.Name
	departureDate := flight.Departure.Format("15:04 02.01.2006")
	go app.FlightReservationEmail(email, AirlineName, name, seatClass, seatNumber, departureDate, from, to, price)
	for i := 1; i < len(query.Users); i++ {
		if query.Users[i].Email != "" {
			name2 := query.Users[i].Name + string(" ") +query.Users[i].Surname
			go app.FlightInvitationEmail(query.Users[i].Email, AirlineName, name, name2, departureDate, from, to)
		}
	}
}

func (app *Application) ReserveHotel(w http.ResponseWriter, r *http.Request) {
	// GET MASTER REF ID, HOTEL ID
	vars := mux.Vars(r)
	masterID, err := strconv.ParseUint(vars["master_id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get master reservation ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hotelID, err := strconv.ParseUint(vars["hotel_id"], 10, 64)
	if err != nil {
		app.ErrorLog.Println("Could not get hotel ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIND LOGGED USER
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// DECODE QUERY
	var searchQuery models.HotelReservationParamsDTO
	err = json.NewDecoder(r.Body).Decode(&searchQuery)
	if err != nil {
		app.ErrorLog.Println("Could not decode JSON")
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

	// CLEANUP QUERY
	var query models.HotelReservationParams
	query.Rooms = searchQuery.Rooms
	query.From = dateFrom
	query.To = dateTo
	query.IsQuickReserve = searchQuery.IsQuickReserve
	for k := range searchQuery.Features {
		query.Features = append(query.Features, &searchQuery.Features[k])
	}

	// RESERVE
	reservationID, hotel, users, rooms, features, price, from, to, err :=
		app.Store.ReserveHotel(uint(masterID), uint(hotelID), user.ID, query)
	if err != nil {
		app.ErrorLog.Println("Could not complete reservation")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(reservationID); err != nil {
		app.ErrorLog.Printf("Cannot encode reservation data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go app.HotelReservationEmail(email, hotel, users, rooms, features, from, to, price)
}

func (app *Application) CancelFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelFlight(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) CancelHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelHotel(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) CancelVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resID, err := strconv.ParseUint(vars["id"], 10, 64)

	if err = app.Store.CancelVehicle(uint(resID)); err != nil {
		app.ErrorLog.Printf("Could not cancel reservation")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (app *Application) FlightReservationEmail(receiver string, airline string, user string, SeatClass string,
	SeatNumber string, dateFrom string,from string, to string, price string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := "From: " + app.EmailAddress + "\n" +
		"To: " + receiver + "\n" +
		"Subject: Flight reservation\n" +
		mime +
		`<html><head></head>
		<body><h1>ISA/MRS TIM6</h1>
		Your reservation is successful!
		Reservation details: <br>
		Name: ` + user + `<br>
		Airline: ` + airline + `<br>
		Seat Class: ` + SeatClass + `<br>
		Seat Number: ` + SeatNumber + `<br>
		From: ` + from + `<br>
		To: ` + to + `<br>
		Departure Date: ` + dateFrom + `<br>
		Price: ` + price + `<br>
		</body></html>`

	_ = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
}

func (app *Application) FlightInvitationEmail(receiver string, airline string, user string, user2,
	dateFrom string,from string, to string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := "From: " + app.EmailAddress + "\n" +
		"To: " + receiver + "\n" +
		"Subject: Flight reservation\n" +
		mime +
		`<html><head></head>
		<body><h1>ISA/MRS TIM6</h1>
		You have been invited to flight by ` + user + `
		<br>Flight details: <br>
		Airline: ` + airline + `<br>
		From: ` + from + `<br>
		To: ` + to + `<br>
		Departure Date: ` + dateFrom + `<br>
		You can accept invitation on next link and look for more details about it: <br> 
		<a href="http://localhost:8080/login?inv=true">here</a>
		</body></html>`

	_ = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
}

func (app *Application) HotelReservationEmail(receiver string, hotel string, user string, hotelRooms string,
	hotelFeatures string, dateFrom string, dateTo string, price string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := "From: " + app.EmailAddress + "\n" +
		"To: " + receiver + "\n" +
		"Subject: Hotel reservation\n" +
		mime +
		`<html><head></head>
		<body><h1>ISA/MRS TIM6</h1>
		Your reservation is successful!
		Reservation details: <br>
		Name: ` + user + `<br>
		Hotel: ` + hotel + `<br>
		Rooms: ` + hotelRooms + `<br>
		Features: ` + hotelFeatures + `<br>
		From: ` + dateFrom + `<br>
		To: ` + dateTo + `<br>
		Price: ` + price + `<br>
		</body></html>`

	_ = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", app.EmailAddress, app.EmailPassword, "smtp.gmail.com"),
		app.EmailAddress, []string{receiver}, []byte(message))
}

func (app *Application) GetPriceScale(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Println("Could not retrieve user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var ret models.ReservationScaleDAO

	if ret.Scale, ret.Count, err = app.Store.GetPriceScale(uint(user.ID)); err != nil {
		app.ErrorLog.Printf("Could not get scale")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		app.ErrorLog.Printf("Cannot encode scale data into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
