package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Flowers
type Flower struct {
	ID         uint    `json:"id"` // ID of the flower
	ColorPetal string `json:"colorPetal"` // Hex color
}

// Boxes flowers are stored in
type Box struct {
	ID uint `json:"id"` // ID of the box
	Flower *Flower `json:"box_flower"` // Pointer to the Flower it contains
}

// Cloning a flower producing a child
type CloneRelationship struct {
	ID uint `json:"ID"` // ID of this relationship
	ParentID uint `json:"parentID"` // ID of the parent flower
	ChildID uint `json:"childID"` // ID of the child flower
}

// Breeding 2 flowers producing a child
type BreedRelationship struct {
	ID uint `json:"ID"` // ID of this relationship
	Parent1ID uint `json:"parent1ID"` // ID of instigator parent flower
	Parent2ID uint `json:"parent2ID"` // ID of the receiver parent flower
	ChildID uint `json:"childID"` // ID of the child flower
}

// Dummy flowers to delete later
var flowers = []Flower {
	{ColorPetal: "000000", ID: 0}, {ColorPetal: "ff0000", ID: 1},
	{ColorPetal: "800000", ID: 2},
}

// Dummy boxes
var boxes = []Box {
	{ID: 0, Flower: &flowers[0]}, {ID: 1, Flower: &flowers[1]}, {ID: 2, Flower: &flowers[2]}, {ID: 3},
	{ID: 4}, {ID: 5}, {ID: 6}, {ID: 7},
}

// Dummy relationships
var relationships = []BreedRelationship {
	{Parent1ID: 0, Parent2ID: 1, ChildID: 2},
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

	// /relationships GET endpoint: Return JSON of all relationships
	r.Get("/relationships", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")

		jsonRlns, err := json.Marshal(relationships)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(jsonRlns)
	})

	fmt.Println("Serving on port 3000")

	http.ListenAndServe(":3000", r)

}
