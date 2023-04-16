package main

import (
	"fmt"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	sp "github.com/Cesarmarti/FUN-project/internal/sports"
)

func main() {
	// TODO: Determine and implement run parameters
	// PARAMS:
	// filePath
	// option to give a sequence or run the generator for given length

	filePath := "skiing.json"
	seq := "abcdefg"

	sport, err := sp.ParseSport(filePath)
	if err != nil {
		fmt.Println(err)
	}

	algorithm := al.NewAlgorithm(sport)
	sequence := al.NewSequence(seq)

	value := algorithm.Evaluate(sequence)

	fmt.Printf("Sequence value: %v\n", value)
}
