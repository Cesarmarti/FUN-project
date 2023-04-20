package generator

import (
	"math"

	al "github.com/Cesarmarti/FUN-project/internal/algorithm"
	"github.com/Cesarmarti/FUN-project/internal/models"
)

func TestSequences(algo *al.Algorithm, sequences []models.Sequence) ([]models.Sequence, int) {
	maxValue := math.MinInt

	maxSequences := []models.Sequence{}

	for _, sequence := range sequences {
		value := algo.Evaluate(sequence)
		if value > maxValue {
			maxSequences = []models.Sequence{sequence}
			maxValue = value
		} else if value == maxValue {
			maxSequences = append(maxSequences, sequence)
		}
	}

	return maxSequences, maxValue
}
