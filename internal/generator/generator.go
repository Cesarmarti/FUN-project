package generator

import (
	"math"

	"github.com/Cesarmarti/FUN-project/internal/models"
	"github.com/samber/lo"
)

func GenerateSequences(skills []models.Skill, length int) []models.Sequence {
	labels := lo.Map(skills, func(skill models.Skill, index int) string {
		return skill.Label
	})

	numSkills := len(skills)
	seqs := []models.Sequence{}

	for lens := 1; lens <= length; lens++ {
		numSequences := int(math.Pow(float64(numSkills), float64(lens)))
		for i := 0; i < numSequences; i++ {
			seqs = append(seqs, models.NewSequence(iToSequence(i, labels, lens)))
		}
	}

	return seqs
}

func iToSequence(n int, skills []string, length int) string {
	numSkills := len(skills)
	seq := ""

	for i := 0; i < length; i++ {
		seq += skills[n%numSkills]
		n /= numSkills
	}

	return seq
}
