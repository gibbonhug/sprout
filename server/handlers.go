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

	box, err := data.GetAllBox()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	boxJson, _ := json.Marshal(box)

	w.Write(boxJson)
}

func GetBoxesParam(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	// Grab url param
	intid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := (int32)(intid)

	// Get JSON data
	box, err := data.GetBoxFromID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	boxJson, _ := json.Marshal(box)

	w.Write(boxJson)
	return
}

func GetFlowers(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	flower, err := data.GetAllFlower()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	flowerJson, _ := json.Marshal(flower)

	w.Write(flowerJson)
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
	flower, err := data.GetFlowerFromID(id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	flowerJson, _ := json.Marshal(flower)

	w.Write(flowerJson)
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

	pair, err := data.GetAllPair()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	pairJson, _ := json.Marshal(pair)

	w.Write(pairJson)
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
	pair, _ := data.GetPairFromID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	pairJson, _ := json.Marshal(pair)

	w.Write(pairJson)
	return
}

func GetCloneRlns(w http.ResponseWriter, r *http.Request) {
	setLocalJSONHeaders(w)

	clone, err := data.GetAllClone()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cloneJson, _ := json.Marshal(clone)

	w.Write(cloneJson)
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
	clone, err := data.GetCloneFromID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	cloneJson, _ := json.Marshal(clone)

	w.Write(cloneJson)
	return
}
