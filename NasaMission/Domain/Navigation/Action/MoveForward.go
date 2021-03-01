package Action

import (
	"errors"
	"fmt"
	"go-test-works/NasaMission/Domain/Navigation"
	"go-test-works/NasaMission/Domain/Navigation/Enum"
)

func MoveForward(position *Navigation.Position) error {
	var err error

	direction := position.GetDirection()
	coordinate := position.GetCoordinate()

	switch position.GetDirection() {
	case Enum.North:
		coordinate = Navigation.CreateCoordinate(coordinate.GetX(), coordinate.GetY()+1)
		break

	case Enum.East:
		coordinate = Navigation.CreateCoordinate(coordinate.GetX()+1, coordinate.GetY())
		break

	case Enum.South:
		coordinate = Navigation.CreateCoordinate(coordinate.GetX(), coordinate.GetY()-1)
		break

	case Enum.West:
		coordinate = Navigation.CreateCoordinate(coordinate.GetX()-1, coordinate.GetY())
		break

	default:
		err = errors.New(fmt.Sprintf("unknown direction %s, not accept action", direction))
	}

	*position = Navigation.CreatePosition(coordinate, direction)

	return err
}
