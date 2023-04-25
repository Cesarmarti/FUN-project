package structures

import (
	"github.com/Cesarmarti/FUN-project/internal/models"
)

type AntiRepetition struct {
	// each skills is pointed to it's group counter,
	// if it has more than one, we have a non-hierarhical anti repetition rule
	skillCounters  []map[string][]int
	counts         [][]int
	countsOriginal [][]int
	groups         int
	hierarhical    bool
}

func NewAntiRepetition(groups []models.AntiRepetitionGroup, routines int) AntiRepetition {
	antiRepetition := AntiRepetition{
		skillCounters:  []map[string][]int{},
		counts:         make([][]int, routines),
		countsOriginal: make([][]int, routines),
		groups:         len(groups),
	}

	for i := 0; i < routines; i++ {
		antiRepetition.skillCounters = append(antiRepetition.skillCounters, make(map[string][]int))
		antiRepetition.counts[i] = make([]int, len(groups))
		antiRepetition.countsOriginal[i] = make([]int, len(groups))
	}

	// Go over groups and point each move to the right ruleCounter(s)
	for gi, g := range groups {
		for i := 0; i < routines; i++ {
			antiRepetition.counts[i][gi] = g.K
			antiRepetition.countsOriginal[i][gi] = g.K
		}

		for _, s := range g.Skills {
			for i := 0; i < routines; i++ {
				antiRepetition.skillCounters[i][s] = append(antiRepetition.skillCounters[i][s], gi)
			}
		}
	}

	// Check if the rule is hierarhical
	antiRepetition.hierarhical = antiRepetition.isHierarhical()

	return antiRepetition
}

func (r AntiRepetition) GetBitstring(seq models.Sequence, routine int) []float64 {
	bitstring := make([]float64, len(seq))

	// Keep track of modified group counters
	modified := make([]int, len(seq))

	if r.hierarhical {
		for i, s := range seq {
			val, ok := r.skillCounters[routine][s]
			modified[i] = val[0]
			if ok {
				if r.counts[routine][val[0]] > 0 {
					bitstring[i] = 1.0
					r.counts[routine][val[0]]--
				} else {
					bitstring[i] = 0.0
				}

			} else {
				bitstring[i] = 1.0
			}

		}
	} else {
		// Non hierarhical rule
	}

	// Reset used group counters
	for _, v := range modified {
		r.counts[routine][v] = r.countsOriginal[routine][v]
	}

	return bitstring
}

// Check if each move belongs to only one rule by
func (r *AntiRepetition) isHierarhical() bool {
	for _, v := range r.skillCounters[0] {
		// move having more than 1 ruleCounter implies non-hierarhical rule
		if len(v) > 1 {
			return false
		}
	}
	return true
}

/*
func (r *AntiRepetition) ResetCounters(id int) {
	for _, sc := range r.skillCounters[id] {
		sc[0].Reset()
	}
}*/

// RuleCounter keeps a counter for each group of elements
type RuleCounter struct {
	k        int
	original int
}

func NewRuleCounter(k int) RuleCounter {
	return RuleCounter{
		k:        k,
		original: k,
	}
}

func (r *RuleCounter) Use() float64 {
	if r.k > 0 {
		r.k--
		return 1.0
	} else {
		return 0.0
	}
}

func (r *RuleCounter) Reset() {
	r.k = r.original
}
