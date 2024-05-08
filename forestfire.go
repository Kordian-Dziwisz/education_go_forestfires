package main

func startForestfire(plantantion *Plantantion, startingTree *Tree, maxBurnDistance float64, wind *Vector) {
	startingTree.isBurning = true
	for _, tree := range plantantion.trees {
		if tree.position.getDistanceFrom(startingTree.position.getShiftedPosition(wind)) < maxBurnDistance && !tree.isBurning {
			startForestfire(plantantion, &tree, maxBurnDistance, wind)
		}
	}
}
