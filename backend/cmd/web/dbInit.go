package main

import (
	"encoding/json"
	"fmt"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func DeepCopy(a, b interface{}) {
	bytes, _ := json.Marshal(a)
	err := json.Unmarshal(bytes, b)
	if err != nil {
		panic(err)
	}
}

func initTables(db *gorm.DB) {
	fmt.Println("DATABASE: Dropping tables")
	timeDroppingTables := time.Now()
	// Drops all the tables if they exist for a clean initialization
	db.DropTableIfExists(
		&models.HotelAdmin{}, &models.Hotel{}, &models.Feature{}, &models.Room{},
		&models.RentACarAdmin{}, &models.RentACarCompany{}, &models.Location{}, &models.Vehicle{},
		&models.AirlineAdmin{}, &models.Airplane{}, &models.Layovers{}, &models.Airline{}, &models.Seat{}, &models.Flight{},
		&models.SystemAdmin{}, &models.Friendship{}, &models.User{}, &models.Reservation{}, &models.RentACarReservation{},
		&models.HotelReservation{}, &models.FlightReservation{}, &models.Destination{}, &models.RoomRating{}, &models.ReservationReward{},
		&models.VehicleRating{}, &models.FeatureAirline{}, &models.HotelReservationReward{}, &models.RoomQuickReserveDays{},
		"user_reservations", "room_reservations", "flight_reservation_feature", "room_reservation_feature", "vehicle_reservations", "reward_features")

	fmt.Printf("DATABASE: Finished dropping, time taken: %f seconds\n", time.Since(timeDroppingTables).Seconds())

	fmt.Println("DATABASE: Auto migrating schema")
	timeAutoMigration := time.Now()
	// Creates tables, missing columns and missing indexes
	db.AutoMigrate(
		&models.HotelAdmin{}, &models.Hotel{}, &models.Feature{}, &models.Room{},
		&models.RentACarAdmin{}, &models.RentACarCompany{}, &models.Location{}, &models.Vehicle{},
		&models.AirlineAdmin{}, &models.Airplane{}, &models.Layovers{}, &models.Airline{}, &models.Seat{}, &models.Flight{},
		&models.SystemAdmin{}, &models.Friendship{}, &models.User{}, &models.Reservation{}, &models.RentACarReservation{},
		&models.HotelReservation{}, &models.FlightReservation{}, &models.Destination{}, &models.RoomRating{}, &models.ReservationReward{},
		&models.VehicleRating{}, &models.FeatureAirline{}, &models.HotelReservationReward{}, &models.RoomQuickReserveDays{}, models.SeatQuickReservation{}, models.FlightQuickReserveSeats{})
	fmt.Printf("DATABASE: Finished automigration, time taken: %f seconds\n", time.Since(timeAutoMigration).Seconds())

	// HASH PASSWORD
	pass1, _ := bcrypt.GenerateFromPassword([]byte("S_ADMIN1"), bcrypt.DefaultCost)
	// CREATING SYSTEM ADMINS
	systemAdmin := models.SystemAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "S_ADMIN1@email.com", Password: string(pass1)},
			UserInfo:    models.UserInfo{Name: "S_ADMIN1_IME", Surname: "S_ADMIN1_PREZIME"},
			Address:     "S_ADMIN1_ADDRESS",
			Phone:       "S_ADMIN1_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     true,
	}
	db.Create(&systemAdmin)
	fmt.Printf("DATABASE: Finished creating system admin, time taken: %f seconds\n", time.Since(timeAutoMigration).Seconds())
}

func AddDemoModels(db *gorm.DB) {
	fmt.Println("DATABASE: Adding models")
	timeStart := time.Now()

	// hash passwords
	pass1, err := bcrypt.GenerateFromPassword([]byte("A_ADMIN1"), bcrypt.DefaultCost)
	pass2, err := bcrypt.GenerateFromPassword([]byte("A_ADMIN2"), bcrypt.DefaultCost)
	pass3, err := bcrypt.GenerateFromPassword([]byte("A_ADMIN3"), bcrypt.DefaultCost)

	// CREATING AIRLINE ADMINS
	airlineAdmin := models.AirlineAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "A_ADMIN1@email.com", Password: string(pass1)},
			UserInfo:    models.UserInfo{Name: "A_ADMIN1_IME", Surname: "A_ADMIN1_PREZIME"},
			Address:     "A_ADMIN1_ADDRESS",
			Phone:       "A_ADMIN1_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     true,
	}
	airlineAdmin2 := models.AirlineAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "A_ADMIN2@email.com", Password: string(pass2)},
			UserInfo:    models.UserInfo{Name: "A_ADMIN2_IME", Surname: "A_ADMIN2_PREZIME"},
			Address:     "A_ADMIN2_ADDRESS",
			Phone:       "A_ADMIN2_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     false,
	}
	airlineAdmin3 := models.AirlineAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "A_ADMIN3@email.com", Password: string(pass3)},
			UserInfo:    models.UserInfo{Name: "A_ADMIN3_IME", Surname: "A_ADMIN3_PREZIME"},
			Address:     "A_ADMIN3_ADDRESS",
			Phone:       "A_ADMIN3_PHONE",
		},
		RegistrationComplete: false,
		FirstPassChanged:     true,
	}
	db.Create(&airlineAdmin)
	db.Create(&airlineAdmin2)
	db.Create(&airlineAdmin3)

	// hash passwords
	pass1, err = bcrypt.GenerateFromPassword([]byte("H_ADMIN1"), bcrypt.DefaultCost)
	pass2, err = bcrypt.GenerateFromPassword([]byte("H_ADMIN2"), bcrypt.DefaultCost)
	pass3, err = bcrypt.GenerateFromPassword([]byte("H_ADMIN3"), bcrypt.DefaultCost)

	// CREATING HOTEL ADMINS
	hotelAdmin := models.HotelAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "H_ADMIN1@email.com", Password: string(pass1)},
			UserInfo:    models.UserInfo{Name: "H_ADMIN1_IME", Surname: "H_ADMIN1_PREZIME"},
			Address:     "H_ADMIN1_ADDRESS",
			Phone:       "H_ADMIN1_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     true,
		HotelID:              1,
	}
	hotelAdmin2 := models.HotelAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "H_ADMIN2@email.com", Password: string(pass2)},
			UserInfo:    models.UserInfo{Name: "H_ADMIN2_IME", Surname: "H_ADMIN2_PREZIME"},
			Address:     "H_ADMIN2_ADDRESS",
			Phone:       "H_ADMIN2_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     false,
		HotelID:              1,
	}
	hotelAdmin3 := models.HotelAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "H_ADMIN3@email.com", Password: string(pass3)},
			UserInfo:    models.UserInfo{Name: "H_ADMIN3_IME", Surname: "H_ADMIN3_PREZIME"},
			Address:     "H_ADMIN3_ADDRESS",
			Phone:       "H_ADMIN3_PHONE",
		},
		RegistrationComplete: false,
		FirstPassChanged:     true,
		HotelID:              2,
	}
	db.Create(&hotelAdmin)
	db.Create(&hotelAdmin2)
	db.Create(&hotelAdmin3)

	pass1, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN1"), bcrypt.DefaultCost)
	pass2, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN2"), bcrypt.DefaultCost)
	pass3, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN3"), bcrypt.DefaultCost)

	// CREATING RENT-A-CAR ADMINS
	rentACarAdmin := models.RentACarAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "R_ADMIN1@email.com", Password: string(pass1)},
			UserInfo:    models.UserInfo{Name: "R_ADMIN1_IME", Surname: "R_ADMIN1_PREZIME"},
			Address:     "R_ADMIN1_ADDRESS",
			Phone:       "R_ADMIN1_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     true,
		RentACarCompanyID:    1,
	}
	rentACarAdmin2 := models.RentACarAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "R_ADMIN2@email.com", Password: string(pass2)},
			UserInfo:    models.UserInfo{Name: "R_ADMIN2_IME", Surname: "R_ADMIN2_PREZIME"},
			Address:     "R_ADMIN2_ADDRESS",
			Phone:       "R_ADMIN2_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     false,
		RentACarCompanyID:    1,
	}
	rentACarAdmin3 := models.RentACarAdmin{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "R_ADMIN3@email.com", Password: string(pass3)},
			UserInfo:    models.UserInfo{Name: "R_ADMIN3_IME", Surname: "R_ADMIN3_PREZIME"},
			Address:     "R_ADMIN3_ADDRESS",
			Phone:       "R_ADMIN3_PHONE",
		},
		RegistrationComplete: true,
		FirstPassChanged:     true,
		RentACarCompanyID:    2,
	}
	db.Create(&rentACarAdmin)
	db.Create(&rentACarAdmin2)
	db.Create(&rentACarAdmin3)

	// CREATING AIRLINES
	airline := models.Airline{
		AirlineProfile: models.AirlineProfile{
			Name: "Air France",
			Promo: "Air France is a leading global player in its three main areas of activity: " +
				"passenger transport, cargo transport and aircraft maintenance. From its hubs at Paris-Orly and " +
				"Paris-Charles de Gaulle airports, the airline operates flights to 195 destinations and 91 countries. " +
				"Air France is a founding member of the SkyTeam global alliance, alongside Korean Air, Aeromexico and Delta. " +
				"With the North American airline, Air France has also set up a joint venture dedicated to the joint operation of " +
				"several hundred transatlantic flights every day.",
			Address: models.Address{
				Address:    "Paris",
				Coordinate: models.Coordinate{Latitude: 48.8566, Longitude: 2.3522}},
		},
		Admins: []*models.AirlineAdmin{
			&airlineAdmin,
			&airlineAdmin2,
		},
		Airplanes: []models.Airplane{
			{Name: "Airplane1",
				RowWidth: 6,
				Seats: []models.Seat{
					{Number: 1, Class: "FIRST", QuickReserve: false},
					{Number: 2, Class: "FIRST", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false},
					{Number: 4, Class: "FIRST", QuickReserve: false},
					{Number: 5, Class: "FIRST", QuickReserve: false},
					{Number: 6, Class: "FIRST", QuickReserve: false},
					{Number: 7, Class: "FIRST", QuickReserve: false},
					{Number: 8, Class: "FIRST", QuickReserve: false},
					{Number: 9, Class: "FIRST", QuickReserve: false},
					{Number: 10, Class: "BUSINESS", QuickReserve: false},
					{Number: 11, Class: "BUSINESS", QuickReserve: false},
					{Number: 12, Class: "BUSINESS", QuickReserve: false},
					{Number: 13, Class: "BUSINESS", QuickReserve: false},
					{Number: 14, Class: "BUSINESS", QuickReserve: false},
					{Number: 15, Class: "BUSINESS", QuickReserve: false},
					{Number: 16, Class: "ECONOMIC", QuickReserve: false, Disabled: true},
					{Number: 17, Class: "ECONOMIC", QuickReserve: false},
					{Number: 18, Class: "ECONOMIC", QuickReserve: false},
					{Number: 19, Class: "ECONOMIC", QuickReserve: false, Disabled: true},
					{Number: 20, Class: "ECONOMIC", QuickReserve: false},
					{Number: 21, Class: "ECONOMIC", QuickReserve: false},
					{Number: 21, Class: "ECONOMIC", QuickReserve: false}},
			},
			{Name: "Airplane2",
				RowWidth: 2,
				Seats: []models.Seat{
					{Number: 1, Class: "FIRST", QuickReserve: false},
					{Number: 2, Class: "FIRST", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false},
					{Number: 4, Class: "FIRST", QuickReserve: false},
					{Number: 5, Class: "FIRST", QuickReserve: false},
					{Number: 6, Class: "FIRST", QuickReserve: false},
					{Number: 7, Class: "BUSINESS", QuickReserve: false},
					{Number: 8, Class: "BUSINESS", QuickReserve: false},
					{Number: 9, Class: "BUSINESS", QuickReserve: false},
					{Number: 10, Class: "BUSINESS", QuickReserve: false},
					{Number: 11, Class: "BUSINESS", QuickReserve: false},
					{Number: 12, Class: "BUSINESS", QuickReserve: false},
					{Number: 13, Class: "BUSINESS", QuickReserve: false},
					{Number: 14, Class: "BUSINESS", QuickReserve: false},
					{Number: 15, Class: "ECONOMIC", QuickReserve: false},
					{Number: 16, Class: "ECONOMIC", QuickReserve: false},
					{Number: 17, Class: "ECONOMIC", QuickReserve: false},
					{Number: 18, Class: "ECONOMIC", QuickReserve: false}},
			},
		},
		Destinations: []models.Destination{
			{Name: "Berlin", Coordinate: models.Coordinate{Latitude: 52.5200, Longitude: 13.4050}},
			{Name: "Tokyo", Coordinate: models.Coordinate{Latitude: 35.6762, Longitude: 139.6503}},
			{Name: "Dubai", Coordinate: models.Coordinate{Latitude: 25.2048, Longitude: 55.2708}},
			{Name: "Belgrade", Coordinate: models.Coordinate{Latitude: 44.7866, Longitude: 20.4489}},
			{Name: "Las Vegas", Coordinate: models.Coordinate{Latitude: 36.1699, Longitude: -115.1398}},
			{Name: "Amsterdam", Coordinate: models.Coordinate{Latitude: 52.3680, Longitude: 4.9036}},
			{Name: "Moscow", Coordinate: models.Coordinate{Latitude: 44.7866, Longitude: 20.4489}},
			{Name: "London", Coordinate: models.Coordinate{Latitude: 51.5074, Longitude: 0.1278}},
			{Name: "Rome", Coordinate: models.Coordinate{Latitude: 41.9028, Longitude: 12.4964}},
			{Name: "Vienna", Coordinate: models.Coordinate{Latitude: 48.2082, Longitude: 16.3738}},
			{Name: "Madrid", Coordinate: models.Coordinate{Latitude: 40.4168, Longitude: 3.7038}},
		},
		AirlineFeatures: []models.FeatureAirline{
			{Name: "Coffee", Icon: "add_circle_outline", Description: "Coffee on flight", Price: 50},
			{Name: "Wifi", Icon: "add_circle_outline", Description: "Wifi on flight", Price: 75},
			{Name: "Newspaper", Icon: "add_circle_outline", Description: "Newspaper on flight", Price: 90},
		},
		Flights: []models.Flight{
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 8, 2, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 2, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "F1_L1_ADDRESS 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
					{Address: models.Address{
						Address:    "F1_L2_ADDRESS 2",
						Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 5.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 300, PriceFIRSTCLASS: 500,
					SmallSuitcase: 50, BigSuitcase: 150,
				},
				Duration:  time.Hour * 7,
				Distance:  380,
				Departure: time.Date(2019, 8, 2, 5, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 2, 12, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 53, Longitude: 42.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 4,
				Distance:  780,
				Departure: time.Date(2019, 8, 2, 10, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 2, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 16,
				Distance:  780,
				Departure: time.Date(2019, 8, 2, 22, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 8, 3, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 5.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 4,
				Distance:  780,
				Departure: time.Date(2019, 8, 3, 16, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 20, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 400, PriceFIRSTCLASS: 600,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 8, 4, 10, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 4, 15, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 5.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 4,
				Distance:  780,
				Departure: time.Date(2019, 8, 6, 10, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 6, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 5.24}}},
				},
			},
		},
	}
	DeepCopy(&airline.Airplanes[0], &airline.Flights[0].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[1].Airplane)
	DeepCopy(&airline.Airplanes[0], &airline.Flights[2].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[3].Airplane)
	DeepCopy(&airline.Airplanes[0], &airline.Flights[4].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[5].Airplane)
	DeepCopy(&airline.Airplanes[0], &airline.Flights[6].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[7].Airplane)


	airline.Flights[0].Origin = &airline.Destinations[0]
	airline.Flights[0].Destination = &airline.Destinations[4]
	airline.Flights[1].Origin = &airline.Destinations[1]
	airline.Flights[1].Destination = &airline.Destinations[4]
	airline.Flights[2].Origin = &airline.Destinations[1]
	airline.Flights[2].Destination = &airline.Destinations[4]

	airline.Flights[3].Origin = &airline.Destinations[2]
	airline.Flights[3].Destination = &airline.Destinations[4]
	airline.Flights[4].Origin = &airline.Destinations[3]
	airline.Flights[4].Destination = &airline.Destinations[4]
	airline.Flights[5].Origin = &airline.Destinations[4]
	airline.Flights[5].Destination = &airline.Destinations[0]
	airline.Flights[6].Origin = &airline.Destinations[5]
	airline.Flights[6].Destination = &airline.Destinations[4]
	airline.Flights[7].Origin = &airline.Destinations[5]
	airline.Flights[7].Destination = &airline.Destinations[4]



	airline2 := models.Airline{
		AirlineProfile: models.AirlineProfile{
			Name:  "Emirates Airline",
			Promo: "Established in Oct-1985, Emirates Airline is the national carrier of the emirate of Dubai, United Arab Emirates. " +
				"Emirates is the world's largest airline and is among the fastest-growing carriers worldwide, pursuing an aggressive n" +
				"etwork. The airline operates with a fleet of widebody equipment, and is the largest operator of the Airbus A380 aircraft " +
				"type. Emirates, operating from its hub at Dubai International Airport, provides an extensive network of services within t" +
				"he Middle East as well as to Africa, Asia, the South Pacific, North America, Europe and South America. Emirates SkyCargo is the " +
				"air freight division of Emirates serving over 40 destinations.",
			Address: models.Address{
				Address:    "Dubai",
				Coordinate: models.Coordinate{Latitude: 25.2048, Longitude: 55.2708}},
		},
		Admins: []*models.AirlineAdmin{
			&airlineAdmin3,
		},
		Destinations: []models.Destination{
			{Name: "Berlin", Coordinate: models.Coordinate{Latitude: 52.5200, Longitude: 13.4050}},
			{Name: "Tokyo", Coordinate: models.Coordinate{Latitude: 35.6762, Longitude: 139.6503}},
			{Name: "Dubai", Coordinate: models.Coordinate{Latitude: 25.2048, Longitude: 55.2708}},
			{Name: "Belgrade", Coordinate: models.Coordinate{Latitude: 44.7866, Longitude: 20.4489}},
			{Name: "Las Vegas", Coordinate: models.Coordinate{Latitude: 36.1699, Longitude: -115.1398}},
			{Name: "Amsterdam", Coordinate: models.Coordinate{Latitude: 52.3680, Longitude: 4.9036}},
			{Name: "Moscow", Coordinate: models.Coordinate{Latitude: 44.7866, Longitude: 20.4489}},
			{Name: "London", Coordinate: models.Coordinate{Latitude: 51.5074, Longitude: 0.1278}},
			{Name: "Rome", Coordinate: models.Coordinate{Latitude: 41.9028, Longitude: 12.4964}},
			{Name: "Vienna", Coordinate: models.Coordinate{Latitude: 48.2082, Longitude: 16.3738}},
			{Name: "Madrid", Coordinate: models.Coordinate{Latitude: 40.4168, Longitude: 3.7038}},
		},
		Airplanes: []models.Airplane{
			{Name: "Airplane3",
				RowWidth: 6,
				Seats: []models.Seat{
					{Number: 1, Class: "ECONOMIC", QuickReserve: false},
					{Number: 2, Class: "BUSINESS", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false}},
			},
			{Name: "Airplane4",
				RowWidth: 2,
				Seats: []models.Seat{
					{Number: 11, Class: "ECONOMIC", QuickReserve: false},
					{Number: 22, Class: "BUSINESS", QuickReserve: false},
					{Number: 33, Class: "FIRST", QuickReserve: false}},
			},
		},
		AirlineFeatures: []models.FeatureAirline{
			{Name: "Wifi", Icon: "add_circle_outline", Description: "Wifi on flight", Price: 100},
			{Name: "Movies", Icon: "add_circle_outline", Description: "Movies on flight", Price: 200},
			{Name: "Shower", Icon: "add_circle_outline", Description: "Shower on flight", Price: 300},
		},
		Flights: []models.Flight{
			{
				PriceList: models.PriceList{
					PriceECONOMY: 200, PriceBUSINESS: 300, PriceFIRSTCLASS: 500,
					SmallSuitcase: 120, BigSuitcase: 232, QuickReservationPriceScale: 0.8,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 8, 3, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 14, 15, 0, 0, time.Local),
				Layovers:  nil,
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 8, 3, 10, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 12, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{

				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 8, 4, 12, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 4, 14, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 53, Longitude: 42.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 8, 4, 10, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 4, 12, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{

				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 8, 4, 15, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 4, 17, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "Layover 1",
						Coordinate: models.Coordinate{Latitude: 53, Longitude: 42.414}}},
					{Address: models.Address{
						Address:    "Layover 2",
						Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
				},
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 8, 3, 9, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 8, 3, 11, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{

				},
			},
		},
	}
	DeepCopy(&airline2.Airplanes[0], &airline2.Flights[0].Airplane)
	DeepCopy(&airline2.Airplanes[1], &airline2.Flights[1].Airplane)
	DeepCopy(&airline2.Airplanes[0], &airline2.Flights[2].Airplane)
	DeepCopy(&airline2.Airplanes[1], &airline2.Flights[3].Airplane)
	DeepCopy(&airline2.Airplanes[0], &airline2.Flights[4].Airplane)
	DeepCopy(&airline2.Airplanes[1], &airline2.Flights[5].Airplane)


	airline2.Flights[0].Origin = &airline2.Destinations[0]
	airline2.Flights[0].Destination = &airline2.Destinations[4]
	airline2.Flights[1].Origin = &airline2.Destinations[1]
	airline2.Flights[1].Destination = &airline2.Destinations[4]
	airline2.Flights[2].Origin = &airline2.Destinations[5]
	airline2.Flights[2].Destination = &airline2.Destinations[4]
	airline2.Flights[3].Origin = &airline2.Destinations[6]
	airline2.Flights[3].Destination = &airline2.Destinations[4]
	airline2.Flights[4].Origin = &airline2.Destinations[5]
	airline2.Flights[4].Destination = &airline2.Destinations[4]
	airline2.Flights[5].Origin = &airline2.Destinations[6]
	airline2.Flights[5].Destination = &airline2.Destinations[4]

	db.Create(&airline)
	db.Create(&airline2)

	// CREATING HOTELS
	hotel := models.Hotel{
		HotelProfile: models.HotelProfile{
			Name:        "The Palazzo",
			Description: "The center of a modern renaissance, every experience invites you to open your mind to what luxury can mean.",
			Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Latitude: 36.1699, Longitude: -115.1398},
			},
		},
		Features: []models.Feature{
			{Name: "H1_FEATURE1", Icon: "add_circle_outline", Description: "H1_FEATURE1_DESC", Price: 50},
			{Name: "H1_FEATURE2", Icon: "add_circle_outline", Description: "H1_FEATURE2_DESC", Price: 75},
			{Name: "H1_FEATURE3", Icon: "add_circle_outline", Description: "H1_FEATURE3_DESC", Price: 90},
		},
		Rooms: []models.Room{
			{Number: 1, Capacity: 2, Price: 250, QuickReserveDays: []models.RoomQuickReserveDays{
				{
					Start: time.Date(2019, 7, 3, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 7, 23, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 8, 3, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 8, 13, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 9, 3, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 9, 23, 0, 0, 0, 0, time.UTC),
				}},
			},
			{Number: 2, Capacity: 3, Price: 350, QuickReserveDays: []models.RoomQuickReserveDays{
				{
					Start: time.Date(2019, 7, 5, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 7, 15, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 8, 13, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 8, 23, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 9, 3, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 9, 7, 0, 0, 0, 0, time.UTC),
				}},
			},
			{Number: 3, Capacity: 4, Price: 450, QuickReserveDays: []models.RoomQuickReserveDays{
				{
					Start: time.Date(2019, 4, 2, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 7, 3, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 9, 4, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 9, 9, 0, 0, 0, 0, time.UTC),
				},
				{
					Start: time.Date(2019, 10, 3, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC),
				}},
			},
			{Number: 4, Capacity: 5, Price: 650},
			{Number: 5, Capacity: 2, Price: 250},
		},
	}

	hotel.Rewards = []models.HotelReservationReward{
		{Name: "H1_REWARD1", Description: "H1_DESC1", PriceScale: 0.95, Features: []*models.Feature{
			&hotel.Features[0], &hotel.Features[1]}},
		{Name: "H1_REWARD2", Description: "H1_DESC2", PriceScale: 0.85, Features: []*models.Feature{
			&hotel.Features[1], &hotel.Features[2]}},
	}

	hotel2 := models.Hotel{
		HotelProfile: models.HotelProfile{
			Name:        "The Venetian",
			Description: "A tribute to the Italian opulence, warmth and lust-for-life that informs each and every day.",
			Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Latitude: 36.1699, Longitude: -115.1398}},
		},
		Features: []models.Feature{
			{Name: "Wifi", Icon: "add_circle_outline", Description: "Wifi in room", Price: 40},
			{Name: "Room service", Icon: "add_circle_outline", Description: "Room service desc", Price: 55},
			{Name: "Bar", Icon: "add_circle_outline", Description: "Mini bar", Price: 60},
		},
		Rooms: []models.Room{
			{Number: 1, Capacity: 2, Price: 150},
			{Number: 2, Capacity: 3, Price: 250},
			{Number: 3, Capacity: 4, Price: 350},
			{Number: 4, Capacity: 5, Price: 450},
			{Number: 5, Capacity: 2, Price: 550},
		},
	}
	hotel2.Rewards = []models.HotelReservationReward{
		{Name: "H2_REWARD1", Description: "H2_DESC1", PriceScale: 0.90, Features: []*models.Feature{
			&hotel2.Features[0], &hotel2.Features[1]}},
		{Name: "H2_REWARD2", Description: "H2_DESC2", PriceScale: 0.85, Features: []*models.Feature{
			&hotel2.Features[1], &hotel2.Features[2]}},
	}
	db.Create(&hotel)
	db.Create(&hotel2)

	// CREATING RENT-A-CAR COMPANIES
	rentACarCompany := models.RentACarCompany{
		RentACarCompanyProfile: models.RentACarCompanyProfile{
			Name:  "LAS VEGAS CARS",
			Promo: "THE BEST VEHICLES IN LAS VEGAS",
			Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Longitude: 36.1699, Latitude: -115.1398}}},
		Locations: []models.Location{
			{Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Longitude: 36.1699, Latitude: -115.1398}}},
			{Address: models.Address{
				Address:    "New york",
				Coordinate: models.Coordinate{Longitude: 40.7128, Latitude: 74.0060}}},
		},
		Vehicles: []models.Vehicle{
			{Name: "VW GOLF 7", Capacity: 4, Type: "HATCHBACK", PricePerDay: 45, Discount: false},
			{Name: "AUDI R8", Capacity: 2, Type: "COUPE", PricePerDay: 150, Discount: true},
			{Name: "VW GOLF 8", Capacity: 6, Type: "HATCHBACK", PricePerDay: 75, Discount: false},
		},
	}
	rentACarCompany2 := models.RentACarCompany{
		RentACarCompanyProfile: models.RentACarCompanyProfile{
			Name:  "LAS VEGAS VEHICLES",
			Promo: "THE THE THE BEST VEHICLES IN LAS VEGAS",
			Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Longitude: 36.1699, Latitude: -115.1398}}},
		Locations: []models.Location{
			{Address: models.Address{
				Address:    "Las Vegas",
				Coordinate: models.Coordinate{Longitude: 36.1699, Latitude: -115.1398}}},
			{Address: models.Address{
				Address:    "New york",
				Coordinate: models.Coordinate{Longitude: 40.7128, Latitude: 74.0060}}},
		},
		Vehicles: []models.Vehicle{
			{Name: "Porsche Cayene", Capacity: 6, Type: "SUV", PricePerDay: 45, Discount: false},
			{Name: "VW TOUAREG", Capacity: 6, Type: "SUV", PricePerDay: 65, Discount: true},
			{Name: "VW GOLF 8", Capacity: 5, Type: "HATCHBACK", PricePerDay: 75, Discount: false},
		},
	}
	db.Create(&rentACarCompany)
	db.Create(&rentACarCompany2)

	// hash passwords
	pass1, err = bcrypt.GenerateFromPassword([]byte("USER1"), bcrypt.DefaultCost)
	pass2, err = bcrypt.GenerateFromPassword([]byte("USER2"), bcrypt.DefaultCost)
	pass3, err = bcrypt.GenerateFromPassword([]byte("USER3"), bcrypt.DefaultCost)
	pass4, err := bcrypt.GenerateFromPassword([]byte("USER4"), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	// CREATING USERS
	user := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER1@email.com", Password: string(pass1)},
			UserInfo:    models.UserInfo{Name: "Nikola", Surname: "Nikolic", Passport:    "1234567",},
			Address:     "USER1_ADDRESS",
			Phone:       "USER1_PHONE",
		},
		RegistrationComplete: true,
	}
	user2 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER2@email.com", Password: string(pass2)},
			UserInfo:    models.UserInfo{Name: "Marko", Surname: "Markovic",Passport:    "1234567",},
			Address:     "USER2_ADDRESS",
			Phone:       "USER2_ADDRESS",
		},
		RegistrationComplete: true,
	}
	user3 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER3@email.com", Password: string(pass3)},
			UserInfo:    models.UserInfo{Name: "Petar", Surname: "Petrovic",Passport:    "1234567",},
			Address:     "USER3_ADDRESS",
			Phone:       "USER3_ADDRESS",
		},
		RegistrationComplete: true,
	}
	user4 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER4@email.com", Password: string(pass4)},
			UserInfo:    models.UserInfo{Name: "Milos", Surname: "Milosevic",Passport:    "1234567",},
			Address:     "USER4_ADDRESS",
			Phone:       "USER4_ADDRESS",
		},
		RegistrationComplete: true,
	}
	db.Create(&user)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)

	// CREATING FRIENDSHIPS
	friendship := models.Friendship{
		Status: "ACCEPTED",
		User1:  &user,
		User2:  &user2,
	}
	friendship2 := models.Friendship{
		Status: "PENDING",
		User1:  &user,
		User2:  &user3,
	}
	db.Create(&friendship)
	db.Create(&friendship2)

	// CREATING RESERVATIONS
	//reservation := models.Reservation{
	//	Passenger: models.Passenger{
	//		UserID:   user.ID,
	//		UserInfo: user.UserInfo,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price:         250,
	//		Seat:          &airline2.Flights[1].Airplane.Seats[0],
	//		Flight:        &airline2.Flights[1],
	//		FlightRating:  2,
	//		CompanyRating: 3,
	//		Features: []*models.FeatureAirline{
	//			&airline2.AirlineFeatures[0],
	//			&airline2.AirlineFeatures[1],
	//		},
	//		IsQuickReserve: false,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 100,
	//		Rooms: []models.Room{
	//			hotel2.Rooms[0],
	//			hotel2.Rooms[2],
	//			hotel2.Rooms[3],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 3, RoomID: hotel2.Rooms[0].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2020, 4, 7, 0, 0, 0, 0, time.Local),
	//		},
	//		Features: []*models.Feature{
	//			&hotel2.Features[0],
	//			&hotel2.Features[2],
	//			&hotel2.Features[1],
	//		},
	//		Hotel:          hotel2,
	//		HotelRating:    4,
	//		IsQuickReserve: false,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:           70,
	//		Location:        rentACarCompany.Locations[0].Address.Address,
	//		RentACarCompany: rentACarCompany,
	//		Vehicle:         rentACarCompany.Vehicles[0],
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyRating:  4,
	//		VehicleRating:  2,
	//		IsQuickReserve: false,
	//	},
	//}
	//reservation2 := models.Reservation{
	//	Passenger: models.Passenger{
	//		UserID:   user2.ID,
	//		UserInfo: user2.UserInfo,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price:         250,
	//		Seat:          &airline.Flights[2].Airplane.Seats[1],
	//		Flight:        &airline.Flights[2],
	//		FlightRating:  3,
	//		CompanyRating: 3,
	//		Features: []*models.FeatureAirline{
	//			&airline2.AirlineFeatures[0],
	//			&airline2.AirlineFeatures[1],
	//		},
	//		IsQuickReserve: false,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 180,
	//		Rooms: []models.Room{
	//			hotel.Rooms[0],
	//			hotel.Rooms[1],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 4, RoomID: hotel.Rooms[0].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 9, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		Features: []*models.Feature{
	//			&hotel.Features[0],
	//			&hotel.Features[1],
	//		},
	//		Hotel:          hotel,
	//		HotelRating:    3,
	//		IsQuickReserve: false,
	//	},
	//}
	//reservation3 := models.Reservation{
	//	Passenger: models.Passenger{
	//		UserID:   user3.ID,
	//		UserInfo: user3.UserInfo,
	//	},
	//	MasterRef: reservation.ID,
	//	ReservationFlight: models.FlightReservation{
	//		Price:         250,
	//		Seat:          &airline.Flights[2].Airplane.Seats[4],
	//		Flight:        &airline.Flights[2],
	//		FlightRating:  5,
	//		CompanyRating: 3,
	//		Features: []*models.FeatureAirline{
	//			&airline.AirlineFeatures[0],
	//			&airline.AirlineFeatures[1],
	//		},
	//		IsQuickReserve: false,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:           70,
	//		Location:        rentACarCompany.Locations[0].Address.Address,
	//		RentACarCompany: rentACarCompany,
	//		Vehicle:         rentACarCompany.Vehicles[1],
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyRating:  3,
	//		VehicleRating:  4,
	//		IsQuickReserve: false,
	//	},
	//}
	//
	//reservation4 := models.Reservation{
	//	Passenger: models.Passenger{
	//		UserID:   user4.ID,
	//		UserInfo: user4.UserInfo,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price:         350,
	//		Seat:          &airline.Flights[0].Airplane.Seats[4],
	//		Flight:        &airline.Flights[0],
	//		FlightRating:  2,
	//		CompanyRating: 5,
	//		Features: []*models.FeatureAirline{
	//			&airline.AirlineFeatures[0],
	//			&airline.AirlineFeatures[1],
	//		},
	//		IsQuickReserve: false,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:         70,
	//		Location:      rentACarCompany2.Locations[0].Address.Address,
	//		Vehicle:       rentACarCompany2.Vehicles[0],
	//		CompanyRating: 3,
	//		VehicleRating: 4,
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		IsQuickReserve: false,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 180,
	//		Rooms: []models.Room{
	//			hotel.Rooms[0],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 4, RoomID: hotel.Rooms[0].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 5, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel.ID,
	//		HotelRating: 4,
	//		Features: []*models.Feature{
	//			&hotel.Features[0],
	//			&hotel.Features[1],
	//		},
	//		IsQuickReserve: false,
	//	},
	//}

	//reservation5 := models.Reservation{
	//	Holders: []*models.User{
	//		&user3, &user,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[0],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 2, VehicleID: rentACarCompany2.Vehicles[0].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 6, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 6, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    4,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 180,
	//		Rooms: []models.Room{
	//			hotel.Rooms[2],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 5, RoomID: hotel.Rooms[2].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 6, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 9, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel.ID,
	//		HotelRating: 4,
	//	},
	//}
	//
	//reservation6 := models.Reservation{
	//	Holders: []*models.User{
	//		&user4, &user2,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[1],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 5, VehicleID: rentACarCompany2.Vehicles[1].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    5,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 280,
	//		Rooms: []models.Room{
	//			hotel.Rooms[1],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 2, RoomID: hotel.Rooms[1].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 7, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 8, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel.ID,
	//		HotelRating: 4,
	//	},
	//}
	//
	//reservation7 := models.Reservation{
	//	Holders: []*models.User{
	//		&user,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[1],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 5, VehicleID: rentACarCompany2.Vehicles[1].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    5,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 380,
	//		Rooms: []models.Room{
	//			hotel.Rooms[0],
	//			hotel.Rooms[4],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 4, RoomID: hotel.Rooms[0].ID},
	//			{Rating: 1, RoomID: hotel.Rooms[4].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 9, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 9, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel.ID,
	//		HotelRating: 4,
	//	},
	//}
	//
	//reservation8 := models.Reservation{
	//	Holders: []*models.User{
	//		&user2,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[1],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 4, VehicleID: rentACarCompany2.Vehicles[1].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    5,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 480,
	//		Rooms: []models.Room{
	//			hotel2.Rooms[2],
	//			hotel2.Rooms[1],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 2, RoomID: hotel2.Rooms[2].ID},
	//			{Rating: 5, RoomID: hotel2.Rooms[1].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 9, 10, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel2.ID,
	//		HotelRating: 4,
	//	},
	//}
	//
	//reservation9 := models.Reservation{
	//	Holders: []*models.User{
	//		&user3,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[2],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 5, VehicleID: rentACarCompany2.Vehicles[2].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    5,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 580,
	//		Rooms: []models.Room{
	//			hotel2.Rooms[3],
	//			hotel2.Rooms[4],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 3, RoomID: hotel2.Rooms[3].ID},
	//			{Rating: 2, RoomID: hotel2.Rooms[4].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 10, 4, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 11, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel2.ID,
	//		HotelRating: 4,
	//	},
	//}
	//
	//reservation10 := models.Reservation{
	//	Holders: []*models.User{
	//		&user4,
	//	},
	//	ReservationFlight: models.FlightReservation{
	//		Price: 250,
	//		Seats: []models.Seat{
	//			airline2.Flights[0].Airplane.Seats[1],
	//		},
	//		FlightID: airline2.Flights[1].ID,
	//	},
	//	ReservationRentACar: models.RentACarReservation{
	//		Price:    70,
	//		Location: rentACarCompany2.Locations[0].Address.Address,
	//		Vehicles: []*models.Vehicle{
	//			&rentACarCompany2.Vehicles[2],
	//		},
	//		Ratings: []models.VehicleRating{
	//			{Rating: 1, VehicleID: rentACarCompany.Vehicles[2].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		CompanyID: rentACarCompany2.ID,
	//		Rating:    1,
	//	},
	//	ReservationHotel: models.HotelReservation{
	//		Price: 680,
	//		Rooms: []models.Room{
	//			hotel2.Rooms[2],
	//		},
	//		Ratings: []models.RoomRating{
	//			{Rating: 4, RoomID: hotel2.Rooms[2].ID},
	//		},
	//		Occupation: models.Occupation{
	//			Beginning: time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
	//			End:       time.Date(2021, 5, 6, 0, 0, 0, 0, time.Local),
	//		},
	//		HotelID:     hotel2.ID,
	//		HotelRating: 4,
	//	},
	//}

	//db.Create(&reservation)
	//db.Create(&reservation2)
	//db.Create(&reservation3)
	//db.Create(&reservation4)
	//db.Create(&reservation5)
	//db.Create(&reservation6)
	//db.Create(&reservation7)
	//db.Create(&reservation8)
	//db.Create(&reservation9)
	//db.Create(&reservation10)

	reward := models.ReservationReward{
		RequiredNumber: 5,
		PriceScale:     0.95,
	}

	reward2 := models.ReservationReward{
		RequiredNumber: 10,
		PriceScale:     0.90,
	}

	reward3 := models.ReservationReward{
		RequiredNumber: 20,
		PriceScale:     0.80,
	}

	reward4 := models.ReservationReward{
		RequiredNumber: 0,
		PriceScale:     1.0,
	}

	db.Create(&reward)
	db.Create(&reward2)
	db.Create(&reward3)
	db.Create(&reward4)


	fmt.Printf("DATABASE: Finished adding models, time taken: %f seconds\n", time.Since(timeStart).Seconds())
}
