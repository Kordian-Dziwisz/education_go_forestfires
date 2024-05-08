package main

type Tree struct {
	position  Coordinates `json: "position"`
	isBurning bool        `json: "isBurning"`
}
