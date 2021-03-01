package Service

import (
	"bufio"
	"fmt"
	"go-test-works/NasaMission/Domain/Plateau"
	"go-test-works/NasaMission/Domain/Plateau/Validation"
	"go-test-works/NasaMission/Domain/Rover"
	"go-test-works/NasaMission/IO"
	RoverRunner "go-test-works/NasaMission/Service/Rover"
)

func Execute(input *bufio.Scanner, plateau *Plateau.Plateau) {
	IO.PrintDebugMessage("Getting rover")

	rover, roverNum := getRover(input, plateau)

	IO.PrintDebugMessage(fmt.Sprintf("Rover #%d control", roverNum))

	controlRover(input, rover, plateau)

	if true == IO.GetBool(input, "Control else rover?") {
		Execute(input, plateau)
	}
}

func getRover(input *bufio.Scanner, plateau *Plateau.Plateau) (rover *Rover.Rover, num int) {
	roverNum := 1

	if plateau.CountRovers() > 1 {
		roverNum = IO.GetInteger(input, "Enter rover num:")
	}

	rover, success := plateau.GetRover(roverNum)

	if false == success {
		IO.PrintErrorMessage(fmt.Sprintf("Rover #%d not exists", roverNum))

		return getRover(input, plateau)
	}

	return rover, roverNum
}

func controlRover(input *bufio.Scanner, rover *Rover.Rover, plateau *Plateau.Plateau) {
	roverRunner := RoverRunner.CreateRoverRunner()

	for isWaitAction := true; isWaitAction; {
		actions := IO.GetActions(input)
		for _, action := range actions {
			err := roverRunner.Run(rover, action)

			if nil != err {
				IO.PrintErrorMessage(err.Error())
				break
			}

			position := rover.GetPosition()
			coordinate := position.GetCoordinate()

			IO.PrintDebugMessage(fmt.Sprintf("Rover position: %d %d %s",
				coordinate.GetX(), coordinate.GetY(), position.GetDirection()))
		}

		if false == Validation.IsValidCoordinate(plateau, rover.GetPosition().GetCoordinate()) {
			IO.PrintErrorMessage(fmt.Sprintf("Rover goes out of plateau: X:(%d-%d), Y:(%d-%d)",
				plateau.GetStartX(), plateau.GetEndX(), plateau.GetStartY(), plateau.GetEndY()))
		}

		isWaitAction = IO.GetBool(input, "Do action else?")
	}
}
