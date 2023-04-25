package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type ElementGroup struct {
	skillGroups map[string]float64
}

func NewElementGroup(skills []models.Skill, groups []models.ElementGroupGroup) ElementGroup {
	elementGroup := ElementGroup{
		skillGroups: make(map[string]float64),
	}

	for _, g := range groups {
		for _, s := range g.Skills {
			elementGroup.skillGroups[s] = g.Value
		}
	}

	return elementGroup
}

func (e *ElementGroup) GetElementValue(skill string) float64 {
	return e.skillGroups[skill]
}
