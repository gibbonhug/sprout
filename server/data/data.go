/* Package is a work in progress
*/
package data

import (
	"context"
	"os"
	"encoding/json"

	"github.com/jackc/pgx/v5"

	"github.com/gibbonhug/sprout/flower"
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
	var flowerSlice []flower.Flower

	for flowerRows.Next() {
		values, err := flowerRows.Values()

		if err != nil {
			return nil, err
		}

		id := uint(values[0].(int32))
		color := values[1].(string)

		thisFlower := flower.Flower{ID: id, ColorPetal: color}

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