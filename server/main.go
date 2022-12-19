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
	ID         uint    `json:"id"`
	ParentID   []uint `json:"parentID"` // Flowers may have 0,1,2, or more parents
	ChildID    []uint `json:"childID"`
}

type Box struct {
	ID uint `json:"id"`
	FlowerID uint `json:"box_flower"`
}

// Dummy flowers to delete later
var flowers = []Flower {
	{ColorPetal: "000000", ID: 0}, {ColorPetal: "ff0000", ID: 1},
}

// Dummy boxes
var boxes = []Box {
	{ID: 0, Flower: &flowers[0]}, {ID: 1, Flower: &flowers[1]}, {ID: 2}, {ID: 3},
	{ID: 4}, {ID: 5}, {ID: 6}, {ID: 7},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	// Boxes GET endpoint: Return JSON of all Boxes
	r.Get("/boxes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")

		jsonBoxes, err := json.Marshal(boxes)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(jsonBoxes)
	})


	// flowers GET endpoint: Return JSON of all flowers
	r.Get("/flowers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")

		jsonFlowers, err := json.Marshal(flowers)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(jsonFlowers)
	})

	http.ListenAndServe(":3000", r)

	fmt.Println("Serving on port 3000")
}
