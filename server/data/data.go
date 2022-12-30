/* Package is a work in progress
Not sure how to split up /data and /flower packages
*/
package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// Global connection variable
var DB *pgx.Conn

// Sets DB global connection to new pgx connection with .env connection url
func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("SPROUT_DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	DB = conn

	return DB, nil
}

// Get all flowers from database and return them as json array
func GetAllFlowersAsJson(conn *pgx.Conn) ([]byte, error) {
	flowerRows, err := conn.Query(context.Background(), "SELECT * FROM flower")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var flowerSlice []Flower

	for flowerRows.Next() {
		values, err := flowerRows.Values()

		if err != nil {
			return nil, err
		}

		id := uint(values[0].(int32))
		color := values[1].(string)

		thisFlower := Flower{ID: id, ColorPetal: color}

		flowerSlice = append(flowerSlice, thisFlower)
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flowerSlice)

	if err != nil {
		return nil, err
	}

	// no error
	return flowerJson, nil
}

func GetFlowerFromIDAsJson(conn *pgx.Conn, flowerID uint) ([]byte, error) {
	paramAsStr := fmt.Sprint(flowerID)
	queryString := "SELECT * FROM flower WHERE flower_id = " + paramAsStr

	flowerRows, err := conn.Query(context.Background(), queryString)

	if err != nil {
		return nil, err
	}

	var flower Flower
	valid := false // To avoid returning 0-valued flower if the flower does not exist

	for flowerRows.Next() {
		valid = true // There is a flower with the param id so do not return error for invalid

		values, err := flowerRows.Values()

		if err != nil {
			return nil, err
		}

		id := uint(values[0].(int32))
		color := values[1].(string)

		flower.ID = id
		flower.ColorPetal = color
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flower)

	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("Flower not found")
	}

	// no error
	return flowerJson, nil	
}