package algorithm

import (
	"github.com/Cesarmarti/FUN-project/internal/models"
)

// Calculates the basic value of sequence
func (a *Algorithm) CalculateBasicRule(seq models.Sequence) int {
	value := 0

	for _, s := range seq {
		value += (a.Values[s] - a.Deductions[s])
	}

	return value
}

// Calculates the value of sequence according to the anti repetition rule formula
func (a *Algorithm) CalculateAntiRepetitionRule(seq models.Sequence) int {
	value := 0
	bitstring := a.Structures.AntiRepetition.GetBitstring(seq)

	for i, s := range seq {
		value += (a.Values[s]*bitstring[i] - a.Deductions[s])
	}

	return value
}

// Calculates the value of sequence according to the element group rule formula
func (a *Algorithm) CalculateElementGroupRule(seq models.Sequence) int {
	value := 0

	for _, s := range seq {
		value += a.Structures.ElementGroup.GetElementValue(s)
	}

	return value
}

// Calculates the value of sequence according to the connection rule formula
func (a *Algorithm) CalculateConnectionRule(seq models.Sequence) int {
	value := 0

	for i := 0; i < len(seq)-1; i++ {
		value += a.Structures.Connection.GetConnectionValue(seq[i], seq[i+1])
	}

	return value
}

// Calculates the value of sequence according to the incomplete graph rule formula
func (a *Algorithm) CalculateIncompleteGraphRule(seq models.Sequence) int {
	for i := 0; i < len(seq)-1; i++ {
		if !a.Structures.IncompleteGraph.IsEdgeAllowed(seq[i], seq[i+1]) {
			return 0
		}
	}

	// Return 1 if the sequence is allowed, meaning no disallowed edges were found
	return 1
}
