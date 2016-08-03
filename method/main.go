// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p *player) average() float64 {
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.

	players := []player{
		{name: "Justin", atBats: 27, hits: 14},
		{name: "Bill", atBats: 32, hits: 26},
		{name: "Kyle", atBats: 24, hits: 10},
	}

	// Display the batting average for each player in the slice.
	for _, person := range players {
		fmt.Println(person.name, person.average())
	}
}
