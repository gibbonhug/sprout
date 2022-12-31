package data

import (
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

func GetFlowerFromID(flowerID int32) (*Flower, error) {
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

	return &flower, nil
}


// Get all boxes from database and return them as json array
// Entire flower which is inside box (if exists) is inside the json array
func GetAllBox() ([]Box, error) {
	boxRows, err := DB.Query(CTX, "SELECT * FROM box")

	if err != nil {
		return nil, err
	}

	// First turn the data to a slice
	var boxSlice []Box

	// Special logic for adding flower struct to box
	for boxRows.Next() {
		values, err := boxRows.Values()
		
		if err != nil {
			return nil, errors.New("Error scanning box")
		}

		var box Box

		boxID := values[0].(int32)
		box.ID = boxID

		// Type assertion to check if foregin key is not null

		flowerID, flowerExists := values[1].(int32)

		var flower *Flower 

		if flowerExists {
			flower, err = GetFlowerFromID(flowerID)
			box.Flower = flower

			if err != nil {
				return nil, err
			}
		}
		
		boxSlice = append(boxSlice, box)
	}

	// no error
	return boxSlice, nil
}

func GetBoxFromID(boxID int32) (*Box, error) {
	paramAsStr := fmt.Sprint(boxID)
	queryString := "SELECT * FROM box WHERE box_id = " + paramAsStr

	boxRows, err := DB.Query(CTX, queryString)

	if err != nil {
		return nil, err
	}

	if !boxRows.Next() {
		return nil, errors.New("Box not found")
	}
	
	values, err := boxRows.Values()
		
	if err != nil {
		return nil, errors.New("Error scanning box")
	}

	var box Box

	boxIDdata := values[0].(int32)
	box.ID = boxIDdata

	// Type assertion to check if foregin key is not null

	flowerID, flowerExists := values[1].(int32)

	var flower *Flower 

	if flowerExists {
		flower, err = GetFlowerFromID(flowerID)
		box.Flower = flower

		if err != nil {
			return nil, err
		}
	}

	return &box, nil
}



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

func GetPairFromID(pairID int32) (*PairRln, error) {
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
func GetAllClone() ([]CloneRln, error) {
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


func GetCloneFromID(cloneID int32) (*CloneRln, error) {
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

