package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"runtime"
	"time"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/generator"
	"github.com/Cesarmarti/FUN-project/internal/models"
)

var errFailed error = errors.New("failed to execute")

func main() {

	fileFlag := flag.String("file", "", "path to config file")
	sequenceFlag := flag.String("seq", "", "sequence to evaluate given as string")
	generatorFlag := flag.Int("gen", 0, "upper length of sequences to generate")
	generatorMinimumFlag := flag.Int("gen-min", 1, "minimum length of sequences to generate")
	printAllFlag := flag.Bool("print-all", false, "print all sequences")
	annealingFlag := flag.Int("ann", 0, "use annealing generation for sequences of given length")
	geneticFlag := flag.Int("genetic", 0, "use genetic generation for sequences of given length")

	flag.Parse()

	execute(fileFlag, sequenceFlag, generatorFlag, generatorMinimumFlag, printAllFlag, annealingFlag, geneticFlag)
}

func execute(fileFlag *string, sequenceFlag *string, generatorFlag *int, generatorMinimumFlag *int, printAllFlag *bool, annealingFlag *int, geneticFlag *int) error {
	start := time.Now()
	filePath := ""

	threads := runtime.NumCPU()
	// Divide by 2 assuming threads are 2 per core, hyper threading is useless a the workload is the same
	// Take 2, first serves as main, second is used for collector
	routines := threads/2 - 2

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
		fmt.Printf("Sequence value: %.2f\n", maxValue)
	}

	if *annealingFlag != 0 {
		aStart := time.Now()
		// Routines add work, more means longer execution time
		seq, cost := generator.Annealing(algorithm, *annealingFlag, 1000, 0.995, math.Exp(-20), routines, 100)
		fmt.Println("----------------------------------------------")
		fmt.Println("Annealing generation:")
		fmt.Printf("Optimal sequence: %v\n", seq)
		fmt.Printf("Sequence value: %.2f\n", cost)
		fmt.Printf("Exectution time: %v\n", time.Since(aStart))
	}

	if *geneticFlag != 0 {
		gStart := time.Now()
		seq, cost := generator.Genetic(algorithm, *geneticFlag, 1000, 1000, routines)
		fmt.Println("----------------------------------------------")
		fmt.Println("Genetic generation:")
		fmt.Printf("Optimal sequence: %v\n", seq)
		fmt.Printf("Sequence value: %.2f\n", cost)
		fmt.Printf("Exectution time: %v\n", time.Since(gStart))
	}

	elapsed := time.Since(start)
	fmt.Println("----------------------------------------------")
	fmt.Printf("Total execution time: %s", elapsed)
	return nil
}
