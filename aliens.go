package alieninvasion

import (
	"fmt"
	"strings"
)

// Aliens represents a list of aliens.
type Aliens []Alien

// Empty returns true if the list of aliens is empty.
func (a Aliens) Empty() bool {
	return len(a) == 0
}

// RemoveByCity removes all aliens that have the same location as the given city from the list and returns a new list.
func (a Aliens) RemoveByCity(c *City) Aliens {
	var newAliens Aliens
	for _, alien := range a {
		if alien.location != c {
			newAliens = append(newAliens, alien)
		}
	}
	return newAliens
}

func (a Aliens) String() string {
	var sb strings.Builder
	for i, alien := range a {
		sb.WriteString(fmt.Sprintf("(%d, '%s')", alien.id, alien.location.name))

		if i < len(a)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// Alien represents an alien on the map.
type Alien struct {
	id       int
	location *City
}

// Move moves the alien to another location.
func (a *Alien) Move(loc *City) {
	a.location = loc
}
