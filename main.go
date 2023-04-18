package main

import (
	"fmt"
	"os"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/models"
)

func main() {
	// TODO: Add the option to run sequence generator instead

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Invalid arguments")
		fmt.Println("Correct usage: ")
		fmt.Println("go run main.go <configFilePath> <sequence>")
		fmt.Println("OR")
		fmt.Println("./fun-project <configFilePath> <sequence>")
		return
	}

	filePath := args[0]
	seq := args[1]

	sport, err := models.ParseSport(filePath)
	if err != nil {
		fmt.Println(err)
	}

	algorithm := al.NewAlgorithm(sport)
	sequence := models.NewSequence(seq)

	valid := algorithm.ValidateSequence(sequence)
	if !valid {
		fmt.Println("Invalid sequence")
		return
	}

	value := algorithm.Evaluate(sequence)

	fmt.Printf("Sequence value: %v\n", value)
}
