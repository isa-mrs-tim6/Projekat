package postgre

import (
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"strings"
)

func (db *Store) OneWaySearch(query models.OneWayQuery) ([]models.Flight, error) {
	var flights []models.Flight
	if err := db.Preload("Origin").Preload("Destination").Preload("Airline").
		Joins("JOIN destinations fromDest on fromDest.id=flights.origin_id").
		Where("LOWER(fromDest.name) LIKE ?", "%"+query.From+"%").
		Joins("JOIN destinations toDest on toDest.id=flights.destination_id").
		Where("LOWER(toDest.name) LIKE ? and date(flights.departure)=?", "%"+query.To+"%", query.Date).
		Joins("JOIN airplanes on airplanes.id=flights.airplane_id").
		Joins("JOIN seats on flights.airplane_id=seats.airplane_id").
		Where("seats.class=? and seats.reservation_id = 0", query.SeatClass).
		Group("flights.id").
		Having("count(flights.id)>=?", query.Passengers).
		Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (db *Store) UserSearch(id uint, query models.UserQueryDto) ([]models.UserResultDTO, error) {
	var results []models.UserResultDTO
	if err := db.Table("users").
		Where("users.id != ? AND (Lower(users.email) LIKE Lower(?) OR Lower(users.name) LIKE Lower(?)) AND (status is null or status = 'PENDING')", id, string('%')+query.Query+string('%'), string('%')+query.Query+string('%')).
		Joins("LEFT JOIN friendships on (users.ID = friendships.user1_id) AND (friendships.user2_id = ?) OR (friendships.user1_id = ? AND friendships.user2_id = users.ID)", id, id).
		Select("users.id as id, users.name as name, users.surname as surname, users.email as email, friendships.status as status").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (db *Store) HotelSearch(query models.HotelQuery) ([]models.HotelSearchResults, error) {
	var hotels []models.Hotel
	var retval []models.Hotel

	query.Name = strings.ToLower(query.Name)
	query.Address = strings.ToLower(query.Address)

	if err := db.Preload("Rooms").Where("LOWER(hotels.name) LIKE ? AND LOWER(hotels.address) LIKE ?", "%"+query.Name+"%", "%"+query.Address+"%").Find(&hotels).Error; err != nil {
		return nil, err
	}

	var reservations []models.HotelReservation
	if err := db.Find(&reservations).Error; err != nil {
		return nil, err
	}

	for _, hotel := range hotels {
		for _, room := range hotel.Rooms {
			available := true
			for _, reservation := range reservations {
				for _, reservedRoom := range reservation.Rooms {
					if room.ID == reservedRoom.ID && ((reservation.Beginning.After(query.From) && reservation.Beginning.Before(query.To)) || (reservation.End.After(query.From) && reservation.End.Before(query.To))) {
						available = false
						break
					}
				}
				if available {
					break
				}
			}
			if available {
				retval = append(retval, hotel)
				break
			}
		}
	}

	var results []models.HotelSearchResults
	for _, hotel := range hotels {
		var rating []float64

		if err := db.Table("hotel_reservations").Where("hotel_id = ? AND hotel_rating != 0", hotel.ID).
			Pluck("avg(hotel_rating)", &rating).Error; err != nil {
			return nil, err
		}
		results = append(results, models.HotelSearchResults{Hotel: hotel, Rating: rating[0]})
	}

	return results, nil
}

func (db *Store) RoomSearch(query models.RoomQuery, userID uint, isQuickReserve bool) ([]models.RoomSearchResults, error) {
	var rooms []models.Room

	if err := db.Raw(`
			SELECT "rooms".* FROM "rooms" WHERE (("rooms"."deleted_at" IS NULL AND (rooms.hotel_id = ? AND rooms.capacity in (?))))
			EXCEPT (
				SELECT "rooms".* FROM "rooms"
			    JOIN room_reservations on room_reservations.room_id = rooms.id
			    JOIN hotel_reservations ON hotel_reservations.id = room_reservations.hotel_reservation_id
				WHERE "rooms"."deleted_at" IS NULL AND
					((hotel_reservations.beginning BETWEEN ? AND ?) OR
					 (hotel_reservations.end BETWEEN ? AND ?)) AND
						rooms.hotel_id = ? AND rooms.capacity in (?)
					GROUP BY rooms.id
			)
			EXCEPT (
				SELECT "rooms".* FROM "rooms"
			    JOIN room_quick_reserve_days ON room_quick_reserve_days.room_id = rooms.id
			);`, query.HotelID, query.Capacities,
		query.From, query.To, query.From, query.To, query.HotelID, query.Capacities).Scan(&rooms).Error; err != nil {
		return nil, err
	}

	var results []models.RoomSearchResults
	for _, room := range rooms {
		var rating []float64
		db.Table("room_ratings").Where("room_id = ? AND rating != 0", room.ID).
			Pluck("avg(rating)", &rating)

		price := room.Price
		if isQuickReserve {
			price *= 0.9
		}

		// Add discount based on number of reservations
		var numOfReservations uint
		var reward models.ReservationReward
		db.Table("reservations").Where("user_id = ?", userID).Count(&numOfReservations)
		db.First(&reward, "required_number < ?", numOfReservations)
		if reward.PriceScale != 0 {
			price *= reward.PriceScale
		}

		results = append(results, models.RoomSearchResults{Room: room, Rating: rating[0], Price: price})
	}

	return results, nil
}

func (db *Store) QuickReservationRoomSearch(query models.RoomQuery, userID uint, isQuickReserve bool) ([]models.RoomSearchResults, error) {
	var rooms []models.Room

	if err := db.Raw(`
			SELECT "rooms".* FROM "rooms"
			JOIN room_quick_reserve_days ON room_quick_reserve_days.room_id = rooms.id
			WHERE (("rooms"."deleted_at" IS NULL AND (rooms.hotel_id = ?)))
				EXCEPT (
					SELECT "rooms".* FROM "rooms"
											  JOIN room_reservations on room_reservations.room_id = rooms.id
											  JOIN hotel_reservations ON hotel_reservations.id = room_reservations.hotel_reservation_id
					WHERE "rooms"."deleted_at" IS NULL AND
						((hotel_reservations.beginning BETWEEN ? AND ?) OR
						 (hotel_reservations.end BETWEEN ? AND ?)) AND
							rooms.hotel_id = ?
						GROUP BY rooms.id
				)
				EXCEPT (
					SELECT "rooms".* FROM "rooms"
											  JOIN room_quick_reserve_days ON room_quick_reserve_days.room_id = rooms.id
											  LEFT JOIN room_reservations ON room_reservations.room_id = rooms.id
					WHERE "rooms"."deleted_at" IS NULL AND ((rooms.hotel_id = ?)) AND
						  ((? BETWEEN room_quick_reserve_days.start AND room_quick_reserve_days.end) OR
						(? BETWEEN room_quick_reserve_days.start AND room_quick_reserve_days.end))
					GROUP BY rooms.id
				);`, query.HotelID,
		query.From, query.To, query.From, query.To, query.HotelID,
		query.HotelID, query.From, query.To,
	).Scan(&rooms).Error; err != nil {
		return nil, err
	}
	var results []models.RoomSearchResults
	for _, room := range rooms {
		var rating []float64
		db.Table("room_ratings").Where("room_id = ? AND rating != 0", room.ID).
			Pluck("avg(rating)", &rating)

		price := room.Price
		if isQuickReserve {
			price *= 0.9
		}

		// Add discount based on number of reservations
		var numOfReservations uint
		var reward models.ReservationReward
		db.Table("reservations").Where("user_id = ?", userID).Count(&numOfReservations)
		db.First(&reward, "required_number < ?", numOfReservations)
		if reward.PriceScale != 0 {
			price *= reward.PriceScale
		}

		results = append(results, models.RoomSearchResults{Room: room, Rating: rating[0], Price: price})
	}

	return results, nil
}

func (db *Store) RacSearch(query models.RacQuery) ([]models.RentACarCompany, error) {
	var companies []models.RentACarCompany
	var retval []models.RentACarCompany

	query.Name = strings.ToLower(query.Name)
	query.Address = strings.ToLower(query.Address)

	// mora da se ucitaju rent a kar kompanije gledajuci tabele rent_a_car_companies i locations
	if err := db.Preload("Vehicles").Preload("Locations").Joins("FULl JOIN locations ON locations.rent_a_car_company_id = rent_a_car_companies.id").
		Where("LOWER(locations.address) like ?", "%"+strings.ToLower(query.Address)+"%").
		Where("LOWER(rent_a_car_companies.name) like ?", "%"+strings.ToLower(query.Name)+"%").
		Group("rent_a_car_companies.id").
		Find(&companies).Error; err != nil {
		return nil, err
	}

	var reservations []models.RentACarReservation
	if err := db.Find(&reservations).Error; err != nil {
		return nil, err
	}

	for _, rac := range companies {
		for _, v := range rac.Vehicles {
			available := true
			for _, reservation := range reservations {
				if v.ID == reservation.VehicleID &&
					((reservation.Beginning.After(query.From) && reservation.Beginning.Before(query.To)) ||
						(reservation.End.After(query.From) && reservation.End.Before(query.To))) {
					available = false
					break
				}
				if available {
					break
				}
			}
			if available {
				retval = append(retval, rac)
				break
			}
		}
	}

	return retval, nil
}
