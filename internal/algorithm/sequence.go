package algorithm

import (
	"strings"

	sp "github.com/Cesarmarti/FUN-project/internal/sports"
	"github.com/samber/lo"
)

// Sequence of moves in a routine
type Sequence []string

func NewSequence(seq string) []string {
	return strings.Split(seq, "")
}

func (a *Algorithm) ValidateSequence(sequence Sequence) bool {
	// Check that every move is valid
	for _, move := range sequence {
		check := lo.ContainsBy(a.Sport.Skills, func(s sp.Skill) bool {
			return s.Label == move
		})
		if !check {
			return false
		}
	}

	return true
}
