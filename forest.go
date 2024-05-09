package main

type Forest struct {
	Trees []Tree `json:"trees"` // gotta add burned trees arr, might help with optimization
	maxX  float64
	maxY  float64
}

func (f *Forest) plantTrees(amount uint64, generate Generator) {
	for _, position := range generate(amount, f.maxX, f.maxY) {
		f.Trees = append(f.Trees, Tree{Position: position, IsBurning: false})
	}
}

func (f *Forest) copy() *Forest {
	trees := make([]Tree, len(f.Trees))
	copy(trees, f.Trees)
	return &Forest{
		trees,
		f.maxX,
		f.maxY,
	}
}
