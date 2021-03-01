package Rover

import (
	"go-test-works/NasaMission/Domain/Navigation"
)

type Rover struct {
	position Navigation.Position
}

func CreateRover(position Navigation.Position) Rover {
	return Rover{position: position}
}

func (rover *Rover) GetPosition() *Navigation.Position {
	return &rover.position
}
