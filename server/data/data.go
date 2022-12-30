/* Package holds the 'global' database pgx connection variable, functions to query database, as well as models for data (flowers, boxes etc)
 */
package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Global connection variables
var DB *pgxpool.Pool
var CTX context.Context

// Sets DB global connection to new pgx connection with .env connection url
func Connect() (*pgxpool.Pool, error) {
	CTX = context.Background()

	conn, err := pgxpool.New(CTX, os.Getenv("SPROUT_DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	DB = conn

	return DB, nil
}

// Get all flowers from database and return them as json array
func GetAllFlowersAsJson() ([]byte, error) {
	flowerRows, err := DB.Query(CTX, "SELECT * FROM flower")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var flowerSlice []Flower

	rs := pgxscan.NewRowScanner(flowerRows)

	for flowerRows.Next() {
		var flower Flower

		err = rs.Scan(&flower)

		if err != nil {
			return nil, errors.New("Error scanning flower")
		}
		
		flowerSlice = append(flowerSlice, flower)
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flowerSlice)

	if err != nil {
		return nil, errors.New("Problem marshalling into JSON")
	}

	// no error
	return flowerJson, nil
}

func GetFlowerFromIDAsJson(flowerID int32) ([]byte, error) {
	paramAsStr := fmt.Sprint(flowerID)
	queryString := "SELECT * FROM flower WHERE flower_id = " + paramAsStr

	flowerRows, err := DB.Query(CTX, queryString)

	if err != nil {
		return nil, err
	}

	if !flowerRows.Next() {
		return nil, errors.New("Flower not found")
	}
	
	rs := pgxscan.NewRowScanner(flowerRows)

	var flower Flower

	err = rs.Scan(&flower)

	if err != nil {
		return nil, errors.New("Error scanning flower")
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flower)

	if err != nil {
		return nil, errors.New("Error marshaling flower")
	}

	// no error
	return flowerJson, nil	
}