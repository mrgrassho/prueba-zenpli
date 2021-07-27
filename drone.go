package main

import (
	"fmt"
	"errors"
	"math"
)

var CardinalPoints = []string{"N", "E", "S", "W"}

var CardinalNames = map[string]string {
	"N": "Norte",
	"E": "Oriente",
	"S": "Sur",
	"W": "Occidente",
}

type Drone struct {
	X, Y int
	Dir string
	Limit int
}

func (d *Drone) Rotate(i int) {
	currentDir := IndexOf(d.Dir, CardinalPoints)
	nextDir := currentDir + i
	if nextDir < 0 {
		nextDir = len(CardinalPoints) - 1
	} else if nextDir >= len(CardinalPoints) {
		nextDir = 0
	}
	d.Dir = CardinalPoints[nextDir]
}

func (d *Drone) RotateLeft() {
	d.Rotate(-1)
}

func (d *Drone) RotateRight() {
	d.Rotate(1)
}

func (d *Drone) GoFoward(i int) error {
	switch c := d.Dir; c {
	case "N": // North
		d.Y += i 
	case "E": // East
		d.X += i
	case "S": // South
		d.Y -= i
	case "W": // West
		d.X -= i
	}
	if int(math.Abs(float64(d.Y))) > d.Limit || int(math.Abs(float64(d.X))) > d.Limit {
		return errors.New(fmt.Sprintf("Limit %d exceeded", d.Limit))
	}
	return nil
}

func (d *Drone) GoFowardSingle() error {
	return d.GoFoward(1)
}

func (d *Drone) Report() string {
	return fmt.Sprintf("(%d, %d) dirección %s\n", d.X, d.Y, CardinalNames[d.Dir])
}