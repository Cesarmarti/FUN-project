package sports

import (
	"encoding/json"
	"io/ioutil"
)

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
