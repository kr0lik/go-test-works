package IO

import "fmt"

const (
	debugColor = "\033[0;36m%s\033[0m"
	infoColor  = "\033[1;34m%s\033[0m"
	errorColor = "\033[1;31m%s\033[0m"
)

func PrintInfoMessage(message string) {
	fmt.Println(fmt.Sprintf(infoColor, message))
}

func PrintDebugMessage(message string) {
	fmt.Println(fmt.Sprintf(debugColor, message))
}

func PrintErrorMessage(message string) {
	fmt.Println(fmt.Sprintf(errorColor, message))
}
