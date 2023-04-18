package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type Structures struct {
	AntiRepetition  *AntiRepetition
	ElementGroup    *ElementGroup
	Connection      *Connection
	IncompleteGraph *IncompleteGraph
}

func NewStructures() *Structures {
	return &Structures{}
}

func (s *Structures) InitAntiRepetition(groups []models.AntiRepetitionGroup) {
	s.AntiRepetition = NewAntiRepetition(groups)
}

func (s *Structures) InitElementGroup(skills []models.Skill, groups []models.ElementGroupGroup) {
	s.ElementGroup = NewElementGroup(skills, groups)
}

func (s *Structures) InitConnection(connections []models.Connection) {
	s.Connection = NewConnection(connections)
}

func (s *Structures) InitIncompleteGraph(edges []models.Edge) {
	s.IncompleteGraph = NewIncompleteGraph(edges)
}
