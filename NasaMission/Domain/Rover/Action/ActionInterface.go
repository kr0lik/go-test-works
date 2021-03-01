package Action

import (
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/Domain/Rover/Enum"
)

type RoverActionInterface interface {
	Execute(rover *Rover.Rover) error
	GetCommand() Enum.Command
}
