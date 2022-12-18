package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Flower struct {
	ColorPetal string `json:"colorPetal"` // Hex color
	ID int `json:"id"`
	ParentID []uint `json:"parentID"` // Flowers may have 0,1,2, or more parents
	ChildID []uint `json:"childID"`
}

// Dummy flowers to delete later
var flowers = []Flower {
	{ColorPetal: "000000", ID: 0}, {ColorPetal: "ff0000", ID: 1},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// flowers GET endpoint: Return JSON of all flowers
	r.Get("/flowers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jsonFlowers, err := json.Marshal(flowers)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(jsonFlowers)

	})

	http.ListenAndServe(":3000", r)

	fmt.Println("Serving on port 3000")
}