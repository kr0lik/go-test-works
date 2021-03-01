package Builder

import (
	"bufio"
	"fmt"
	"go-test-works/NasaMission/Domain/Navigation"
	"go-test-works/NasaMission/Domain/Plateau"
	"go-test-works/NasaMission/Domain/Plateau/Validation"
	"go-test-works/NasaMission/IO"
)

func BuildPosition(input *bufio.Scanner, plateau *Plateau.Plateau) Navigation.Position {
	IO.PrintDebugMessage("Creating position")

	coordinate := getCoordinate(input, plateau)

	direction := IO.GetDirection(input)

	return Navigation.CreatePosition(coordinate, direction)
}

func getCoordinate(input *bufio.Scanner, plateau *Plateau.Plateau) Navigation.Coordinate {
	coordinate := Navigation.CreateCoordinate(
		IO.GetInteger(input, "Enter rover start X coordinate:"),
		IO.GetInteger(input, "Enter rover start Y coordinate:"))

	if false == Validation.IsValidCoordinate(plateau, coordinate) {
		IO.PrintErrorMessage(fmt.Sprintf("Coordinates out of plateau: X:(%d-%d), Y:(%d-%d)",
			plateau.GetStartX(), plateau.GetEndX(), plateau.GetStartY(), plateau.GetEndY()))

		return getCoordinate(input, plateau)
	}

	return coordinate
}
