package generator

import (
	"math/rand"
	"sort"
	"sync"
	"time"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/models"
	"github.com/samber/lo"
)

type geneticResult struct {
	sequences []Element
}

type geneticCollector struct {
	sequences []Element
	result    chan geneticResult
	finish    chan chan geneticResult
	done      chan bool
	wg        *sync.WaitGroup
}

type Element struct {
	sequence []string
	value    float64
}

func NewGeneticCollector(routines int, wg *sync.WaitGroup) *geneticCollector {
	return &geneticCollector{
		sequences: []Element{},
		result:    make(chan geneticResult),
		finish:    make(chan chan geneticResult),
		done:      make(chan bool),
		wg:        wg,
	}
}

func (c *geneticCollector) Run() {
	for {
		select {
		case r := <-c.result:
			c.sequences = append(c.sequences, r.sequences...)
			c.wg.Done()
		case returnChan := <-c.finish:
			returnChan <- geneticResult{sequences: c.sequences}
			c.sequences = []Element{}
		case <-c.done:
			return
		}
	}
}

func Genetic(algo al.Algorithm, length int, n_pop int, maxIter int, routines int) (models.Sequence, float64) {
	rand.Seed(time.Now().UnixNano())
	skills := lo.Map(algo.Sport.Skills, func(skill models.Skill, index int) string {
		return skill.Label
	})

	var wg sync.WaitGroup

	collector := NewGeneticCollector(routines, &wg)
	go collector.Run()

	population := make([]Element, n_pop)

	for i := 0; i < n_pop; i++ {
		seq := make([]string, length)
		for i := 0; i < len(seq); i++ {
			seq[i] = skills[rand.Intn(len(skills))]
		}
		population[i] = Element{seq, 0}
	}

	// Evaluate population
	for i := 0; i < len(population); i++ {
		population[i].value = algo.Evaluate(population[i].sequence, 0)
	}

	for iter := 0; iter < maxIter; iter++ {
		// Sort population
		sort.Slice(population, func(i, j int) bool {
			return population[i].value > population[j].value
		})

		topK := int(float64(len(population)) * 0.1)

		newPopulation := population[:topK]
		generations := n_pop - topK

		chunk := generations / routines

		collector.wg.Add(routines)
		for i := 0; i < routines; i++ {
			if i < routines-1 {
				go geneticRoutine(algo, population, chunk, skills, collector.result, &wg, i)
			} else {
				go geneticRoutine(algo, population, chunk+(generations%routines), skills, collector.result, &wg, i)
			}
		}

		collector.wg.Wait()

		resultChan := make(chan geneticResult)

		collector.finish <- resultChan

		result := <-resultChan

		newPopulation = append(newPopulation, result.sequences...)

		population = newPopulation
	}

	collector.done <- true

	return population[0].sequence, population[0].value
}

func geneticRoutine(algo al.Algorithm, population []Element, iterations int, skills []string, resultChan chan geneticResult, wg *sync.WaitGroup, id int) {
	newPopulation := []Element{}
	for i := 0; i < iterations; i++ {
		parent1 := population[rand.Intn(len(population))]
		parent2 := population[rand.Intn(len(population))]
		newPopulation = append(newPopulation, MergeParents(algo, parent1, parent2, skills))
	}
	resultChan <- geneticResult{sequences: newPopulation}
}

func MergeParents(algo al.Algorithm, parent1, parent2 Element, skills []string) Element {
	var child Element
	child.sequence = make([]string, len(parent1.sequence))

	for i := 0; i < len(parent1.sequence); i++ {
		rng := rand.Float64()

		if rng < 0.45 {
			child.sequence[i] = parent1.sequence[i]
		} else if rng < 0.90 {
			child.sequence[i] = parent2.sequence[i]
		} else {
			rng2 := rand.Intn(len(skills))
			child.sequence[i] = skills[rng2]
		}

	}

	child.value = algo.Evaluate(child.sequence, 0)

	return child
}
