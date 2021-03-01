package Plateau

import (
	"errors"
	"fmt"
	"go-test-works/NasaMission/Domain/Rover"
)

type Plateau struct {
	rovers                     map[int]Rover.Rover
	startX, startY, endX, endY int `default:"0"`
}

func CreatePlateau(endX int, endY int) *Plateau {
	return &Plateau{endX: endX, endY: endY, rovers: make(map[int]Rover.Rover)}
}

func (plateau *Plateau) AddRover(num int, rover Rover.Rover) error {
	if _, ok := plateau.rovers[num]; ok {
		return errors.New(fmt.Sprintf("Rover #%d already added", num))
	}

	plateau.rovers[num] = rover
	return nil
}

func (plateau *Plateau) CountRovers() int {
	return len(plateau.rovers)
}

func (plateau *Plateau) GetRover(num int) (rover *Rover.Rover, success bool) {
	r, success := plateau.rovers[num]

	return &r, success
}

func (plateau *Plateau) GetStartX() int {
	return plateau.startX
}

func (plateau *Plateau) GetStartY() int {
	return plateau.startY
}

func (plateau *Plateau) GetEndX() int {
	return plateau.endX
}

func (plateau *Plateau) GetEndY() int {
	return plateau.endY
}
