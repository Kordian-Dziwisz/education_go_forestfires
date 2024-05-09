package main

type FFArgs struct {
	maxBurnDistance float64
	wind            *Vector
}

type History struct {
	args      FFArgs
	Snapshots []Forest `json:"snapshots"`
}

func (h *History) startForestfire(treeIndex int, forest *Forest) {
	h.Snapshots = append(h.Snapshots, *forest.copy())
	h.Snapshots[0].Trees[treeIndex].IsBurning = true
	h.step(1)
}

func (h *History) step(t int) {
	h.Snapshots = append(h.Snapshots, *h.Snapshots[t-1].copy())
	burnFurther := false
	for i, tree := range h.Snapshots[t].Trees {
		if tree.IsBurning && h.Snapshots[t-1].Trees[i].IsBurning {
			h.Snapshots[t].Trees[i].IsBurning = false
			h.Snapshots[t].Trees[i].IsBurned = true
			for j, anotherTree := range h.Snapshots[t].Trees {
				if h.args.maxBurnDistance > anotherTree.Position.getDistanceFrom(tree.Position.getShiftedPosition(h.args.wind)) &&
					!anotherTree.IsBurned &&
					!anotherTree.IsBurning {
					burnFurther = true
					h.Snapshots[t].Trees[j].IsBurning = true
				}
			}
		}
	}
	if burnFurther {
		h.step(t + 1)
	}
}
