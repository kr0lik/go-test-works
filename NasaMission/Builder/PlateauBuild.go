package Builder

import (
	"bufio"
	"fmt"
	"go-test-works/NasaMission/Domain/Plateau"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/IO"
)

func BuildPlateau(input *bufio.Scanner) *Plateau.Plateau {
	IO.PrintDebugMessage("Creating Plateau")

	plateau := Plateau.CreatePlateau(
		IO.GetInteger(input, "Enter the max upper(X) coordinate of the plateau:"),
		IO.GetInteger(input, "Enter the max right(Y) coordinate of the plateau:"),
	)

	IO.PrintDebugMessage(fmt.Sprintf("Plateau (%d,%d) created",
		plateau.GetEndX(), plateau.GetEndY()))

	addRovers(input, plateau)

	return plateau
}

func addRovers(input *bufio.Scanner, plateau *Plateau.Plateau) {
	roverNum := 0
	for isWaitRover := true; isWaitRover; {
		roverNum++

		if err := addRover(input, plateau, roverNum); nil != err {
			IO.PrintErrorMessage(err.Error())
		}

		isWaitRover = IO.GetBool(input, "Add rover else?")
	}
}

func addRover(input *bufio.Scanner, plateau *Plateau.Plateau, roverNum int) error {
	IO.PrintDebugMessage(fmt.Sprintf("Adding rover #%d", roverNum))

	position := BuildPosition(input, plateau)

	rover := Rover.CreateRover(position)

	if err := plateau.AddRover(roverNum, rover); nil != err {
		return err
	}

	coordinate := position.GetCoordinate()

	IO.PrintDebugMessage(fmt.Sprintf("Rover #%d with position %d %d %s added",
		roverNum, coordinate.GetX(), coordinate.GetY(), position.GetDirection()))

	return nil
}
