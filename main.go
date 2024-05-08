package main

import (
	"encoding/json"
	"os"
)

func main() {
	// start := time.Now()
	plantation := Forest{Trees: []Tree{}, MaxX: 10000, MaxY: 10000}
	plantation.plantTrees(64e+5, generateFullRandom)
	startForestfire(&plantation, &plantation.Trees[0], 1, &Vector{angle: 0, norm: 0})

	data, err := json.Marshal(plantation.Trees)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("forest.json")
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
