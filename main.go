package main

import (
	"encoding/json"
	"os"
)

func main() {
	plantation := Forest{Trees: []Tree{}, MaxX: 10, MaxY: 10}
	plantation.plantTrees(64, generateFullRandom)
	startForestfire(&plantation, &plantation.Trees[0], 1, &Vector{angle: 0, velocity: 0})

	data, err := json.Marshal(plantation.Trees)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("forest.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(data)
}
