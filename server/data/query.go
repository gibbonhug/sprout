package data

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// Get all flowers from database and return them as json array
func GetAllFlower() ([]Flower, error) {
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

	// no error
	return flowerSlice, nil
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


// Get all boxes from database and return them as json array
// Entire flower which is inside box (if exists) is inside the json array
//func GetAllBoxJson() ([]byte, error) {
	//boxRows, err := DB.Query(CTX, "SELECT * FROM box")

	//if err != nil {
		//return nil, err
	//}

	// First turn the data to a slice
	//var BoxSlice []Box

	//rs := pgxscan.NewRowScanner(boxRows)

	//for boxRows.Next() {
		//var flower Flower

		//err = rs.Scan(&flower)

		//if err != nil {
			//return nil, errors.New("Error scanning box")
		//}
		
		//flowerSlice = append(flowerSlice, flower)
	//}

	// Then turn it to JSON
	//flowerJson, err := json.Marshal(flowerSlice)

	//if err != nil {
		//return nil, errors.New("Problem marshaling into JSON")
	//}

	// no error
	//return flowerJson, nil
//}



// Get all pairs from database and return them as json array
func GetAllPair() ([]PairRln, error) {
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

	return pairSlice, nil
}

func GetPairFromIDJson(pairID int32) (*PairRln, error) {
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

	// no error
	return &pair, nil	
}

// Get all clones from database and return them as json array
func GetAllCloneJson() ([]CloneRln, error) {
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

	// no error
	return cloneSlice, nil
}


func GetCloneFromIDJson(cloneID int32) (*CloneRln, error) {
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

	// no error
	return &clone, nil	
}

