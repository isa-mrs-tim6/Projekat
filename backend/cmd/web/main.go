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
	router.HandleFunc("/api/user/login", app.LoginUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/getUsers", Validate(app.GetUser, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/user/getFriends", Validate(app.GetFriends, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/user/getProfile", Validate(app.GetUserProfile, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/user/updateProfile", Validate(app.UpdateUserProfile, []string{"User"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/register", app.RegisterUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/rate", app.Rate).Methods("POST", "OPTIONS")

	// ADMIN API
	router.HandleFunc("/api/admin/checkFirstPass", Validate(app.CheckFirstPass, []string{"SystemAdmin", "HotelAdmin", "AirlineAdmin", "Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/admin/login", app.LoginAdmin).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/admin/register", Validate(app.RegisterAdmin, []string{"SystemAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/admin/{type}/completeRegistration/q={email}", app.CompleteRegistration).Methods("GET")
	router.HandleFunc("/api/admin/getProfile", Validate(app.GetAdminProfile, []string{"SystemAdmin", "HotelAdmin", "AirlineAdmin", "Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/admin/updateProfile", Validate(app.UpdateAdminProfile, []string{"SystemAdmin", "HotelAdmin", "AirlineAdmin", "Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	//

	// RESERVATION API
	router.HandleFunc("/api/airline/getAirlineReservations", Validate(app.GetAirlineReservations, []string{"AirlineAdmin"})).Methods("GET")
	router.HandleFunc("/api/hotel/getHotelReservations", Validate(app.GetHotelReservations, []string{"HotelAdmin"})).Methods("GET")
	router.HandleFunc("/api/reservations/rewards", Validate(app.GetRewards, []string{"SystemAdmin", "User"})).Methods("GET")
	router.HandleFunc("/api/reservations/getReservations", Validate(app.GetUserReservations, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/reservations/rewards", Validate(app.UpdateRewards, []string{"SystemAdmin"})).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/reservations/{id}/reservation", Validate(app.GetReservation, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/reservations/airline/reserve/q={id}", Validate(app.ReserveFlight, []string{"User"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/reservations/hotel/{hotel_id}/reserve/masterRef={master_id}", Validate(app.ReserveHotel, []string{"User"})).Methods("POST", "OPTIONS")

	// FLIGHT API
	router.HandleFunc("/api/flight/add", app.CreateFlight).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/flight/getCompanyFlights", app.GetCompanyFlights).Methods("GET")
	router.HandleFunc("/api/flight/{id}/getFlight", app.GetFlight).Methods("GET")
	router.HandleFunc("/api/flight/updateSeats", app.UpdateSeats).Methods("POST", "OPTIONS")

	// AIRPLANE API
	router.HandleFunc("/api/airplane/getAirplanes", app.GetAirplanes).Methods("GET")
	router.HandleFunc("/api/airplane/getCompanyAirplanes", app.GetCompanyAirplanes).Methods("GET")

	// PRICELIST API
	router.HandleFunc("/api/priceList/update", app.UpdatePriceList).Methods("POST", "OPTIONS")

	// AIRLINE API
	router.HandleFunc("/api/airline/getAirlines", app.GetAirlines).Methods("GET")
	router.HandleFunc("/api/airline/getProfile", app.GetAirlineProfiles).Methods("GET")
	router.HandleFunc("/api/airline/updateProfile", app.UpdateAirlineProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/airline/addAirline", Validate(app.CreateAirline, []string{"SystemAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/airline/getFlightRatings", Validate(app.GetFlightRatings, []string{"AirlineAdmin"})).Methods("GET")
	router.HandleFunc("/api/airline/getGraphData", Validate(app.GetAirlineGraphData, []string{"AirlineAdmin"})).Methods("GET")

	// DESTINATION API
	router.HandleFunc("/api/destination/getDestinations", app.GetDestinations).Methods("GET")
	router.HandleFunc("/api/destination/getCompanyDestinations", app.GetCompanyDestinations).Methods("GET")
	router.HandleFunc("/api/destination/{id}/getDestination", app.GetDestination).Methods("GET")
	router.HandleFunc("/api/destination/{id}/updateDestination", app.UpdateDestination).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/destination/add", app.CreateDestination).Methods("POST", "OPTIONS")

	// HOTEL API
	router.HandleFunc("/api/hotel/getHotels", app.GetHotels).Methods("GET")
	router.HandleFunc("/api/hotel/getProfile", app.GetHotelProfile).Methods("GET")
	router.HandleFunc("/api/hotel/updateProfile", app.UpdateHotelProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/addHotel", Validate(app.CreateHotel, []string{"SystemAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/{id}/getRoomCapacities", Validate(app.GetRoomCapacities, []string{"User"})).Methods("GET")
	router.HandleFunc("/api/hotel/getRooms", Validate(app.GetRooms, []string{"User", "HotelAdmin"})).Methods("GET")
	router.HandleFunc("/api/hotel/getRoomRatings", Validate(app.GetRoomRatings, []string{"HotelAdmin"})).Methods("GET")
	router.HandleFunc("/api/hotel/addRooms", Validate(app.AddRooms, []string{"HotelAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/deleteRooms", Validate(app.DeleteRooms, []string{"HotelAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/updateRoom", Validate(app.UpdateRoom, []string{"HotelAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/features", Validate(app.GetHotelFeatures, []string{"HotelAdmin", "User"})).Methods("GET")
	router.HandleFunc("/api/hotel/features", Validate(app.AddHotelFeature, []string{"HotelAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/features", Validate(app.UpdateHotelFeature, []string{"HotelAdmin"})).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/hotel/features", Validate(app.DeleteHotelFeature, []string{"HotelAdmin"})).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/hotel/rewards", Validate(app.GetHotelRewards, []string{"HotelAdmin", "User"})).Methods("GET")
	router.HandleFunc("/api/hotel/rewards", Validate(app.AddHotelReward, []string{"HotelAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotel/rewards", Validate(app.UpdateHotelReward, []string{"HotelAdmin"})).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/hotel/rewards", Validate(app.DeleteHotelReward, []string{"HotelAdmin"})).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/hotel/updateQuickReserveDays", Validate(app.UpdateQuickReserveDays, []string{"HotelAdmin"})).Methods("PUT", "OPTIONS")

	// RENT-A-CAR API
	router.HandleFunc("/api/rentACarCompany/getCompanyVehicles", Validate(app.GetCompanyVehicles, []string{"Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/addVehicle", Validate(app.AddVehicle, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/updateVehicle", Validate(app.UpdateVehicle, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/deleteVehicle", Validate(app.DeleteVehicle, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/addLocation", Validate(app.AddLocation, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/deleteLocation", Validate(app.DeleteLocation, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/updateLocation", Validate(app.UpdateLocation, []string{"Rent-A-CarAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/getCompanyLocations", Validate(app.GetCompanyLocations, []string{"Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/getRentACarCompanies", app.GetRentACarCompanies).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/getProfile", Validate(app.GetRentACarCompanyProfile, []string{"Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/getVehicleRatings", Validate(app.GetVehicleRatings, []string{"Rent-A-CarAdmin"})).Methods("GET")
	router.HandleFunc("/api/rentACarCompany/updateProfile", app.UpdateRentACarCompanyProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/findVehicles", app.FindVehicles).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/{id}/reserveVehicle", Validate(app.ReserveVehicle, []string{"User"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/addRentACarCompany", Validate(app.CreateRentACarCompany, []string{"SystemAdmin"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rentACarCompany/getReservations", Validate(app.GetRACReservations, []string{"Rent-A-CarAdmin"})).Methods("GET")

	// SEARCH API
	router.HandleFunc("/api/search/oneWay", app.OneWaySearch).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/search/multi", app.MultiSearch).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/search/hotels", app.HotelSearch).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/search/rac", app.RacSearch).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/search/{id}/rooms", Validate(app.RoomSearch, []string{"User"})).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/search/{id}/quickRooms", Validate(app.QuickReserveRoomSearch, []string{"User"})).Methods("POST", "OPTIONS")

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
			handlers.AllowCredentials(),
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
