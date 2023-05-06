package generator

import (
	"math"
	"math/rand"
	"sync"
	"time"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/models"
	"github.com/samber/lo"
)

type annealingResult struct {
	maxValue float64
	sequence []string
}

type annealingCollector struct {
	maxValue float64
	sequence []string
	result   chan annealingResult
	finish   chan chan annealingResult
	done     chan bool
	wg       *sync.WaitGroup
}

func NewAneallingCollector(routines int, wg *sync.WaitGroup) *annealingCollector {
	return &annealingCollector{
		maxValue: math.SmallestNonzeroFloat64,
		sequence: []string{},
		result:   make(chan annealingResult),
		finish:   make(chan chan annealingResult),
		wg:       wg,
		done:     make(chan bool),
	}
}

func (c *annealingCollector) Run() {
	for {
		select {
		case r := <-c.result:
			if r.maxValue > c.maxValue {
				c.maxValue = r.maxValue
				c.sequence = r.sequence
			}
			c.wg.Done()
		case returnChan := <-c.finish:
			returnChan <- annealingResult{maxValue: c.maxValue, sequence: c.sequence}
		case <-c.done:
			return
		}

	}
}

func Annealing(algo al.Algorithm, length int, maxIter int, alpha, tMin float64, routines int, branching int) (models.Sequence, float64) {
	rand.Seed(time.Now().UnixNano())
	skills := lo.Map(algo.Sport.Skills, func(skill models.Skill, index int) string {
		return skill.Label
	})

	var wg sync.WaitGroup

	collector := NewAneallingCollector(routines, &wg)
	go collector.Run()

	bestSequence := make([]string, length)
	for i := 0; i < len(bestSequence); i++ {
		bestSequence[i] = skills[rand.Intn(len(skills))]
	}

	bestCost := algo.Evaluate(bestSequence, 0)
	t := 0.5

	iter := 0

	for ; iter < maxIter; iter++ {
		newSequence, newCost := generateNeighbor(algo, bestSequence, bestCost, skills, t, collector, branching, routines)

		if newCost > bestCost || rand.Float64() < t {
			bestSequence = newSequence
			bestCost = newCost
		}

		if iter%10 == 0 {
			t *= alpha
		}
	}

	collector.done <- true

	return bestSequence, bestCost
}

func generateNeighbor(algo al.Algorithm, sequence []string, score float64, skills []string, t float64, collector *annealingCollector, branching int, routines int) ([]string, float64) {
	collector.wg.Add(routines)
	for i := 0; i < routines; i++ {
		go annealingRoutine(algo, sequence, score, i, branching, skills, t, collector.result)
	}

	collector.wg.Wait()

	resultChan := make(chan annealingResult)

	collector.finish <- resultChan

	result := <-resultChan

	return result.sequence, result.maxValue
}

func annealingRoutine(algo al.Algorithm, sequence []string, bestScore float64, routine int, branching int, skills []string, t float64, resultChan chan annealingResult) {
	rand.Seed(time.Now().UnixNano() * int64(routine+1))
	bestSequence := make([]string, len(sequence))
	copy(bestSequence, sequence)

	for i := 0; i < branching; i++ {
		newSequence := make([]string, len(sequence))
		copy(newSequence, bestSequence)
		newSkill := rand.Intn(len(skills))
		index := rand.Intn(len(sequence))
		newSequence[index] = skills[newSkill]
		newScore := algo.Evaluate(newSequence, routine)
		if newScore > bestScore || rand.Float64() < t {
			bestSequence = newSequence
			bestScore = newScore
		}
	}

	resultChan <- annealingResult{maxValue: bestScore, sequence: bestSequence}
}
