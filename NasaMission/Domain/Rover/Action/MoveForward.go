package Action

import (
	"go-test-works/NasaMission/Domain/Navigation/Action"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/Domain/Rover/Enum"
)

type MoveForward struct {
}

func (MoveForward) Execute(rover *Rover.Rover) error {
	return Action.MoveForward(rover.GetPosition())
}

func (MoveForward) GetCommand() Enum.Command {
	return Enum.Move
}
