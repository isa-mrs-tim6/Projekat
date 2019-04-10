package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/db/gorm/postgre"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
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
	emailDomain := flag.String("email", "", "Email domain")
	emailPassword := flag.String("emailPassword", "", "Email password")
	flag.Parse()

	// SETUP LOGGING
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// CHECK IF INFO IS GIVEN
	if *emailDomain == "" || *emailPassword == "" {
		log.Fatal("Email domain and password not given")
	}

	// SETUP DATABASE
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, *username, *password, dbname)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Could not establish connection to the database")
	}
	defer Close(db)

	// CHECK FOR FIRST TIME SETUP
	if !(*dbPersist) {
		initTables(db)
		addModels(db)
	}

	// INJECT TO APPLICATION
	app := Application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Store:         &postgre.Store{db},
		EmailAddress:  *emailDomain,
		EmailPassword: *emailPassword,
	}

	// RUN SERVER
	app.RunServer()
}

func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter()

	// USER API
	// TODO
	router.HandleFunc("/api/user/login", app.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/getUsers", app.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}/getProfile", app.GetUserProfile).Methods("GET")
	router.HandleFunc("/api/user/{id}/updateProfile", app.UpdateUserProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/register", app.RegisterUser).Methods("POST", "OPTIONS")

	// ADMIN API
	router.HandleFunc("/api/admin/register", app.RegisterAdmin).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/admin/{type}/completeRegistration/q={email}", app.CompleteRegistration).Methods("GET")

	// RESERVATION API
	// TODO

	// AIRLINE API
	router.HandleFunc("/api/airline/getAirlines", app.GetAirlines).Methods("GET")
	router.HandleFunc("/api/airline/{id}/getProfile", app.GetAirlineProfiles).Methods("GET")
	router.HandleFunc("/api/airline/{id}/updateProfile", app.UpdateAirlineProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/airline/addAirline", app.CreateAirline).Methods("POST", "OPTIONS")

	//DESTINATION API
	router.HandleFunc("/api/destination/getDestinations", app.GetDestinations).Methods("GET")
	router.HandleFunc("/api/destination/{id}/getDestination", app.GetDestination).Methods("GET")
	router.HandleFunc("/api/destination/{id}/updateDestination", app.UpdateDestination).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/destination/add", app.CreateDestination).Methods("POST", "OPTIONS")

	// HOTEL API
	router.HandleFunc("/api/hotel/getHotels", app.GetHotels).Methods("GET")
	router.HandleFunc("/api/hotel/{id}/getProfile", app.GetHotelProfile).Methods("GET")
	router.HandleFunc("/api/hotel/{id}/updateProfile", app.UpdateHotelProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/addHotel", app.CreateHotel).Methods("POST", "OPTIONS")

	// RENT-A-CAR API
	router.HandleFunc("/api/rentACarCompany/getRentACarCompanies", app.GetRentACarCompanies).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/{id}/getProfile", app.GetRentACarCompanyProfile).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/{id}/updateProfile", app.UpdateRentACarCompanyProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/{id}/findVehicles", app.FindVehicles).Methods("POST", "OPTIONS")

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
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}))(app.Routes()),
	}
	app.InfoLog.Printf("Starting server on %s", address)
	err := server.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
