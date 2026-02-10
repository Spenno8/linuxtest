package model

import (
	"backend/config"
	"context"
	"fmt"
)

// MapPin represents a user-created pin on the map.
type MapPin struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Pintitle string `json:"pintitle"`
	Pindesc  string `json:"pindesc"`
	Pincolor string `json:"pincolor"`
	Pinlat   string `json:"pinlat"`
	Pinlong  string `json:"pinlong"`
}

// GetUserMapPins fetches all map pins for a given user.
// Returns a slice of MapPin objects.
func GetUserMapPins(UserID string) ([]MapPin, error) {
	fmt.Println("QUERYING FOR MAP PINS - USER:", UserID)

	// Query the database for pins belonging to the given user
	rows, err := config.DB.Query(context.Background(),
		"SELECT id, user_id, title, description, color, latitude, longitude FROM user_map_points WHERE user_id = $1", UserID)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	defer rows.Close()

	var pins []MapPin
	for rows.Next() {
		var pin MapPin
		err := rows.Scan(&pin.ID,
			&pin.UserID,
			&pin.Pintitle,
			&pin.Pindesc,
			&pin.Pincolor,
			&pin.Pinlat,
			&pin.Pinlong)
		if err != nil {
			fmt.Println("DB SCAN ERROR:", err)
			return nil, err
		}

		pins = append(pins, pin)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("ROWS ERROR:", err)
		return nil, err
	}

	fmt.Println("User map pins:", pins)
	return pins, nil
}

// NewUserPinDB inserts a new map pin into the database for the given user.
// Returns the newly created MapPin with its ID populated.
func NewUserPinDB(UUID, Pintitle, Pindesc, Pincolor, Pinlat, Pinlong string) (*MapPin, error) {
	var u MapPin
	row := config.DB.QueryRow(context.Background(),
		`INSERT INTO user_map_points (user_id, title, description, color, latitude, longitude)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, user_id, title, description, color, latitude, longitude`,
		UUID, Pintitle, Pindesc, Pincolor, Pinlat, Pinlong)
	err := row.Scan(&u.ID, &u.UserID, &u.Pintitle, &u.Pindesc, &u.Pincolor, &u.Pinlat, &u.Pinlong)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &u, nil
}

// UpdateUserPinDB updates the details of an existing pin.
// Only the owner of the pin (UserID) can update it.
// Returns the updated MapPin.
func UpdateUserPinDB(
	PinID, UserID, Pintitle, Pindesc, Pincolor, Pinlat, Pinlong string,
) (*MapPin, error) {

	var u MapPin

	row := config.DB.QueryRow(context.Background(),
		`UPDATE user_map_points
		 SET title = $1,
		     description = $2,
		     color = $3,
		     latitude = $4,
		     longitude = $5,
		     updated_at = NOW()
		 WHERE id = $6 AND user_id = $7
		 RETURNING id, user_id, title, description, color, latitude, longitude`,
		Pintitle,
		Pindesc,
		Pincolor,
		Pinlat,
		Pinlong,
		PinID,
		UserID,
	)

	err := row.Scan(
		&u.ID,
		&u.UserID,
		&u.Pintitle,
		&u.Pindesc,
		&u.Pincolor,
		&u.Pinlat,
		&u.Pinlong,
	)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}

	return &u, nil
}

// DeletedUserMapPin removes a pin from the database.
// Only the owner of the pin (UserID) can delete it.
// Returns the deleted MapPin as confirmation.
func DeletedUserMapPin(UserID, PinID string) (*MapPin, error) {
	var u MapPin
	row := config.DB.QueryRow(context.Background(),
		`DELETE FROM user_map_points
		 WHERE id = $1 AND user_id = $2
		 RETURNING id, user_id, title, description, color, latitude, longitude`,
		PinID, UserID)

	err := row.Scan(
		&u.ID,
		&u.UserID,
		&u.Pintitle,
		&u.Pindesc,
		&u.Pincolor,
		&u.Pinlat,
		&u.Pinlong,
	)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}

	return &u, nil
}
