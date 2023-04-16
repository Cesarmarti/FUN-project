package algorithm

import (
	sp "github.com/Cesarmarti/FUN-project/internal/sports"
)

type Algorithm struct {
	Sport                 *sp.Sport
	Values                map[string]int
	Deductions            map[string]int
	AntiRepetitionVector  interface{} // TYPE TBD
	ElementGroupPairs     interface{} // TYPE TBD
	ConnectionTriples     interface{} // TYPE TBD
	IncompleteGraphMatrix interface{} // TYPE TBD
}

func NewAlgorithm(sport *sp.Sport) *Algorithm {
	algorithm := &Algorithm{
		Sport: sport,
	}

	// Generate values and dedutions
	values := make(map[string]int)
	deduction := make(map[string]int)

	for _, skill := range sport.Skills {
		values[skill.Label] = skill.Value
		deduction[skill.Label] = skill.Deduction
	}
	algorithm.Values = values
	algorithm.Deductions = deduction

	// Generate anti repetition vector if the rule is present
	if sport.AntiRepetitionRule != nil {
		// TODO: Implement
		algorithm.AntiRepetitionVector = nil
	}

	// Generate element group pairs if the rule is present
	if sport.ElementGroupRule != nil {
		// TODO: Implement
		algorithm.ElementGroupPairs = nil
	}

	// Generate connection triples if the rule is present
	if sport.ConnectionRule != nil {
		// TODO: Implement
		algorithm.ConnectionTriples = nil
	}

	// Generate incomplete graph matrix if the rule is present
	if sport.IncompleteGraphRule != nil {
		// TODO: Implement
		algorithm.IncompleteGraphMatrix = nil
	}

	return algorithm
}

// Evaluate the sequence according to the rules of the sport
func (a *Algorithm) Evaluate(seq Sequence) int {
	value := a.CalculateBasicRule(seq)

	// Calculate anti repetition rule if preset
	if a.Sport.AntiRepetitionRule != nil {
		value = value + a.CalculateAntiRepetitionRule(seq)
	}

	// Calculate element group rule if preset
	if a.Sport.ElementGroupRule != nil {
		value = value + a.CalculateElementGroupRule(seq)
	}

	// Calculate connection rule if preset
	if a.Sport.ConnectionRule != nil {
		value = value + a.CalculateConnectionRule(seq)
	}

	// Calculate incomplete graph rule if preset
	if a.Sport.IncompleteGraphRule != nil {
		value = value + a.CalculateIncompleteGraphRule(seq)
	}

	return value
}
