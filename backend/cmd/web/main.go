package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/db/gorm/postgre"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

const (
	host    = "localhost"
	port    = 5432
	dbname  = "postgres"
	address = ":8000"
)

func main() {
	// SETUP DB CONNECTION DETAILS
	username := flag.String("username", "postgres", "Database username")
	password := flag.String("password", "admin", "Database password")
	dbPersist := flag.Bool("persist", true, "Recreate database tables and add mock-up objects")
	flag.Parse()

	// SETUP LOGGING
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// SETUP DATABASE
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, *username, *password, dbname)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Could not establish connection to the database")
		panic(err)
	}
	defer db.Close()

	// CHECK FOR FIRST TIME SETUP
	if !(*dbPersist) {
		initTables(db)
		addModels(db)
	}

	// INJECT TO APPLICATION
	app := Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Store:    &postgre.Store{db},
	}

	// RUN SERVER
	app.RunServer()
}

func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter()

	// USER API
	// TODO
	router.HandleFunc("/api/user/getUsers", app.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}/getProfile", app.GetUserProfile).Methods("GET")
	router.HandleFunc("/api/user/{id}/updateProfile", app.UpdateUserProfile).Methods("POST")

	// SYSTEM ADMIN API
	// TODO

	// AIRLINE ADMIN API
	// TODO
	router.HandleFunc("/api/airline/getAirline", app.GetAirline).Methods("GET")
	router.HandleFunc("/api/airline/{id}/getProfile", app.GetAirlineProfiles).Methods("GET")
	router.HandleFunc("/api/airline/{id}/updateProfile", app.UpdateAirlineProfile).Methods("POST")

	// HOTEL ADMIN API
	// TODO

	// RENT-A-CAR ADMIN API
	router.HandleFunc("/api/rentACarCompany/getRentACarCompanies", app.GetRentACarCompanies).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/{id}/getProfile", app.GetRentACarCompanyProfile).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/{id}/updateProfile", app.UpdateRentACarCompanyProfile).Methods("POST")
	router.HandleFunc("/api/rentACarCompany/{id}/findVehicles", app.FindVehicles).Methods("POST")
	// TODO

	// RESERVATION API
	// TODO

	// AIRLINE API
	// TODO

	// HOTEL API
	router.HandleFunc("/api/hotel/getHotels", app.GetHotels).Methods("GET")
	router.HandleFunc("/api/hotel/{id}/getProfile", app.GetHotelProfile).Methods("GET")
	router.HandleFunc("/api/hotel/{id}/updateProfile", app.UpdateHotelProfile).Methods("POST")
	router.HandleFunc("/api/hotel/addHotel", app.CreateHotel).Methods("POST", "OPTIONS")

	// RENT-A-CAR API
	// TODO

	// STATIC FILE HANDLER
	staticFileDirectory := http.Dir("./ui/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	return router
}

func (app *Application) RunServer() {
	server := &http.Server{
		Addr:     address,
		ErrorLog: app.ErrorLog,
		Handler:  app.Routes(),
	}
	app.InfoLog.Printf("Starting server on %s", address)
	err := server.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

func (app *Application) setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
