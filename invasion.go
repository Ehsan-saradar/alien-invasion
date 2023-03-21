package alieninvasion

import (
	"fmt"
	"math/rand"
)

const MaxInvadeIterations = 10000

type Invasion struct {
	// Map is the map of cities that aliens will invade.
	Map *Map
	// Aliens is the list of aliens that will invade the map.
	Aliens Aliens

	// The random generator used to generate random numbers.
	rnd *rand.Rand
}

// NewInvasion creates a new invasion with the given map and number of aliens.
func NewInvasion(m *Map, n int, rnd *rand.Rand) *Invasion {
	return &Invasion{
		Map:    m,
		Aliens: spawnAliens(m, n, rnd),
		rnd:    rnd,
	}
}

// spawnAliens spawns n aliens on the map and returns a list of all aliens spawned.
func spawnAliens(m *Map, n int, rnd *rand.Rand) Aliens {
	var aliens Aliens
	for i := 0; i < n; i++ {
		aliens = append(aliens, Alien{
			id:       i,
			location: m.CityByIndex(rnd.Intn(m.Len())),
		})
	}
	return aliens
}

// Run runs the main algorithm of this challenge.
// It takes a map of cities and a number determining the initial number of spawned aliens on the map.
// then it randomly spawns aliens on the map and moves them around in each iteration. If more than one alien
// meet each other in a city, they will fight and the city will be destroyed. The algorithm stops when all
// aliens are destroyed or each of them moved at least 10,000 times.
// Note that the algorithm is based on math/rand package so to make the result truly random make sure to seed the random generator.
func (inv *Invasion) Run() {
	// Destroy all cities with more than one alien once before the first iteration.
	inv.destroy()

	for i := 0; i < MaxInvadeIterations; i++ {
		if inv.Aliens.Empty() {
			break
		}

		// Move all aliens in this iteration.
		inv.moveAliens()

		// Destroy all cities and aliens when more than one aliens are in the same city.
		inv.destroy()
	}
}

// moveAliens moves all aliens in the list to a random neighbor city.
func (inv *Invasion) moveAliens() {
	for i := 0; i < len(inv.Aliens); i++ {
		neighborCities := inv.Aliens[i].location.Neighbors()
		// If the alien is in a city with no neighbors, skip it.
		if len(neighborCities) == 0 {
			continue
		}

		inv.Aliens[i].Move(neighborCities[inv.rnd.Intn(len(neighborCities))])
	}
}

// destroy destroys all aliens that have same location from the list along with their cities.
func (inv *Invasion) destroy() {
	// Create a map of cities and all the aliens that are in the same city.
	cities := make(map[*City][]int)
	for i := range inv.Aliens {
		cities[inv.Aliens[i].location] = append(cities[inv.Aliens[i].location], inv.Aliens[i].id)
	}

	// Destroy all cities that have more than one alien in it.
	for city, aliens := range cities {
		if len(aliens) > 1 {
			// Remove all aliens that have the same location as a destroyed city.
			inv.Aliens = inv.Aliens.RemoveByCity(city)

			// Remove the destroyed city from the map.
			inv.Map.RemoveCity(city.name)

			fmt.Printf("%s has been destroyed by aliens %v!\n", city.name, aliens)
		}
	}
}
