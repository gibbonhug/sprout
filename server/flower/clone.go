package flower

import (

)

// Function to take a flower and clone it
// Temporarily takes child's ID as a param
func BreedClone(parent *Flower, id uint) *Flower {
	child := Flower{}

	child.ID = id
	child.ColorPetal = parent.ColorPetal

	return &child
}