package Navigation

import (
	"go-test-works/NasaMission/Domain/Navigation/Enum"
)

type Position struct {
	coordinate Coordinate
	direction  Enum.Direction
}

func CreatePosition(coordinate Coordinate, direction Enum.Direction) Position {
	return Position{coordinate: coordinate, direction: direction}
}

func (position *Position) GetCoordinate() Coordinate {
	return position.coordinate
}

func (position *Position) GetDirection() Enum.Direction {
	return position.direction
}
