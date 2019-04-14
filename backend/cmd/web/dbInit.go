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
		&models.HotelReservation{}, &models.FlightReservation{}, &models.Destination{}, "user_reservations", "vehicle_reservations")
	fmt.Printf("DATABASE: Finished dropping, time taken: %f seconds\n", time.Since(timeDroppingTables).Seconds())

	fmt.Println("DATABASE: Auto migrating schema")
	timeAutoMigration := time.Now()
	// Creates tables, missing columns and missing indexes
	db.AutoMigrate(
		&models.HotelAdmin{}, &models.Hotel{}, &models.Feature{}, &models.Room{},
		&models.RentACarAdmin{}, &models.RentACarCompany{}, &models.Location{}, &models.Vehicle{},
		&models.AirlineAdmin{}, &models.Airplane{}, &models.Layovers{}, &models.Airline{}, &models.Seat{}, &models.Flight{},
		&models.SystemAdmin{}, &models.Friendship{}, &models.User{}, &models.Reservation{}, &models.RentACarReservation{},
		&models.HotelReservation{}, &models.FlightReservation{}, &models.Destination{})
	fmt.Printf("DATABASE: Finished automigration, time taken: %f seconds\n", time.Since(timeAutoMigration).Seconds())
}

func addModels(db *gorm.DB) {
	fmt.Println("DATABASE: Adding models")
	timeStart := time.Now()

	// hash passwords
	pass1, err := bcrypt.GenerateFromPassword([]byte("S_ADMIN1"), bcrypt.DefaultCost)

	// CREATING SYSTEM ADMINS
	systemAdmin := models.SystemAdmin{
		Credentials:          models.Credentials{Email: "S_ADMIN1@email.com", Password: string(pass1)},
		UserInfo:             models.UserInfo{Name: "S_ADMIN1_IME", Surname: "S_ADMIN1_PREZIME"},
		RegistrationComplete: true,
	}
	db.Create(&systemAdmin)

	// hash passwords
	pass1, err = bcrypt.GenerateFromPassword([]byte("A_ADMIN1"), bcrypt.DefaultCost)
	pass2, err := bcrypt.GenerateFromPassword([]byte("A_ADMIN2"), bcrypt.DefaultCost)
	pass3, err := bcrypt.GenerateFromPassword([]byte("A_ADMIN3"), bcrypt.DefaultCost)

	// CREATING AIRLINE ADMINS
	airlineAdmin := models.AirlineAdmin{
		Credentials:          models.Credentials{Email: "A_ADMIN1@email.com", Password: string(pass1)},
		UserInfo:             models.UserInfo{Name: "A_ADMIN1_IME", Surname: "A_ADMIN1_PREZIME"},
		RegistrationComplete: true,
	}
	airlineAdmin2 := models.AirlineAdmin{
		Credentials:          models.Credentials{Email: "A_ADMIN2@email.com", Password: string(pass2)},
		UserInfo:             models.UserInfo{Name: "A_ADMIN2_IME", Surname: "A_ADMIN2_PREZIME"},
		RegistrationComplete: true,
	}
	airlineAdmin3 := models.AirlineAdmin{
		Credentials:          models.Credentials{Email: "A_ADMIN3@email.com", Password: string(pass3)},
		UserInfo:             models.UserInfo{Name: "A_ADMIN3_IME", Surname: "A_ADMIN3_PREZIME"},
		RegistrationComplete: false,
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
		Credentials:          models.Credentials{Email: "H_ADMIN1@email.com", Password: string(pass1)},
		UserInfo:             models.UserInfo{Name: "H_ADMIN1_IME", Surname: "H_ADMIN1_PREZIME"},
		RegistrationComplete: true,
	}
	hotelAdmin2 := models.HotelAdmin{
		Credentials:          models.Credentials{Email: "H_ADMIN2@email.com", Password: string(pass2)},
		UserInfo:             models.UserInfo{Name: "H_ADMIN2_IME", Surname: "H_ADMIN2_PREZIME"},
		RegistrationComplete: true,
	}
	hotelAdmin3 := models.HotelAdmin{
		Credentials:          models.Credentials{Email: "H_ADMIN3@email.com", Password: string(pass3)},
		UserInfo:             models.UserInfo{Name: "H_ADMIN3_IME", Surname: "H_ADMIN3_PREZIME"},
		RegistrationComplete: false,
	}
	db.Create(&hotelAdmin)
	db.Create(&hotelAdmin2)
	db.Create(&hotelAdmin3)

	pass1, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN1"), bcrypt.DefaultCost)
	pass2, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN2"), bcrypt.DefaultCost)
	pass3, err = bcrypt.GenerateFromPassword([]byte("R_ADMIN3"), bcrypt.DefaultCost)

	// CREATING RENT-A-CAR ADMINS
	rentACarAdmin := models.RentACarAdmin{
		Credentials:          models.Credentials{Email: "R_ADMIN1@email.com", Password: string(pass1)},
		UserInfo:             models.UserInfo{Name: "R_ADMIN1_IME", Surname: "R_ADMIN1_PREZIME"},
		RegistrationComplete: true,
	}
	rentACarAdmin2 := models.RentACarAdmin{
		Credentials:          models.Credentials{Email: "R_ADMIN2@email.com", Password: string(pass2)},
		UserInfo:             models.UserInfo{Name: "R_ADMIN2_IME", Surname: "R_ADMIN2_PREZIME"},
		RegistrationComplete: true,
	}
	rentACarAdmin3 := models.RentACarAdmin{
		Credentials:          models.Credentials{Email: "R_ADMIN3@email.com", Password: string(pass3)},
		UserInfo:             models.UserInfo{Name: "R_ADMIN3_IME", Surname: "R_ADMIN3_PREZIME"},
		RegistrationComplete: false,
	}
	db.Create(&rentACarAdmin)
	db.Create(&rentACarAdmin2)
	db.Create(&rentACarAdmin3)

	// CREATING AIRLINES
	airline := models.Airline{
		AirlineProfile: models.AirlineProfile{
			Name:  "A1_NAME",
			Promo: "A1_PROMO",
			Address: models.Address{
				Address:    "A1_ADDRESS 1",
				Coordinate: models.Coordinate{Latitude: 123, Longitude: 432}},
		},
		Admins: []*models.AirlineAdmin{
			&airlineAdmin,
			&airlineAdmin2,
		},
		Airplanes: []models.Airplane{
			{Name: "Airplane1",
				Seats: []models.Seat{
					{Number: 1, Class: "ECONOMIC", QuickReserve: false},
					{Number: 2, Class: "BUSINESS", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false}},
			},
			{Name: "Airplane2",
				Seats: []models.Seat{
					{Number: 11, Class: "ECONOMIC", QuickReserve: false},
					{Number: 22, Class: "BUSINESS", QuickReserve: false},
					{Number: 33, Class: "FIRST", QuickReserve: false}},
			},
		},
		Destinations: []models.Destination{
			{Name: "A1_DEST1", Coordinate: models.Coordinate{Latitude: -11.124, Longitude: 4.24}},
			{Name: "A1_DEST2", Coordinate: models.Coordinate{Latitude: 12.124, Longitude: 44.24}},
			{Name: "A1_DEST3", Coordinate: models.Coordinate{Latitude: 3.124, Longitude: 54.24}},
		},
		Flights: []models.Flight{
			{
				PriceECONOMY: 200, PriceBUSINESS: 300, PriceFIRSTCLASS: 500, QuickReservationPriceScale: 0.8,
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 2, 3, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 2, 3, 14, 15, 0, 0, time.Local),
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
				PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 4, 3, 10, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 4, 3, 12, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "F1_L1_ADDRESS 1",
						Coordinate: models.Coordinate{Latitude: 53, Longitude: 42.414}}},
					{Address: models.Address{
						Address:    "F2_L2_ADDRESS 2",
						Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
				},
			},
		},
	}
	DeepCopy(&airline.Airplanes[0], &airline.Flights[0].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[1].Airplane)

	airline.Flights[1].Airplane.Seats[0].QuickReserve = true
	airline.Flights[1].Airplane.Seats[2].QuickReserve = true
	airline.Flights[0].Airplane.Seats[1].QuickReserve = true

	airline.Flights[0].Origin = &airline.Destinations[0]
	airline.Flights[0].Destination = &airline.Destinations[1]
	airline.Flights[1].Origin = &airline.Destinations[1]
	airline.Flights[1].Destination = &airline.Destinations[2]

	airline2 := models.Airline{
		AirlineProfile: models.AirlineProfile{
			Name:  "A2_NAME",
			Promo: "A2_PROMO",
			Address: models.Address{
				Address:    "A2_ADDRESS 1",
				Coordinate: models.Coordinate{Latitude: 123, Longitude: 432}},
		},
		Admins: []*models.AirlineAdmin{
			&airlineAdmin3,
		},
		Destinations: []models.Destination{
			{Name: "A2_DEST1", Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 4.24}},
			{Name: "A2_DEST2", Coordinate: models.Coordinate{Latitude: 12.124, Longitude: 1.24}},
			{Name: "A2_DEST3", Coordinate: models.Coordinate{Latitude: 35.124, Longitude: 44.24}},
		},
		Airplanes: []models.Airplane{
			{Name: "Airplane3",
				Seats: []models.Seat{
					{Number: 1, Class: "ECONOMIC", QuickReserve: false},
					{Number: 2, Class: "BUSINESS", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false}},
			},
			{Name: "Airplane4",
				Seats: []models.Seat{
					{Number: 11, Class: "ECONOMIC", QuickReserve: false},
					{Number: 22, Class: "BUSINESS", QuickReserve: false},
					{Number: 33, Class: "FIRST", QuickReserve: false}},
			},
		},
		Flights: []models.Flight{
			{
				PriceECONOMY: 200, PriceBUSINESS: 300, PriceFIRSTCLASS: 500, QuickReservationPriceScale: 0.8,
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 2, 3, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 2, 3, 14, 15, 0, 0, time.Local),
				Layovers:  nil,
			},
			{
				PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
				Duration:  time.Hour * 2,
				Distance:  380,
				Departure: time.Date(2019, 4, 3, 10, 17, 0, 0, time.Local),
				Landing:   time.Date(2019, 4, 3, 12, 17, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "F1_L1_ADDRESS 1",
						Coordinate: models.Coordinate{Latitude: 53, Longitude: 42.414}}},
					{Address: models.Address{
						Address:    "F2_L2_ADDRESS 2",
						Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
				},
			},
		},
	}
	DeepCopy(&airline2.Airplanes[0], &airline2.Flights[0].Airplane)
	airline2.Flights[0].Airplane.Seats[0].QuickReserve = true
	airline2.Flights[0].Airplane.Seats[2].QuickReserve = true
	DeepCopy(&airline2.Airplanes[1], &airline2.Flights[1].Airplane)
	airline2.Flights[1].Airplane.Seats[1].QuickReserve = true

	airline2.Flights[0].Origin = &airline2.Destinations[0]
	airline2.Flights[0].Destination = &airline2.Destinations[1]
	airline2.Flights[1].Origin = &airline2.Destinations[1]
	airline2.Flights[1].Destination = &airline2.Destinations[2]

	db.Create(&airline)
	db.Create(&airline2)

	// CREATING HOTELS
	hotel := models.Hotel{
		HotelProfile: models.HotelProfile{
			Name:        "H1_NAME",
			Description: "H1_DESC",
			Address: models.Address{
				Address:    "H1_ADDRESS",
				Coordinate: models.Coordinate{Latitude: 33.123214, Longitude: -5.21352},
			},
		},
		Admins: []*models.HotelAdmin{
			&hotelAdmin,
			&hotelAdmin2,
		},
		Features: []models.Feature{
			{Description: "H1_FEATURE1", Price: 50},
			{Description: "H1_FEATURE2", Price: 75},
			{Description: "H1_FEATURE3", Price: 90},
		},
		Rooms: []models.Room{
			{Number: 101, Capacity: 2, Price: 250},
			{Number: 102, Capacity: 3, Price: 350},
			{Number: 103, Capacity: 4, Price: 450},
			{Number: 104, Capacity: 5, Price: 650},
			{Number: 105, Capacity: 2, Price: 250},
		},
	}
	hotel2 := models.Hotel{
		HotelProfile: models.HotelProfile{
			Name:        "H2_NAME",
			Description: "H2_DESC",
			Address: models.Address{
				Address:    "H2_ADDRESS",
				Coordinate: models.Coordinate{Latitude: -12.342, Longitude: 4.3242}},
		},
		Admins: []*models.HotelAdmin{
			&hotelAdmin3,
		},
		Features: []models.Feature{
			{Description: "H2_FEATURE1", Price: 40},
			{Description: "H2_FEATURE2", Price: 55},
			{Description: "H2_FEATURE3", Price: 60},
		},
		Rooms: []models.Room{
			{Number: 1, Capacity: 2, Price: 150},
			{Number: 2, Capacity: 3, Price: 250},
			{Number: 3, Capacity: 4, Price: 350},
			{Number: 4, Capacity: 5, Price: 450},
			{Number: 5, Capacity: 2, Price: 550},
		},
	}
	db.Create(&hotel)
	db.Create(&hotel2)

	// CREATING RENT-A-CAR COMPANIES
	rentACarCompany := models.RentACarCompany{
		RentACarCompanyProfile: models.RentACarCompanyProfile{
			Name:  "RAC1_NAME",
			Promo: "RAC1_PROMO",
			Address: models.Address{
				Address:    "RAC1_ADDRESS",
				Coordinate: models.Coordinate{Latitude: 41, Longitude: 32}}},
		Admins: []*models.RentACarAdmin{
			&rentACarAdmin,
			&rentACarAdmin2,
		},
		Locations: []models.Location{
			{Address: models.Address{
				Address:    "RAC1_L1_ADDRESS 1",
				Coordinate: models.Coordinate{Latitude: 123, Longitude: 432}}},
			{Address: models.Address{
				Address:    "RAC1_L2_ADDRESS 2",
				Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
		},
		Vehicles: []models.Vehicle{
			{Name: "RAC1_V1", Capacity: 4, Type: "A", PricePerDay: 45, Discount: false},
			{Name: "RAC1_V2", Capacity: 5, Type: "B", PricePerDay: 65, Discount: true},
			{Name: "RAC1_V3", Capacity: 6, Type: "A", PricePerDay: 75, Discount: false},
		},
	}
	rentACarCompany2 := models.RentACarCompany{
		RentACarCompanyProfile: models.RentACarCompanyProfile{
			Name:  "RAC2_NAME",
			Promo: "RAC2_PROMO",
			Address: models.Address{
				Address:    "RAC2_ADDRESS",
				Coordinate: models.Coordinate{Latitude: 41, Longitude: 32}}},
		Admins: []*models.RentACarAdmin{
			&rentACarAdmin3,
		},
		Locations: []models.Location{
			{Address: models.Address{
				Address:    "RAC2_L1_ADDRESS 1",
				Coordinate: models.Coordinate{Latitude: 123, Longitude: 432}}},
			{Address: models.Address{
				Address:    "RAC2_L2_ADDRESS 2",
				Coordinate: models.Coordinate{Latitude: -21.124, Longitude: 512.24}}},
		},
		Vehicles: []models.Vehicle{
			{Name: "RAC2_V1", Capacity: 4, Type: "B", PricePerDay: 45, Discount: false},
			{Name: "RAC2_V2", Capacity: 5, Type: "A", PricePerDay: 65, Discount: false},
			{Name: "RAC2_V3", Capacity: 6, Type: "A", PricePerDay: 75, Discount: true},
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
			UserInfo:    models.UserInfo{Name: "USER1_IME", Surname: "USER1_PREZIME"},
			Address:     "USER1_ADDRESS",
			Phone:       "USER1_PHONE",
		},
		RegistrationComplete: true,
	}
	user2 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER2@email.com", Password: string(pass2)},
			UserInfo:    models.UserInfo{Name: "USER2_IME", Surname: "USER2_PREZIME"},
			Address:     "USER2_ADDRESS",
			Phone:       "USER2_ADDRESS",
		},
		RegistrationComplete: true,
	}
	user3 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER3@email.com", Password: string(pass3)},
			UserInfo:    models.UserInfo{Name: "USER3_IME", Surname: "USER3_PREZIME"},
			Address:     "USER3_ADDRESS",
			Phone:       "USER3_ADDRESS",
		},
		RegistrationComplete: false,
	}
	user4 := models.User{
		Profile: models.Profile{
			Credentials: models.Credentials{Email: "USER4@email.com", Password: string(pass4)},
			UserInfo:    models.UserInfo{Name: "USER4_IME", Surname: "USER4_PREZIME"},
			Address:     "USER4_ADDRESS",
			Phone:       "USER4_ADDRESS",
		},
		RegistrationComplete: false,
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
	reservation := models.Reservation{
		Holders: []*models.User{
			&user,
		},
		ReservationFlight: models.FlightReservation{
			Price: 250,
			Seats: []models.Seat{
				airline.Flights[0].Airplane.Seats[0],
			},
			FlightID: airline.Flights[0].ID,
		},
		ReservationHotel: models.HotelReservation{
			Price: 100,
			Rooms: []models.Room{
				hotel.Rooms[0],
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 7, 0, 0, 0, 0, time.Local),
			},
			HotelID: hotel.ID,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:    70,
			Location: rentACarCompany.Locations[0].Address.Address,
			Vehicles: []*models.Vehicle{
				&rentACarCompany.Vehicles[0],
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
			},
		},
	}
	reservation2 := models.Reservation{
		Holders: []*models.User{
			&user2,
			&user3,
		},
		ReservationFlight: models.FlightReservation{
			Price: 350,
			Seats: []models.Seat{
				airline.Flights[1].Airplane.Seats[1],
				airline.Flights[1].Airplane.Seats[2],
			},
			FlightID: airline.Flights[1].ID,
		},
		ReservationHotel: models.HotelReservation{
			Price: 180,
			Rooms: []models.Room{
				hotel2.Rooms[1],
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 5, 6, 0, 0, 0, 0, time.Local),
			},
		},
	}
	reservation3 := models.Reservation{
		Holders: []*models.User{
			&user4,
		},
		ReservationFlight: models.FlightReservation{
			Price: 250,
			Seats: []models.Seat{
				airline2.Flights[0].Airplane.Seats[1],
			},
			FlightID: airline2.Flights[1].ID,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:    70,
			Location: rentACarCompany2.Locations[0].Address.Address,
			Vehicles: []*models.Vehicle{
				&rentACarCompany2.Vehicles[0],
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
			},
		},
	}

	reservation4 := models.Reservation{
		Holders: []*models.User{
			&user4,
		},
		ReservationFlight: models.FlightReservation{
			Price: 250,
			Seats: []models.Seat{
				airline2.Flights[0].Airplane.Seats[1],
			},
			FlightID: airline2.Flights[1].ID,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:    70,
			Location: rentACarCompany2.Locations[0].Address.Address,
			Vehicles: []*models.Vehicle{
				&rentACarCompany2.Vehicles[0],
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
			},
		},
	}
	db.Create(&reservation)
	db.Create(&reservation2)
	db.Create(&reservation3)
	db.Create(&reservation4)

	fmt.Printf("DATABASE: Finished adding models, time taken: %f seconds\n", time.Since(timeStart).Seconds())
}
