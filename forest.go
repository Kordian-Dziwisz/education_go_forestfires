package main

type Forest struct {
	Trees []Tree `json:"trees"`
	MaxX  float64
	MaxY  float64
}

func (p *Forest) plantTrees(amount uint64, generate Generator) {
	for _, position := range generate(amount, p.MaxX, p.MaxY) {
		p.Trees = append(p.Trees, Tree{Position: position, IsBurning: false})
	}
}
