package forestfire

import "math"

type Coordinates struct {
	x float64
	y float64
}

func (a *Coordinates) getDistanceFrom(b *Coordinates) float64 {
	return math.Sqrt((math.Abs((*a).x-(*b).x) + math.Abs((*a).y-(*b).y)))
}

func (a *Coordinates) getShiftedPosition(v *Vector) *Coordinates {
	// add vector v to point a to get a final shifted position, as a new point
	return &Coordinates{a.x + v.velocity*math.Cos(v.angle), a.y + v.velocity*math.Sin(v.angle)}
}
