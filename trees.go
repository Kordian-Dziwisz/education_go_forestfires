package main

type Tree struct {
	Position  Coordinates `json:"position"`
	IsBurning bool        `json:"isBurning"`
	IsBurned  bool        `json:"isBurned"`
}
