package flower

import (
	"log"

	"github.com/lucasb-eyer/go-colorful"
)

// Function to take two hex color strings with leading '#' and return a hex color string
// with leading '#' that is a 50/50 blend of param colors
func blend5050(str1 string, str2 string) string {
	// Parse hex codes to Colors
	hex1, err := colorful.Hex(str1)
	
	if err != nil {
		log.Fatal(err)
	}

	hex2, err := colorful.Hex(str2)

	if err != nil {
		log.Fatal(err)
	}

	// Convert to HCL for better blend
	h1, c1, l1 := hex1.Hcl()
	hcl1 := colorful.Hcl(h1, c1, l1)

	h2, c2, l2 := hex2.Hcl()
	hcl2 := colorful.Hcl(h2, c2, l2)

	// Blend 50/50 and get value (50/50 from 0.50 argument)
	// Clamped ensures blended color is in human-vision
	blendedHex := hcl1.BlendHcl(hcl2, 0.50).Clamped().Hex()

	return blendedHex
}

// Function to take two Flower parents and breed them returning child
// Temporarily takes child flower's id as a parameter
func BreedPair(parent1 *Flower, parent2 *Flower, id uint) *Flower {
	color := blend5050(parent1.ColorPetal, parent2.ColorPetal)

	return &Flower{ID: id, ColorPetal: color}
}