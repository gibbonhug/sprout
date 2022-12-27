package flower

import (

)

// Function to take a flower and clone it
// Temporarily takes child's ID as a param
func BreedClone(parent *Flower, id uint) *Flower {
	return &Flower{ID: id, ColorPetal: parent.ColorPetal}
}