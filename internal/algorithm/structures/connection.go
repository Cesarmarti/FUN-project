package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type Connection struct {
	connectionPairs map[string]float64
}

func NewConnection(connections []models.Connection) Connection {
	connection := Connection{
		connectionPairs: make(map[string]float64),
	}

	for _, c := range connections {
		key := c.S1 + c.S2
		connection.connectionPairs[key] = c.Value
	}

	return connection
}

// Returns value for given pair of skills
func (c *Connection) GetConnectionValue(s1, s2 string) float64 {
	key := s1 + s2
	// Returns int zero value(0) if key not present
	value := c.connectionPairs[key]

	return value
}
