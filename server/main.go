package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gibbonhug/sprout/data"
)

// Func creates a server on port 3000 with chi,
// serving the dummy data defined in dummy.go on various GET endpoints.
// This is done because database for this application does not exist yet,
// to be able to have basic functionality while developing
func main() {
	/*
		CONNECT TO DATABASE (connection will be used in /handlers)
	*/

	_, err := data.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	/*
		ENDPOINTS
	*/

	// / GET endpoint: test if "hello world" displays on screen
	r.Get("/", GetHome)

	// /boxes GET endpoint: Return JSON of all Boxes
	r.Get("/boxes", GetBoxes)

	// /boxes/{id} GET endpoint: Return box with this id
	// TEMPORARY CODE
	r.Get("/boxes/{id}", GetBoxesParam)

	// /flowers GET endpoint: Return JSON of all flowers
	r.Get("/flowers", GetFlowers)

	// /flowers/{id} GET endpoint: Return flower with this id
	// TEMPORARY CODE
	r.Get("/flowers/{id}", GetFlowersParam)

	// /flowers POST endpoint: Create a flower and update JSON file
	r.Post("/flowers", PostFlowers)

	// /pairrlns GET endpoint: Return JSON of all pair relationships
	r.Get("/pairrlns", GetPairRlns)

	// /pairrlns/{id} GET endpoint: Return pair rln with this id
	// TEMPORARY CODE
	r.Get("/pairrlns/{id}", GetPairRlnsParam)
	
	// /clonerlns GET endpoint: Return JSON of all clone relationships
	r.Get("/clonerlns", GetCloneRlns)

	// /clonerlns/{id} GET endpoint: Return clone rln with this id
	// TEMPORARY CODE
	r.Get("/clonerlns/{id}", GetCloneRlnsParam)

	/*
		END ENDPOINTS
	*/
	
	fmt.Println("Serving on port 3000")

	http.ListenAndServe(":3000", r)
}
