package model

import (
	"backend/config"
	"context"
	"fmt"
)

type MapPin struct {
	Pinuuid  string `json:"UUID"`
	Pinid    string `json:"id"`
	Pintitle string `json:"pintitle"`
	Pindesc  string `json:"pindesc"`
	Pincolor string `json:"pincolor"`
	Pinlat   string `json:"pinlat"`
	Pinlong  string `json:"pinlong"`
}

func GetUserMapPins(UUID string) ([]MapPin, error) {
	fmt.Println("QUERYING FOR MAP PINS - USER:", UUID)
	//var column string

	rows, err := config.DB.Query(context.Background(),
		"SELECT id, title, description, color, latitude, longitude FROM user_map_points WHERE user_id = $1", UUID)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	defer rows.Close()

	var pins []MapPin
	for rows.Next() {
		var pin MapPin
		err := rows.Scan(&pin.Pinid,
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

func NewUserPinDB(UUID, Pintitle, Pindesc, Pincolor, Pinlat, Pinlong string) (*MapPin, error) {
	var u MapPin
	row := config.DB.QueryRow(context.Background(),
		`INSERT INTO user_map_points (user_id, title, description, color, latitude, longitude)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING user_id, title, description, color, latitude, longitude`,
		UUID, Pintitle, Pindesc, Pincolor, Pinlat, Pinlong)
	err := row.Scan(&u.Pinuuid, &u.Pintitle, &u.Pindesc, &u.Pincolor, &u.Pinlat, &u.Pinlong)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &u, nil
}

func DeletedUserMapPin(UUID, Pinid string) (*MapPin, error) {
	var u MapPin
	row := config.DB.QueryRow(context.Background(),
		`DELETE FROM user_map_points WHERE id = $1 AND user_id = $2`,
		Pinid, UUID)
	// Potentially deletable
	err := row.Scan(&u.Pinuuid, &u.Pintitle, &u.Pindesc, &u.Pincolor, &u.Pinlat, &u.Pinlong)
	if err == nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &u, nil
}
