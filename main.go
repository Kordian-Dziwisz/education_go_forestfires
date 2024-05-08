package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	plantation := Forest{Trees: []Tree{}, MaxX: 10000, MaxY: 10000}
	plantation.plantTrees(64e+5, generateFullRandom)
	startForestfire(&plantation, &plantation.Trees[0], 1, &Vector{angle: 0, norm: 0})

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

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
