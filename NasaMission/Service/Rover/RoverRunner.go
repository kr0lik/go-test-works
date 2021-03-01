package Rover

import (
	"errors"
	"fmt"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/Domain/Rover/Action"
	"go-test-works/NasaMission/Domain/Rover/Enum"
)

type roverRunner struct {
	actions []Action.RoverActionInterface
}

func CreateRoverRunner() *roverRunner {
	return &roverRunner{
		actions: []Action.RoverActionInterface{
			new(Action.MoveForward),
			new(Action.TurnLeft),
			new(Action.TurnRight),
		},
	}
}

func (r *roverRunner) Run(rover *Rover.Rover, command Enum.Command) error {
	for _, action := range r.actions {
		if command == action.GetCommand() {
			return action.Execute(rover)
		}
	}

	return errors.New(fmt.Sprintf("unknown action %s, not accept action", command))
}
