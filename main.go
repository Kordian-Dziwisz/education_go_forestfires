package main

import (
	"encoding/json"
	"os"
)

func main() {
	// start := time.Now()
	forest := Forest{Trees: []Tree{}, maxX: 10, maxY: 10}
	forest.plantTrees(192, generateFullRandom)

	history := History{Snapshots: []Forest{},
		args: FFArgs{maxBurnDistance: 1,
			wind: &Vector{angle: 0, norm: 1}}}

	history.startForestfire(0, &forest)

	data, err := json.Marshal(history)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("plotter/forest.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	// elapsed := time.Since(start)
	// fmt.Println(elapsed)
}
