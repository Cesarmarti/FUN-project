package structures

import (
	"github.com/Cesarmarti/FUN-project/internal/models"
)

type AntiRepetition struct {
	// each skills is pointed to it's group counter,
	// if it has more than one, we have a non-hierarhical anti repetition rule
	skillCounters map[string][]*RuleCounter
	hierarhical   bool
	counters      []*RuleCounter
}

func NewAntiRepetition(groups []models.AntiRepetitionGroup) AntiRepetition {
	antiRepetition := AntiRepetition{
		skillCounters: make(map[string][]*RuleCounter),
		counters:      []*RuleCounter{},
	}

	// Go over groups and point each move to the right ruleCounter(s)
	for _, g := range groups {
		ruleCounter := NewRuleCounter(g.K)
		antiRepetition.counters = append(antiRepetition.counters, ruleCounter)
		for _, s := range g.Skills {
			antiRepetition.skillCounters[s] = append(antiRepetition.skillCounters[s], ruleCounter)
		}
	}

	// Check if the rule is hierarhical
	antiRepetition.hierarhical = antiRepetition.isHierarhical()

	return antiRepetition
}

func (r *AntiRepetition) GetBitstring(seq models.Sequence) []int {
	bitstring := make([]int, len(seq))

	if r.hierarhical {
		for i, s := range seq {
			val, ok := r.skillCounters[s]
			if ok {
				bitstring[i] = val[0].Use()
			} else {
				bitstring[i] = 1
			}

		}
	} else {
		// Non hierarhical rule
	}

	// Reset the counters
	r.ResetCounters()

	return bitstring
}

// Check if each move belongs to only one rule by
func (r *AntiRepetition) isHierarhical() bool {
	for _, v := range r.skillCounters {
		// move having more than 1 ruleCounter implies non-hierarhical rule
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func (r *AntiRepetition) ResetCounters() {
	for _, sc := range r.counters {
		sc.Reset()
	}
}

// RuleCounter keeps a counter for each group of elements
type RuleCounter struct {
	k        int
	original int
}

func NewRuleCounter(k int) *RuleCounter {
	return &RuleCounter{
		k:        k,
		original: k,
	}
}

func (r *RuleCounter) Use() int {
	if r.k > 0 {
		r.k--
		return 1
	} else {
		return 0
	}
}

func (r *RuleCounter) Reset() {
	r.k = r.original
}
