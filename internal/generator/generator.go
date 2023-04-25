package generator

import (
	"fmt"
	"math"
	"sync"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/models"
	"github.com/samber/lo"
)

type result struct {
	maxValue  float64
	sequences []models.Sequence
}

type collector struct {
	maxValue   float64
	sequences  []models.Sequence
	collectAll bool
	result     chan result
	finish     chan chan result
	wg         *sync.WaitGroup
}

func NewCollector(collectAll bool, routines int, wg *sync.WaitGroup) *collector {
	return &collector{
		maxValue:   math.SmallestNonzeroFloat64,
		sequences:  []models.Sequence{},
		collectAll: collectAll,
		result:     make(chan result),
		finish:     make(chan chan result),
		wg:         wg,
	}
}

func (c *collector) Run() {
	for {
		select {
		case r := <-c.result:
			if r.maxValue > c.maxValue {
				c.maxValue = r.maxValue
				c.sequences = r.sequences
			} else if r.maxValue == c.maxValue {
				if c.collectAll {
					c.sequences = append(c.sequences, r.sequences...)
				} else {
					c.sequences = r.sequences
				}
			}
			c.wg.Done()
		case returnChan := <-c.finish:
			returnChan <- result{maxValue: c.maxValue, sequences: c.sequences}
			return
		}
	}
}

func GenerateSequences(algo al.Algorithm, skills []models.Skill, max int, min int, routines int, outputAll bool) ([]models.Sequence, float64) {
	labels := lo.Map(skills, func(skill models.Skill, index int) string {
		return skill.Label
	})

	var wg sync.WaitGroup

	numSkills := len(skills)
	collector := NewCollector(outputAll, routines, &wg)
	go collector.Run()

	for lens := min; lens <= max; lens++ {
		numSequences := int(math.Pow(float64(numSkills), float64(lens)))
		fmt.Printf("Generating %d sequences of length %d\n", numSequences, lens)
		chunk := numSequences / routines
		iter := 1

		wg.Add(routines)

		for i := 0; i < numSequences; i += chunk {
			if iter != routines {
				go testSequences(algo, skills, labels, i, i+chunk, lens, outputAll, collector.result, &wg, iter)
			} else {
				if numSequences%routines != 0 {
					go testSequences(algo, skills, labels, i, i+chunk+(numSequences%routines), lens, outputAll, collector.result, &wg, iter)
				} else {
					go testSequences(algo, skills, labels, i, i+chunk, lens, outputAll, collector.result, &wg, iter)
				}
				break
			}
			iter++
		}
	}

	wg.Wait()

	resultChan := make(chan result)

	collector.finish <- resultChan

	result := <-resultChan

	return result.sequences, result.maxValue
}

func testSequences(algo al.Algorithm, skills []models.Skill, labels []string, from int, to int, length int, outputAll bool, resultChan chan result, wg *sync.WaitGroup, id int) {
	maxValue := math.SmallestNonzeroFloat64
	maxSequences := []models.Sequence{}

	for i := from; i < to; i++ {
		sequence := models.NewSequence(iToSequence(i, labels, length))
		value := algo.Evaluate(sequence, id-1)
		if value > maxValue {
			maxSequences = []models.Sequence{sequence}
			maxValue = value
		} else if value == maxValue {
			if outputAll {
				maxSequences = append(maxSequences, sequence)
			}
		}
	}

	resultChan <- result{maxValue: maxValue, sequences: maxSequences}
}

func iToSequence(n int, skills []string, length int) string {
	numSkills := len(skills)
	seq := ""

	for i := 0; i < length; i++ {
		seq += skills[n%numSkills] + "."
		n /= numSkills
	}

	return seq[:len(seq)-1]
}
