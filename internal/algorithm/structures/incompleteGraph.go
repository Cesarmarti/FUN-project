package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type IncompleteGraph struct {
	disallowedEdges map[string]bool
}

func NewIncompleteGraph(edges []models.Edge) *IncompleteGraph {
	incompleteGraph := &IncompleteGraph{
		disallowedEdges: make(map[string]bool),
	}

	for _, e := range edges {
		key := e.S1 + e.S2
		incompleteGraph.disallowedEdges[key] = true
	}

	return incompleteGraph
}

func (g *IncompleteGraph) IsEdgeAllowed(s1, s2 string) bool {
	key := s1 + s2
	return !g.disallowedEdges[key]
}
