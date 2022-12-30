package data

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// Get all flowers from database and return them as json array
func GetAllFlowerJson() ([]byte, error) {
	flowerRows, err := DB.Query(CTX, "SELECT * FROM flower")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var flowerSlice []Flower

	rs := pgxscan.NewRowScanner(flowerRows)

	for flowerRows.Next() {
		var flower Flower

		err = rs.Scan(&flower)

		if err != nil {
			return nil, errors.New("Error scanning flower")
		}
		
		flowerSlice = append(flowerSlice, flower)
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flowerSlice)

	if err != nil {
		return nil, errors.New("Problem marshaling into JSON")
	}

	// no error
	return flowerJson, nil
}

func GetFlowerFromIDJson(flowerID int32) ([]byte, error) {
	paramAsStr := fmt.Sprint(flowerID)
	queryString := "SELECT * FROM flower WHERE flower_id = " + paramAsStr

	flowerRows, err := DB.Query(CTX, queryString)

	if err != nil {
		return nil, err
	}

	if !flowerRows.Next() {
		return nil, errors.New("Flower not found")
	}
	
	rs := pgxscan.NewRowScanner(flowerRows)

	var flower Flower

	err = rs.Scan(&flower)

	if err != nil {
		return nil, errors.New("Error scanning flower")
	}

	// Then turn it to JSON
	flowerJson, err := json.Marshal(flower)

	if err != nil {
		return nil, errors.New("Error marshaling flower")
	}

	// no error
	return flowerJson, nil	
}

// Get all clones from database and return them as json array
func GetAllPairJson() ([]byte, error) {
	pairRows, err := DB.Query(CTX, "SELECT * FROM pair")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var pairSlice []PairRln

	rs := pgxscan.NewRowScanner(pairRows)

	for pairRows.Next() {
		var pair PairRln

		err = rs.Scan(&pair)

		if err != nil {
			return nil, errors.New("Error scanning pair")
		}
		
		pairSlice = append(pairSlice, pair)
	}

	// Then turn it to JSON
	pairJSON, err := json.Marshal(pairSlice)

	if err != nil {
		return nil, errors.New("Problem marshaling into JSON")
	}

	// no error
	return pairJSON, nil
}


func GetPairFromIDJson(pairID int32) ([]byte, error) {
	paramAsStr := fmt.Sprint(pairID)
	queryString := "SELECT * FROM pair WHERE pair_id = " + paramAsStr

	pairRows, err := DB.Query(CTX, queryString)

	if err != nil {
		return nil, err
	}

	if !pairRows.Next() {
		return nil, errors.New("Pair not found")
	}
	
	rs := pgxscan.NewRowScanner(pairRows)

	var pair PairRln

	err = rs.Scan(&pair)

	if err != nil {
		return nil, errors.New("Error scanning pair")
	}

	// Then turn it to JSON
	pairJson, err := json.Marshal(pair)

	if err != nil {
		return nil, errors.New("Error marshaling pair")
	}

	// no error
	return pairJson, nil	
}



// Get all clones from database and return them as json array
func GetAllCloneJson() ([]byte, error) {
	cloneRows, err := DB.Query(CTX, "SELECT * FROM clone")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var cloneSlice []CloneRln

	rs := pgxscan.NewRowScanner(cloneRows)

	for cloneRows.Next() {
		var clone CloneRln

		err = rs.Scan(&clone)

		if err != nil {
			return nil, errors.New("Error scanning clone")
		}
		
		cloneSlice = append(cloneSlice, clone)
	}

	// Then turn it to JSON
	cloneJson, err := json.Marshal(cloneSlice)

	if err != nil {
		return nil, errors.New("Problem marshaling into JSON")
	}

	// no error
	return cloneJson, nil
}


func GetCloneFromIDJson(cloneID int32) ([]byte, error) {
	paramAsStr := fmt.Sprint(cloneID)
	queryString := "SELECT * FROM clone WHERE clone_id = " + paramAsStr

	cloneRows, err := DB.Query(CTX, queryString)

	if err != nil {
		return nil, err
	}

	if !cloneRows.Next() {
		return nil, errors.New("Clone not found")
	}
	
	rs := pgxscan.NewRowScanner(cloneRows)

	var clone CloneRln

	err = rs.Scan(&clone)

	if err != nil {
		return nil, errors.New("Error scanning clone")
	}

	// Then turn it to JSON
	cloneJson, err := json.Marshal(clone)

	if err != nil {
		return nil, errors.New("Error marshaling clone")
	}

	// no error
	return cloneJson, nil	
}

