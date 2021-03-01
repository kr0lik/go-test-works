package Action

import (
	"errors"
	"fmt"
	"go-test-works/NasaMission/Domain/Navigation"
	"go-test-works/NasaMission/Domain/Navigation/Enum"
)

func TurnLeft(position *Navigation.Position) error {
	var err error

	direction := position.GetDirection()

	switch direction {
	case Enum.North:
		direction = Enum.West
		break

	case Enum.East:
		direction = Enum.North
		break

	case Enum.South:
		direction = Enum.East
		break

	case Enum.West:
		direction = Enum.South
		break

	default:
		err = errors.New(fmt.Sprintf("unknown direction %s, not accept action", direction))
	}

	*position = Navigation.CreatePosition(position.GetCoordinate(), direction)

	return err
}
