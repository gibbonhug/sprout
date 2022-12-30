/*
	Package implements functions for flowers
*/
package flower

import (
	"github.com/gibbonhug/sprout/data"
)

// Function to take a flower and clone it
// Temporarily takes child's ID as a param
func BreedClone(parent *data.Flower, id uint) *data.Flower {
	return &data.Flower{ID: id, ColorPetal: parent.ColorPetal}
}