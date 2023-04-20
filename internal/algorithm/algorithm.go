package algorithm

import (
	"github.com/Cesarmarti/FUN-project/internal/algorithm/structures"
	"github.com/Cesarmarti/FUN-project/internal/models"
	"github.com/samber/lo"
)

type Algorithm struct {
	Sport      models.Sport
	Values     map[string]int
	Deductions map[string]int
	Structures structures.Structures
}

func NewAlgorithm(sport models.Sport) Algorithm {
	algorithm := Algorithm{
		Sport: sport,
	}

	// Initialize values and dedutions
	values := make(map[string]int)
	deduction := make(map[string]int)

	for _, skill := range sport.Skills {
		values[skill.Label] = skill.Value
		deduction[skill.Label] = skill.Deduction
	}
	algorithm.Values = values
	algorithm.Deductions = deduction

	// Initialize structures for each rule that is present
	structures := structures.NewStructures()
	if sport.AntiRepetitionRule != nil {
		structures.InitAntiRepetition(sport.AntiRepetitionRule.Groups)
	}

	if sport.ElementGroupRule != nil {
		structures.InitElementGroup(sport.Skills, sport.ElementGroupRule.Groups)
	}

	if sport.ConnectionRule != nil {
		structures.InitConnection(sport.ConnectionRule.Connections)
	}

	if sport.IncompleteGraphRule != nil {
		structures.InitIncompleteGraph(sport.IncompleteGraphRule.Edges)
	}

	algorithm.Structures = structures

	return algorithm
}

// Evaluate the sequence according to the rules of the sport
func (a *Algorithm) Evaluate(seq models.Sequence) int {
	value := 0

	// Calculate anti repetition rule if preset, otherwise calculate basic rule
	if a.Sport.AntiRepetitionRule != nil {
		// Anti repetition rule already includes the basic rule
		value += a.CalculateAntiRepetitionRule(seq)
	} else {
		value += a.CalculateBasicRule(seq)
	}

	// Calculate element group rule if preset
	if a.Sport.ElementGroupRule != nil {
		value += a.CalculateElementGroupRule(seq)
	}

	// Calculate connection rule if preset
	if a.Sport.ConnectionRule != nil {
		value += a.CalculateConnectionRule(seq)
	}

	// Calculate incomplete graph rule if preset
	if a.Sport.IncompleteGraphRule != nil {
		// Multiply routine with 0 if a disallowed pair of moves was found, otherwise 1
		value *= a.CalculateIncompleteGraphRule(seq)
	}

	return value
}

func (a *Algorithm) ValidateSequence(sequence models.Sequence) bool {
	// Check that every move is valid
	for _, move := range sequence {
		check := lo.ContainsBy(a.Sport.Skills, func(s models.Skill) bool {
			return s.Label == move
		})
		if !check {
			return false
		}
	}

	return true
}
