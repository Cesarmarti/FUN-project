package structures

import "github.com/Cesarmarti/FUN-project/internal/models"

type Connection struct {
	connectionPairs map[string]int
}

func NewConnection(connections []models.Connection) *Connection {
	connection := &Connection{
		connectionPairs: make(map[string]int),
	}

	for _, c := range connections {
		key := c.S1 + c.S2
		connection.connectionPairs[key] = c.Value
	}

	return connection
}

// Returns value for given pair of skills
func (c *Connection) GetConnectionValue(s1, s2 string) int {
	key := s1 + s2
	// Returns int zero value(0) if key not present
	value := c.connectionPairs[key]

	return value
}
