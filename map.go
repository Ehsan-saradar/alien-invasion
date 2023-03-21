package alieninvasion

import (
	"bytes"
	"fmt"
)

type Map struct {
	cities       []*City
	citiesByName map[string]*City
}

// String returns a string representation of the map.
func (m Map) String() string {
	s, _ := m.MarshalText()
	return string(s)
}

// MarshalText implements the encoding.TextMarshaler interface.
// It returns a text representation of the map where each line is a city marshaled as text.
func (m Map) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, city := range m.cities {
		b, err := city.MarshalText()
		if err != nil {
			return nil, fmt.Errorf("marshal city %q: %w", city.name, err) // This should never happen!
		}
		buf.Write(b)

		// Add a new line if this is not the last city.
		if i < len(m.cities)-1 {
			buf.WriteByte('\n')
		}
	}

	return buf.Bytes(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// It parses the text representation of the map and creates the cities.
func (m *Map) UnmarshalText(text []byte) error {
	// Split the text into lines.
	lines := bytes.Split(text, []byte{'\n'})

	// Reset the map.
	m.cities = make([]*City, 0, len(lines))
	m.citiesByName = make(map[string]*City, len(lines))

	for i, line := range lines {
		// Ignore empty lines.
		if len(line) == 0 {
			continue
		}

		// Create a new city and unmarshal the line into it.
		city := &City{}
		if err := city.UnmarshalText(line); err != nil {
			return fmt.Errorf("failed to unmarshal city at line %d: %w", i, err)
		}

		if _, ok := m.citiesByName[city.name]; ok {
			return fmt.Errorf("duplicate city name %q", city.name)
		}

		// Add the city to the map.
		m.cities = append(m.cities, city)
		m.citiesByName[city.name] = city
	}

	// Connect the cities.
	for _, city := range m.cities {
		if city.north != nil {
			city.north = m.citiesByName[city.north.name]
		}
		if city.south != nil {
			city.south = m.citiesByName[city.south.name]
		}
		if city.east != nil {
			city.east = m.citiesByName[city.east.name]
		}
		if city.west != nil {
			city.west = m.citiesByName[city.west.name]
		}
	}

	return nil
}

// Len returns the number of cities in the map.
func (m Map) Len() int {
	return len(m.cities)
}

// CityByName returns the city with the given name.
func (m Map) CityByName(name string) *City {
	return m.citiesByName[name]
}

// CityByIndex returns the city at the given index.
func (m Map) CityByIndex(index int) *City {
	return m.cities[index]
}

// RemoveCity removes the city with the given name from the map.
func (m *Map) RemoveCity(name string) {
	for i, city := range m.cities {
		if city.name == name {
			// Remove the city from the cities slice.
			m.cities = append(m.cities[:i], m.cities[i+1:]...)

			// Remove the city from the its neighbors.
			if city.north != nil {
				city.north.south = nil
			}
			if city.south != nil {
				city.south.north = nil
			}
			if city.east != nil {
				city.east.west = nil
			}
			if city.west != nil {
				city.west.east = nil
			}

			return
		}
	}

	// Remove the city from the citiesByName map.
	delete(m.citiesByName, name)
}
