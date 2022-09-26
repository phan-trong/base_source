package main

import "math"

type circle struct {
	radius float64
}

func (c circle) GetName() string {
	return "circle"
}

func (c circle) CalculateArea() float64 {
	return c.radius * c.radius * math.Pi
}
