package main

import (
	"github.com/gibbonhug/sprout/flower"
)

// This file functions as a temporary pseudo-database for the application.

// Dummy flowers to delete later
var Flowers = []flower.Flower {
	{ColorPetal: "000000", ID: 0}, {ColorPetal: "ff0000", ID: 1},
	{ColorPetal: "800000", ID: 2},
}

// Dummy boxes
var Boxes = []flower.Box {
	{ID: 0, Flower: &Flowers[0]}, {ID: 1, Flower: &Flowers[1]}, {ID: 2, Flower: &Flowers[2]}, {ID: 3},
	{ID: 4}, {ID: 5}, {ID: 6}, {ID: 7},
}

// Dummy relationships
var Relationships = []flower.BreedRelationship {
	{Parent1ID: 0, Parent2ID: 1, ChildID: 2},
}

