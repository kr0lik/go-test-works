package Action

import (
	"errors"
	"fmt"
	"go-test-works/NasaMission/Domain/Navigation"
	"go-test-works/NasaMission/Domain/Navigation/Enum"
)

func TurnRight(position *Navigation.Position) error {
	var err error

	direction := position.GetDirection()

	switch direction {
	case Enum.North:
		direction = Enum.East
		break

	case Enum.East:
		direction = Enum.South
		break

	case Enum.South:
		direction = Enum.West
		break

	case Enum.West:
		direction = Enum.North
		break

	default:
		err = errors.New(fmt.Sprintf("unknown direction %s, not accept action", direction))
	}

	*position = Navigation.CreatePosition(position.GetCoordinate(), direction)

	return err
}
