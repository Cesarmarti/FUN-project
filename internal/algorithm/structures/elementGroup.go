package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type ElementGroup struct {
	skillGroups map[string]int
}

func NewElementGroup(skills []models.Skill, groups []models.ElementGroupGroup) *ElementGroup {
	elementGroup := &ElementGroup{
		skillGroups: make(map[string]int),
	}

	// Init skill values to 0
	for _, s := range skills {
		elementGroup.skillGroups[s.Name] = 0
	}

	for _, g := range groups {
		for _, s := range g.Skills {
			elementGroup.skillGroups[s] = g.Value
		}
	}

	return elementGroup
}

func (e *ElementGroup) GetElementValue(skill string) int {
	return e.skillGroups[skill]
}
