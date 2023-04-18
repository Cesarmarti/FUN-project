package models

import (
	"strings"
)

// Sequence of moves in a routine
type Sequence []string

func NewSequence(seq string) []string {
	return strings.Split(seq, "")
}
