package main

import "math"

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (a *Coordinates) getDistanceFrom(b *Coordinates) float64 {
	return math.Sqrt((math.Abs((*a).X-(*b).X) + math.Abs((*a).Y-(*b).Y)))
}

func (a *Coordinates) getShiftedPosition(v *Vector) *Coordinates {
	// add vector v to point a to get a final shifted position, as a new point
	return &Coordinates{a.X + v.velocity*math.Cos(v.angle), a.Y + v.velocity*math.Sin(v.angle)}
}
