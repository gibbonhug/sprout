/*
	Package implements functions and structs for flower-based functions
	ie Flowers, flower boxes, relationships
*/
package data

// Flowers
type Flower struct {
	ID         uint    `json:"id"` // ID of the flower
	ColorPetal string `json:"colorPetal"` // Hex color with leading '#'
}

// Boxes flowers are stored in
type Box struct {
	ID uint `json:"id"` // ID of the box
	Flower *Flower `json:"flower"` // Pointer to the Flower it contains
}

// Cloning a flower producing a child
type CloneRln struct {
	ID uint `json:"ID"` // ID of this relationship
	ParentID uint `json:"parentID"` // ID of the parent flower
	ChildID uint `json:"childID"` // ID of the child flower
}

// Breeding 2 flowers producing a child
type PairRln struct {
	ID uint `json:"ID"` // ID of this relationship
	Parent1ID uint `json:"parent1ID"` // ID of instigator parent flower
	Parent2ID uint `json:"parent2ID"` // ID of the receiver parent flower
	ChildID uint `json:"childID"` // ID of the child flower
}
