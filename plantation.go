package main

type Plantantion struct {
	trees []Tree
	maxX  float64
	maxY  float64
}

func (p Plantantion) plantTrees(amount uint64, generate Generator) {
	for _, position := range generate(amount, p.maxX, p.maxY) {
		p.trees = append(p.trees, Tree{position: position, isBurning: false})
	}
}
