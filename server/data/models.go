package data

// Flowers
type Flower struct {
	ID         int32    `json:"id" db:"flower_id"` // ID of the flower
	ColorPetal string `json:"colorPetal" db:"petal_color"` // Hex color with leading '#'
}

// Boxes flowers are stored in
type Box struct {
	ID int32 `json:"id" db:"box_id"` // ID of the box
	Flower *Flower `json:"flower"` // Pointer to the Flower it contains
}

// Cloning a flower producing a child
type CloneRln struct {
	ID int32 `json:"ID" db:"clone_id"` // ID of this relationship
	ParentID int32 `json:"parentID" db:"parent_id"` // ID of the parent flower
	ChildID int32 `json:"childID" db:"child_id"` // ID of the child flower
}

// Breeding 2 flowers producing a child
type PairRln struct {
	ID int32 `json:"ID" db:"pair_id"` // ID of this relationship
	Parent1ID int32 `json:"parent1ID" db:"parent1_id"` // ID of instigator parent flower
	Parent2ID int32 `json:"parent2ID" db:"parent2_id"` // ID of the receiver parent flower
	ChildID int32 `json:"childID" db:"child_id"` // ID of the child flower
}
