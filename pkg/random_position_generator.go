package forestfire

import "math/rand"

type Generator func(amount uint64, maxX float64, maxY float64) []Coordinates

func generateFullRandom(amount uint64, maxX float64, maxY float64) []Coordinates {
	var result []Coordinates
	for i := 0; i < int(amount); i++ {
		result = append(result, Coordinates{x: maxX * (rand.Float64() - 0.5), y: maxY * (rand.Float64() - 0.5)})
	}
	return result
	// generate an slice of random Coordinates, where max x is maxX and max Y is maxX
}
