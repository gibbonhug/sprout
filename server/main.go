package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gibbonhug/sprout/flower"

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
// This is done because database for this application does not exist yet, 
// to be able to have basic functionality while developing
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

	// /boxes/{id} GET endpoint: Return box with this id
	// TEMPORARY CODE
	r.Get("/boxes/{id}", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		// Grab url param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Get JSON data
		boxJSON, err := ioutil.ReadFile("./data/boxes.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Unmarshal JSON data to boxSlice
		var boxSlice []*flower.Box

		err = json.Unmarshal(boxJSON, &boxSlice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Loop through flowers to find one with this ID

		var returnData []byte

		for _, box := range boxSlice {
			if box.ID == uint(id) {
				returnData, err = json.Marshal(box)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				w.Write(returnData)
				return
			}
		}

		// Box with this ID does not exist
		http.Error(w, "Data does not exist", http.StatusNotFound)

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

	// /flowers/{id} GET endpoint: Return flower with this id
	// TEMPORARY CODE
	r.Get("/flowers/{id}", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		// Grab url param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Get JSON data
		flowersJSON, err := ioutil.ReadFile("./data/flowers.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Unmarshal JSON data to flowersSlice
		var flowersSlice []*flower.Flower

		err = json.Unmarshal(flowersJSON, &flowersSlice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Loop through flowers to find one with this ID

		var returnData []byte

		for _, flower := range flowersSlice {
			if flower.ID == uint(id) {
				returnData, err = json.Marshal(flower)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				w.Write(returnData)
				return
			}
		}

		// Flower with this ID does not exist
		http.Error(w, "Data does not exist", http.StatusNotFound)

	})

	// /pairrlns GET endpoint: Return JSON of all pair relationships
	r.Get("/pairrlns", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		prlns, err := ioutil.ReadFile("./data/pairrelationships.json")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(prlns)
	})

	// /pairrlns/{id} GET endpoint: Return pair rln with this id
	// TEMPORARY CODE
	r.Get("/pairrlns/{id}", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		// Grab url param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Get JSON data
		pairJSON, err := ioutil.ReadFile("./data/pairrelationships.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Unmarshal JSON data to pairSlice
		var pairSlice []*flower.PairRln

		err = json.Unmarshal(pairJSON, &pairSlice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Loop through flowers to find one with this ID

		var returnData []byte

		for _, pair := range pairSlice {
			if pair.ID == uint(id) {
				returnData, err = json.Marshal(pair)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				w.Write(returnData)
				return
			}
		}

		// Pair with this ID does not exist
		http.Error(w, "Data does not exist", http.StatusNotFound)

	})
	

	// /clonerlns GET endpoint: Return JSON of all clone relationships
	r.Get("/clonerlns", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		crlns, err := ioutil.ReadFile("./data/clonerelationships.json")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.Write(crlns)
	})

	// /clonerlns/{id} GET endpoint: Return flower with this id
	// TEMPORARY CODE
	r.Get("/clonerlns/{id}", func(w http.ResponseWriter, r *http.Request) {
		setLocalJSONHeaders(w)

		// Grab url param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Get JSON data
		cloneJSON, err := ioutil.ReadFile("./data/clonerelationships.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Unmarshal JSON data to cloneSlice
		var cloneSlice []*flower.CloneRln

		err = json.Unmarshal(cloneJSON, &cloneSlice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Loop through clones to find one with this ID

		var returnData []byte

		for _, clone := range cloneSlice {
			if clone.ID == uint(id) {
				returnData, err = json.Marshal(clone)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				w.Write(returnData)
				return
			}
		}

		// Clone with this ID does not exist
		http.Error(w, "Data does not exist", http.StatusNotFound)

	})


	fmt.Println("Serving on port 3000")

	http.ListenAndServe(":3000", r)
}
