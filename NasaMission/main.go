package main

import (
	"bufio"
	"go-test-works/NasaMission/Builder"
	"go-test-works/NasaMission/Service"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	plateau := Builder.BuildPlateau(reader)

	Service.Execute(reader, plateau)
}
