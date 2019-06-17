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
		&models.VehicleRating{}, &models.FeatureAirline{}, &models.HotelReservationReward{}, &models.RoomQuickReserveDays{},models.FlightQuickReserveSeats{},
		"user_reservations", "room_reservations", "flight_reservation_feature", "hotel_reservation_feature", "vehicle_reservations", "reward_features")
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
		&models.VehicleRating{}, &models.FeatureAirline{}, &models.HotelReservationReward{}, &models.RoomQuickReserveDays{},models.SeatQuickReservation{},models.FlightQuickReserveSeats{})
	fmt.Printf("DATABASE: Finished automigration, time taken: %f seconds\n", time.Since(timeAutoMigration).Seconds())
}

func addModels(db *gorm.DB) {
	fmt.Println("DATABASE: Adding models")
	timeStart := time.Now()

	// hash passwords
	pass1, err := bcrypt.GenerateFromPassword([]byte("S_ADMIN1"), bcrypt.DefaultCost)

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

	// hash passwords
	pass1, err = bcrypt.GenerateFromPassword([]byte("A_ADMIN1"), bcrypt.DefaultCost)
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
				RowWidth: 6,
				Seats: []models.Seat{
					{Number: 1, Class: "ECONOMIC", QuickReserve: false},
					{Number: 2, Class: "BUSINESS", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false},
					{Number: 4, Class: "ECONOMIC", QuickReserve: false},
					{Number: 5, Class: "BUSINESS", QuickReserve: false},
					{Number: 6, Class: "FIRST", QuickReserve: false},
					{Number: 7, Class: "ECONOMIC", QuickReserve: false},
					{Number: 8, Class: "BUSINESS", QuickReserve: false},
					{Number: 9, Class: "FIRST", QuickReserve: false},
					{Number: 10, Class: "ECONOMIC", QuickReserve: false},
					{Number: 11, Class: "BUSINESS", QuickReserve: false},
					{Number: 12, Class: "FIRST", QuickReserve: false},
					{Number: 13, Class: "ECONOMIC", QuickReserve: false},
					{Number: 14, Class: "BUSINESS", QuickReserve: false},
					{Number: 15, Class: "FIRST", QuickReserve: false},
					{Number: 16, Class: "ECONOMIC", QuickReserve: false, Disabled: true},
					{Number: 17, Class: "BUSINESS", QuickReserve: false},
					{Number: 18, Class: "FIRST", QuickReserve: false},
					{Number: 19, Class: "ECONOMIC", QuickReserve: false, Disabled: true},
					{Number: 20, Class: "BUSINESS", QuickReserve: false},
					{Number: 21, Class: "FIRST", QuickReserve: false},
					{Number: 21, Class: "FIRST", QuickReserve: false}},
			},
			{Name: "Airplane2",
				RowWidth: 2,
				Seats: []models.Seat{
					{Number: 1, Class: "ECONOMIC", QuickReserve: false},
					{Number: 2, Class: "BUSINESS", QuickReserve: false},
					{Number: 3, Class: "FIRST", QuickReserve: false},
					{Number: 4, Class: "ECONOMIC", QuickReserve: false},
					{Number: 5, Class: "BUSINESS", QuickReserve: false},
					{Number: 6, Class: "FIRST", QuickReserve: false},
					{Number: 7, Class: "ECONOMIC", QuickReserve: false},
					{Number: 8, Class: "BUSINESS", QuickReserve: false},
					{Number: 9, Class: "FIRST", QuickReserve: false},
					{Number: 10, Class: "ECONOMIC", QuickReserve: false},
					{Number: 11, Class: "BUSINESS", QuickReserve: false},
					{Number: 12, Class: "FIRST", QuickReserve: false},
					{Number: 13, Class: "ECONOMIC", QuickReserve: false},
					{Number: 14, Class: "BUSINESS", QuickReserve: false},
					{Number: 15, Class: "FIRST", QuickReserve: false},
					{Number: 16, Class: "ECONOMIC", QuickReserve: false},
					{Number: 17, Class: "BUSINESS", QuickReserve: false},
					{Number: 18, Class: "FIRST", QuickReserve: false}},
			},
		},
		Destinations: []models.Destination{
			{Name: "A1_DEST1", Coordinate: models.Coordinate{Latitude: -11.124, Longitude: 4.24}},
			{Name: "A1_DEST2", Coordinate: models.Coordinate{Latitude: 12.124, Longitude: 44.24}},
			{Name: "A1_DEST3", Coordinate: models.Coordinate{Latitude: 3.124, Longitude: 54.24}},
		},
		AirlineFeatures: []models.FeatureAirline{
			{Name: "A1_FEATURE1", Icon: "add_circle_outline", Description: "A1_FEATURE1_DESC", Price: 50},
			{Name: "A1_FEATURE2", Icon: "add_circle_outline", Description: "A1_FEATURE2_DESC", Price: 75},
			{Name: "A1_FEATURE3", Icon: "add_circle_outline", Description: "A1_FEATURE3_DESC", Price: 90},
		},
		Flights: []models.Flight{
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
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
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 300, PriceFIRSTCLASS: 500,
					SmallSuitcase: 50, BigSuitcase: 150,
				},
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
			{
				PriceList: models.PriceList{
					PriceECONOMY: 300, PriceBUSINESS: 600, PriceFIRSTCLASS: 900,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 5, 10, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 5, 11, 14, 15, 0, 0, time.Local),
				Layovers: []models.Layovers{
					{Address: models.Address{
						Address:    "F1_L1_ADDRESS 1",
						Coordinate: models.Coordinate{Latitude: 13, Longitude: 52.414}}},
					{Address: models.Address{
						Address:    "F1_L2_ADDRESS 2",
						Coordinate: models.Coordinate{Latitude: -51.124, Longitude: 5.24}}},
				},
			},
		},
	}
	DeepCopy(&airline.Airplanes[0], &airline.Flights[0].Airplane)
	DeepCopy(&airline.Airplanes[1], &airline.Flights[1].Airplane)
	DeepCopy(&airline.Airplanes[0], &airline.Flights[2].Airplane)

	airline.Flights[1].Airplane.Seats[0].QuickReserve = true
	airline.Flights[1].Airplane.Seats[2].QuickReserve = true
	airline.Flights[0].Airplane.Seats[1].QuickReserve = true

	airline.Flights[0].Origin = &airline.Destinations[0]
	airline.Flights[0].Destination = &airline.Destinations[1]
	airline.Flights[1].Origin = &airline.Destinations[1]
	airline.Flights[1].Destination = &airline.Destinations[2]
	airline.Flights[2].Origin = &airline.Destinations[1]
	airline.Flights[2].Destination = &airline.Destinations[0]

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
			{Name: "A2_FEATURE1", Icon: "add_circle_outline", Description: "A2_FEATURE1_DESC", Price: 50},
			{Name: "A2_FEATURE2", Icon: "add_circle_outline", Description: "A2_FEATURE2_DESC", Price: 75},
			{Name: "A2_FEATURE3", Icon: "add_circle_outline", Description: "A2_FEATURE3_DESC", Price: 90},
		},
		Flights: []models.Flight{
			{
				PriceList: models.PriceList{
					PriceECONOMY: 200, PriceBUSINESS: 300, PriceFIRSTCLASS: 500,
					SmallSuitcase: 120, BigSuitcase: 232, QuickReservationPriceScale: 0.8,
				},
				Duration:  time.Hour * 5,
				Distance:  780,
				Departure: time.Date(2019, 2, 3, 9, 15, 0, 0, time.Local),
				Landing:   time.Date(2019, 2, 3, 14, 15, 0, 0, time.Local),
				Layovers:  nil,
			},
			{
				PriceList: models.PriceList{
					PriceECONOMY: 100, PriceBUSINESS: 200, PriceFIRSTCLASS: 400,
					SmallSuitcase: 100, BigSuitcase: 300,
				},
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
				Address:    "A1_DEST1",
				Coordinate: models.Coordinate{Latitude: 33.123214, Longitude: -5.21352},
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
			Name:        "H2_NAME",
			Description: "H2_DESC",
			Address: models.Address{
				Address:    "H2_ADDRESS",
				Coordinate: models.Coordinate{Latitude: -12.342, Longitude: 4.3242}},
		},
		Features: []models.Feature{
			{Name: "H2_FEATURE1", Icon: "add_circle_outline", Description: "H2_FEATURE1_DESC", Price: 40},
			{Name: "H2_FEATURE2", Icon: "add_circle_outline", Description: "H2_FEATURE2_DESC", Price: 55},
			{Name: "H2_FEATURE3", Icon: "add_circle_outline", Description: "H2_FEATURE3_DESC", Price: 60},
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
			Name:  "RAC1_NAME",
			Promo: "RAC1_PROMO",
			Address: models.Address{
				Address:    "RAC1_ADDRESS",
				Coordinate: models.Coordinate{Latitude: 41, Longitude: 32}}},
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
		Passenger: models.Passenger{
			UserID:   user.ID,
			UserInfo: user.UserInfo,
		},
		ReservationFlight: models.FlightReservation{
			Price:         250,
			Seat:          &airline2.Flights[1].Airplane.Seats[0],
			Flight:        &airline2.Flights[1],
			FlightRating:  2,
			CompanyRating: 3,
			Features: []*models.FeatureAirline{
				&airline2.AirlineFeatures[0],
				&airline2.AirlineFeatures[1],
			},
			IsQuickReserve: false,
		},
		ReservationHotel: models.HotelReservation{
			Price: 100,
			Rooms: []models.Room{
				hotel2.Rooms[0],
				hotel2.Rooms[2],
				hotel2.Rooms[3],
			},
			Ratings: []models.RoomRating{
				{Rating: 3, RoomID: hotel2.Rooms[0].ID},
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
				End:       time.Date(2020, 4, 7, 0, 0, 0, 0, time.Local),
			},
			Features: []*models.Feature{
				&hotel2.Features[0],
				&hotel2.Features[2],
				&hotel2.Features[1],
			},
			Hotel:          hotel2,
			HotelRating:    4,
			IsQuickReserve: false,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:           70,
			Location:        rentACarCompany.Locations[0].Address.Address,
			RentACarCompany: rentACarCompany,
			Vehicle:         rentACarCompany.Vehicles[0],
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
			},
			CompanyRating:  4,
			VehicleRating:  2,
			IsQuickReserve: false,
		},
	}
	reservation2 := models.Reservation{
		Passenger: models.Passenger{
			UserID:   user2.ID,
			UserInfo: user2.UserInfo,
		},
		ReservationFlight: models.FlightReservation{
			Price:         250,
			Seat:          &airline.Flights[2].Airplane.Seats[1],
			Flight:        &airline.Flights[2],
			FlightRating:  3,
			CompanyRating: 3,
			Features: []*models.FeatureAirline{
				&airline2.AirlineFeatures[0],
				&airline2.AirlineFeatures[1],
			},
			IsQuickReserve: false,
		},
		ReservationHotel: models.HotelReservation{
			Price: 180,
			Rooms: []models.Room{
				hotel.Rooms[0],
				hotel.Rooms[1],
			},
			Ratings: []models.RoomRating{
				{Rating: 4, RoomID: hotel.Rooms[0].ID},
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 9, 6, 0, 0, 0, 0, time.Local),
			},
			Features: []*models.Feature{
				&hotel.Features[0],
				&hotel.Features[1],
			},
			Hotel:          hotel,
			HotelRating:    3,
			IsQuickReserve: false,
		},
	}
	reservation3 := models.Reservation{
		Passenger: models.Passenger{
			UserID:   user3.ID,
			UserInfo: user3.UserInfo,
		},
		MasterRef: reservation.ID,
		ReservationFlight: models.FlightReservation{
			Price:         250,
			Seat:          &airline.Flights[2].Airplane.Seats[4],
			Flight:        &airline.Flights[2],
			FlightRating:  5,
			CompanyRating: 3,
			Features: []*models.FeatureAirline{
				&airline.AirlineFeatures[0],
				&airline.AirlineFeatures[1],
			},
			IsQuickReserve: false,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:           70,
			Location:        rentACarCompany.Locations[0].Address.Address,
			RentACarCompany: rentACarCompany,
			Vehicle:         rentACarCompany.Vehicles[1],
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 2, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 3, 0, 0, 0, 0, time.Local),
			},
			CompanyRating:  3,
			VehicleRating:  4,
			IsQuickReserve: false,
		},
	}

	reservation4 := models.Reservation{
		Passenger: models.Passenger{
			UserID:   user4.ID,
			UserInfo: user4.UserInfo,
		},
		ReservationFlight: models.FlightReservation{
			Price:         350,
			Seat:          &airline.Flights[0].Airplane.Seats[4],
			Flight:        &airline.Flights[0],
			FlightRating:  2,
			CompanyRating: 5,
			Features: []*models.FeatureAirline{
				&airline.AirlineFeatures[0],
				&airline.AirlineFeatures[1],
			},
			IsQuickReserve: false,
		},
		ReservationRentACar: models.RentACarReservation{
			Price:         70,
			Location:      rentACarCompany2.Locations[0].Address.Address,
			Vehicle:       rentACarCompany2.Vehicles[0],
			CompanyRating: 3,
			VehicleRating: 4,
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 4, 5, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 4, 6, 0, 0, 0, 0, time.Local),
			},
			IsQuickReserve: false,
		},
		ReservationHotel: models.HotelReservation{
			Price: 180,
			Rooms: []models.Room{
				hotel.Rooms[0],
			},
			Ratings: []models.RoomRating{
				{Rating: 4, RoomID: hotel.Rooms[0].ID},
			},
			Occupation: models.Occupation{
				Beginning: time.Date(2019, 5, 3, 0, 0, 0, 0, time.Local),
				End:       time.Date(2019, 5, 6, 0, 0, 0, 0, time.Local),
			},
			HotelID:     hotel.ID,
			HotelRating: 4,
			Features: []*models.Feature{
				&hotel.Features[0],
				&hotel.Features[1],
			},
			IsQuickReserve: false,
		},
	}

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

	db.Create(&reservation)
	db.Create(&reservation2)
	db.Create(&reservation3)
	db.Create(&reservation4)
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

	db.Create(&reward)
	db.Create(&reward2)
	db.Create(&reward3)

	fmt.Printf("DATABASE: Finished adding models, time taken: %f seconds\n", time.Since(timeStart).Seconds())
}
