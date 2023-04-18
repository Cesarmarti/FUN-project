package models

import (
	"encoding/json"
	"io/ioutil"
)

type Sport struct {
	Discipline          string               `json:"discipline"`
	Skills              []Skill              `json:"skills"`
	AntiRepetitionRule  *AntiRepetitionRule  `json:"antiRepetitionRule"`
	ElementGroupRule    *ElementGroupRule    `json:"elementGroupRule"`
	ConnectionRule      *ConnectionRule      `json:"connectionRule"`
	IncompleteGraphRule *IncompleteGraphRule `json:"incompleteGraphRule"`
}

type Skill struct {
	Name      string `json:"name"`
	Label     string `json:"label"`
	Value     int    `json:"value"`
	Deduction int    `json:"deduction"`
}

type AntiRepetitionRule struct {
	Groups []AntiRepetitionGroup `json:"groups"`
}

type AntiRepetitionGroup struct {
	K      int      `json:"k"`
	Skills []string `json:"skills"`
}

type ElementGroupRule struct {
	Groups []ElementGroupGroup `json:"groups"`
}

type ElementGroupGroup struct {
	Value  int      `json:"value"`
	Skills []string `json:"skills"`
}

type ConnectionRule struct {
	Connections []Connection `json:"connections"`
}

type Connection struct {
	S1    string `json:"s1"`
	S2    string `json:"s2"`
	Value int    `json:"value"`
}

type IncompleteGraphRule struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	S1 string `json:"s1"`
	S2 string `json:"s2"`
}

// Reads and unmarshals a sport json into the Sport struct
func ParseSport(filePath string) (*Sport, error) {
	// Read file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal json
	sport := &Sport{}
	err = json.Unmarshal(file, sport)
	if err != nil {
		return nil, err
	}

	return sport, nil
}
