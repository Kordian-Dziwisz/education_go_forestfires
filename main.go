package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	plantation := Plantantion{trees: []Tree{}}
	plantation.plantTrees(64, generateFullRandom)
	if data, err := json.Marshal(plantation.trees); err != nil {
		fmt.Println("data:")
		fmt.Print(data)
	}
}
