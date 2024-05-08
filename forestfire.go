package main

func startForestfire(plantantion *Forest, startingTree *Tree, maxBurnDistance float64, wind *Vector) {
	startingTree.IsBurning = true
	for i := 0; i < len(plantantion.Trees); i++ {
		tree := &plantantion.Trees[i]
		if tree.Position.getDistanceFrom(startingTree.Position.getShiftedPosition(wind)) < maxBurnDistance && !tree.IsBurning {
			startForestfire(plantantion, tree, maxBurnDistance, wind)
		}
	}
}
