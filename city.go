package alieninvasion

import (
	"fmt"
	"strings"
)

// City is a struct that represents a city and its neighbors in four direction of the map (north, south, east and west).
// nill value means that there is no neighbor in that direction (e.g. North is nil for the northernmost city).
type City struct {
	name                     string
	north, south, east, west *City
}

// MarshalText implements the encoding.TextMarshaler interface.
// It marshals the city as one line text in the following format:
// <name> north=<north> south=<south> east=<east> west=<west>
// Where <name>, <north>, <south>, <east> and <west> are the names of the city and its neighbors.
// If a neighbor is nil, it will be ignored.
func (c City) MarshalText() ([]byte, error) {
	t := c.name
	if c.north != nil {
		t += " north=" + c.north.name
	}
	if c.south != nil {
		t += " south=" + c.south.name
	}
	if c.east != nil {
		t += " east=" + c.east.name
	}
	if c.west != nil {
		t += " west=" + c.west.name
	}
	return []byte(t), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// For more information about the format, see the MarshalText method.
func (c *City) UnmarshalText(text []byte) error {
	fields := strings.Fields(string(text))
	if len(fields) < 1 {
		return fmt.Errorf("empty string is not a valid city")
	}

	c.name = fields[0]
	for _, f := range fields[1:] {
		parts := strings.Split(f, "=")
		if len(parts) != 2 {
			return fmt.Errorf("city neighbor format is invalid, expected <direction>=<name> but got %s", f)
		}
		switch parts[0] {
		case "north":
			c.north = &City{name: parts[1]}
		case "south":
			c.south = &City{name: parts[1]}
		case "east":
			c.east = &City{name: parts[1]}
		case "west":
			c.west = &City{name: parts[1]}
		default:
			return fmt.Errorf("unknown direction %s", parts[0])
		}
	}

	return nil
}

// Neighbors returns a slice of the city's neighbors.
func (c City) Neighbors() []*City {
	neighbors := make([]*City, 0, 4)
	if c.north != nil {
		neighbors = append(neighbors, c.north)
	}
	if c.south != nil {
		neighbors = append(neighbors, c.south)
	}
	if c.east != nil {
		neighbors = append(neighbors, c.east)
	}
	if c.west != nil {
		neighbors = append(neighbors, c.west)
	}
	return neighbors
}
