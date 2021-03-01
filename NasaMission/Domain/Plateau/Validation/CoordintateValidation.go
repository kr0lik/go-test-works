package Validation

import (
	"go-test-works/NasaMission/Domain/Navigation"
	"go-test-works/NasaMission/Domain/Plateau"
)

func IsValidCoordinate(plateau *Plateau.Plateau, coordinate Navigation.Coordinate) bool {
	return coordinate.GetX() <= plateau.GetEndX() &&
		coordinate.GetX() >= plateau.GetStartX() &&
		coordinate.GetY() <= plateau.GetEndY() &&
		coordinate.GetY() >= plateau.GetStartY()
}
