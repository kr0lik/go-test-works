package Action

import (
	"go-test-works/NasaMission/Domain/Navigation/Action"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/Domain/Rover/Enum"
)

type TurnLeft struct {
}

func (TurnLeft) Execute(rover *Rover.Rover) error {
	return Action.TurnLeft(rover.GetPosition())
}

func (TurnLeft) GetCommand() Enum.Command {
	return Enum.Left
}
