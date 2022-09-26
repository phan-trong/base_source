package main

type square struct {
	length float64
}

func (s square) GetName() string {
	return "square"
}

func (s square) CalculateArea() float64 {
	return s.length * s.length
}
