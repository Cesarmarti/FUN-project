package main

import (
	"flag"
	"fmt"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/generator"
	"github.com/Cesarmarti/FUN-project/internal/models"
)

func main() {
	fileFlag := flag.String("file", "", "path to config file")
	sequenceFlag := flag.String("seq", "", "sequence to evaluate given as string")
	generatorFlag := flag.Int("gen", 0, "upper length of sequences to generate")
	generatorMinimumFlag := flag.Int("gen-min", 1, "minimum length of sequences to generate")

	flag.Parse()

	filePath := ""

	if *fileFlag != "" {
		filePath = *fileFlag
	} else {
		flag.Usage()
		return
	}

	sport, err := models.ParseSport(filePath)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Sport: %s\n", sport.Discipline)
	}

	algorithm := al.NewAlgorithm(sport)

	if *sequenceFlag != "" {
		seq := *sequenceFlag
		sequence := models.NewSequence(seq)

		valid := algorithm.ValidateSequence(sequence)
		if !valid {
			fmt.Println("Invalid sequence")
			return
		}

		value := algorithm.Evaluate(sequence)
		fmt.Printf("Value of sequence %s : %v\n", seq, value)
	}

	if *generatorFlag != 0 {
		if *generatorFlag < *generatorMinimumFlag {
			fmt.Println("lower bound should be even or less than greater bound for generation length")
			return
		}
		seqs := generator.GenerateSequences(sport.Skills, *generatorFlag, *generatorMinimumFlag)
		maxSequences, maxValue := generator.TestSequences(&algorithm, seqs)
		fmt.Printf("Optimal sequence(s): %v\n", maxSequences)
		fmt.Printf("Value of optimal sequence(s): %v\n", maxValue)
	}

}
