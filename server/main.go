package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// For local JSON GET requests
// Set content-type and access-control-allow-origin headers
func setLocalJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
}

// Func creates a server on port 3000 with chi,
// serving the dummy data defined in dummy.go on various GET endpoints.
// This is done because database for this application does not exist yet, to be able to have basic functionality
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// / GET endpoint: test if "hello" displays on screen
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	// /boxes GET endpoint: Return JSON of all Boxes
	r.Get("/boxes", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		boxes, err := ioutil.ReadFile("./data/boxes.json")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(boxes)
	})

	// /flowers GET endpoint: Return JSON of all flowers
	r.Get("/flowers", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		flowers, err := ioutil.ReadFile("./data/boxes.json")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(flowers)
	})

	// /pairrelationships GET endpoint: Return JSON of all breed relationships
	r.Get("/pairrelationships", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)
	
		prlns, err := ioutil.ReadFile("./data/pairrelationships.json")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(prlns)
	})

	fmt.Println("Serving on port 3000")

	http.ListenAndServe(":3000", r)
}
