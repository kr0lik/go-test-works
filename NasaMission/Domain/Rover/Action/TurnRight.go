package Action

import (
	"go-test-works/NasaMission/Domain/Navigation/Action"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/Domain/Rover/Enum"
)

type TurnRight struct {
}

func (TurnRight) Execute(rover *Rover.Rover) error {
	return Action.TurnRight(rover.GetPosition())
}

func (TurnRight) GetCommand() Enum.Command {
	return Enum.Right
}
