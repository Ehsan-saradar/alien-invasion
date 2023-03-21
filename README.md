# Invasion
Invasion is a cli app for the alien invasion challenge written in Golang.

---
* [Story](#Story)
* [Usage](#Usage)
* [Test](#Test)
* [Assumption](#Assumption)
---

## Story
This program simulates an alien invasion in the non-existent world of X, using a map containing the names of cities and their connecting roads. Each city has 1-4 directions (north, south, east, or west) representing roads to other cities.

The program takes in a command-line argument specifying the number of aliens to be created. These aliens are randomly placed on the map and wander around, following the roads in any direction. When two aliens end up in the same place, they fight and destroy each other, along with the city they were in. The destroyed city and any roads leading into or out of it are removed from the map.

The program continues running until all aliens have been destroyed or each alien has moved at least 10,000 times. When two aliens fight, a message is printed out indicating the destroyed city and the aliens involved.

After the program finishes, it prints out the remaining map in the same format as the input file.

Please note that the city names do not contain numeric characters, and any additional assumptions are documented in the program's comments or assertions.

## Usage
`invasion [path to map file] -n [aliens-num] -o [output file]`

Invasion takes the path to a map file and the number of aliens to spawn on the map.
and runs the main algorithm of the challenge. and in the end writes the new map to a file.

## Test
`go test ./...`

## Assumption
* Upon their initial landing, the aliens engage in combat.
* The movement of aliens is limited to available directions.
* City names have no spaces or equal signs.
* Only the cardinal directions of "north," "south," "east," and "west" are recognized for road availability in the simulation. Other directions may be inputted, but the road will only function in one direction.
* It is possible for the input file to contain overlapping streets or duplicate cities, in which case the import function will overwrite the previous street.
* The city file may not contain any cities, in which case the simulation runs 10,000 times before completing with an empty output, listing all surviving cities. The program could also be designed to halt when no cities are present, if desired