package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gibbonhug/sprout/data"

	"github.com/go-chi/chi/v5"
)

// For local JSON GET requests
// Set content-type and access-control-allow-origin headers
func setLocalJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func GetBoxes(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	boxes, err := ioutil.ReadFile("./data/boxes.json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(boxes)
}

func GetBoxesParam(w http.ResponseWriter, r *http.Request) {
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
	var boxSlice []*data.Box

	err = json.Unmarshal(boxJSON, &boxSlice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Loop through flowers to find one with this ID

	var returnData []byte

	for _, box := range boxSlice {
		if box.ID == int32(id) {
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
}

func GetFlowers(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	flowers, err := data.GetAllFlowerJson()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(flowers)
}

func GetFlowersParam(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	// Grab url param
	intid, err := strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := (int32)(intid)

	// Get JSON data
	flowerJSON, err := data.GetFlowerFromIDJson(id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Write(flowerJSON)
	return
}

func PostFlowers(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data sent in this request, as a Flower
	decoder := json.NewDecoder(r.Body)
	var f data.Flower // The data as a flower
	err := decoder.Decode(&f)
	if err != nil {
		http.Error(w, "something bad happen", http.StatusBadRequest)
	}

	// Get old JSON data
	oldFlowersJSON, err := ioutil.ReadFile("./data/flowers.json")
	if err != nil {
		http.Error(w, "something bad happen", http.StatusBadRequest)
	}

	// Unmarshall old JSON data into 'flowers'
	var flowers []data.Flower
	err = json.Unmarshal(oldFlowersJSON, &flowers)
	if err != nil {
		http.Error(w, "something bad happen", http.StatusBadRequest)
	}

	// Append data in f (posted data) into flowers
	flowers = append(flowers, f)

	// convert flowers to json
	flowersJSON, _ := json.Marshal(flowers)

	// Update temp JSON database
	err = os.WriteFile("./data/flowers.json", flowersJSON, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPairRlns(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	pairs, err := data.GetAllPairJson()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(pairs)
}

func GetPairRlnsParam(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	// Grab url param
	intid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := (int32)(intid)

	// Get JSON data
	pairJSON, err := data.GetPairFromIDJson(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(pairJSON)
	return
}

func GetCloneRlns(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	clones, err := data.GetAllCloneJson()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(clones)
}

func GetCloneRlnsParam(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	// Grab url param
	intid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := (int32)(intid)

	// Get JSON data
	cloneJSON, err := data.GetCloneFromIDJson(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Write(cloneJSON)
	return
}
