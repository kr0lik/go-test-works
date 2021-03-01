package IO

import (
	"bufio"
	"fmt"
	NavigationEnum "go-test-works/NasaMission/Domain/Navigation/Enum"
	RoverEnum "go-test-works/NasaMission/Domain/Rover/Enum"
	"os"
	"strconv"
	"strings"
)

func GetInteger(input *bufio.Scanner, message string) int {
	PrintInfoMessage(message)

	input.Scan()
	if input.Err() != nil {
		PrintErrorMessage(input.Err().Error())
		os.Exit(3)
	}

	if "" == input.Text() {
		PrintErrorMessage("Incorrect input: " + input.Text())
		return GetInteger(input, message)
	}

	number, err := strconv.Atoi(input.Text())

	if err != nil {
		PrintErrorMessage(err.Error())
		return GetInteger(input, message)
	}

	return number
}

func GetBool(input *bufio.Scanner, message string) bool {
	PrintInfoMessage(message + (" (y/N)"))

	input.Scan()
	if input.Err() != nil {
		PrintErrorMessage(input.Err().Error())
		os.Exit(3)
	}

	res := "y" == input.Text()

	if false == res && "N" != input.Text() {
		PrintErrorMessage("Incorrect input: " + input.Text())
		return GetBool(input, message)
	}

	return res
}

func GetActions(input *bufio.Scanner) []RoverEnum.Command {
	var variants []string

	for _, variant := range new(RoverEnum.Command).GetValues() {
		variants = append(variants, string(variant))
	}

	PrintInfoMessage("Enter action batch: (Available actions " + strings.Join(variants, ",") + ")")

	input.Scan()
	if input.Err() != nil {
		PrintErrorMessage(input.Err().Error())
		os.Exit(3)
	}

	var actions []RoverEnum.Command
	var incorrects []string

	batch := strings.Split(input.Text(), "")
	for _, v := range batch {
		action := RoverEnum.Command(v)

		if !action.IsValid() {
			incorrects = append(incorrects, v)
			continue
		}

		actions = append(actions, action)
	}

	if len(incorrects) > 0 {
		PrintErrorMessage(fmt.Sprint("Incorrect action(s): " + strings.Join(incorrects, ",") + "."))
		return GetActions(input)
	}

	return actions
}

func GetDirection(input *bufio.Scanner) NavigationEnum.Direction {
	var variants []string

	for _, variant := range new(NavigationEnum.Direction).GetValues() {
		variants = append(variants, string(variant))
	}

	PrintInfoMessage("Enter rover start direction: (" + strings.Join(variants, ",") + ")")

	input.Scan()
	if input.Err() != nil {
		PrintErrorMessage(input.Err().Error())
		os.Exit(3)
	}

	direction := NavigationEnum.Direction(input.Text())

	if !direction.IsValid() {
		PrintErrorMessage(fmt.Sprint("Incorrect action " + direction + "."))
		return GetDirection(input)
	}

	return direction
}
