/* Package holds the 'global' database pgx connection variable, functions to query database, as well as models for data (flowers, boxes etc)
 */
package data

import (
	"context"
	"os"

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
