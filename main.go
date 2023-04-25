package main

import (
	"errors"
	"flag"
	"fmt"
	"runtime"
	"time"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/generator"
	"github.com/Cesarmarti/FUN-project/internal/models"
)

var errFailed error = errors.New("Failed to execute")

func main() {

	fileFlag := flag.String("file", "", "path to config file")
	sequenceFlag := flag.String("seq", "", "sequence to evaluate given as string")
	generatorFlag := flag.Int("gen", 0, "upper length of sequences to generate")
	generatorMinimumFlag := flag.Int("gen-min", 1, "minimum length of sequences to generate")
	printAllFlag := flag.Bool("print-all", false, "print all sequences")

	flag.Parse()

	execute(fileFlag, sequenceFlag, generatorFlag, generatorMinimumFlag, printAllFlag)
}

func execute(fileFlag *string, sequenceFlag *string, generatorFlag *int, generatorMinimumFlag *int, printAllFlag *bool) error {
	start := time.Now()
	filePath := ""

	cpus := runtime.NumCPU()
	routines := cpus

	if *fileFlag != "" {
		filePath = *fileFlag
	} else {
		flag.Usage()
		return errFailed
	}

	sport, err := models.ParseSport(filePath)
	if err != nil {
		fmt.Println(err)
		return errFailed
	} else {
		fmt.Printf("Sport: %s\n", sport.Discipline)
	}

	algorithm := al.NewAlgorithm(sport, routines)

	if *sequenceFlag != "" {
		seq := *sequenceFlag
		sequence := models.NewSequence(seq)

		valid := algorithm.ValidateSequence(sequence)
		if !valid {
			fmt.Println("Invalid sequence")
			return errFailed
		}

		value := algorithm.Evaluate(sequence, 0)
		fmt.Printf("Value of sequence %s : %v\n", seq, value)
	}

	if *generatorFlag != 0 {
		if *generatorFlag < *generatorMinimumFlag {
			fmt.Println("lower bound should be even or less than greater bound for generation length")
			return errFailed
		}
		maxSequences, maxValue := generator.GenerateSequences(algorithm, sport.Skills, *generatorFlag, *generatorMinimumFlag, routines, *printAllFlag)
		fmt.Printf("Optimal sequence(s): %v\n", maxSequences)
		fmt.Printf("Value of optimal sequence(s): %.2f\n", maxValue)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
	return nil
}
