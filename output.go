package main

import "fmt"

type output struct{}

func (o output) HTML(sh shape) string {
	return fmt.Sprintf("<h1>  %s area is %f</h1>\n", sh.GetName(), sh.CalculateArea())
}

func (o output) JSON(sh shape) string {
	return fmt.Sprintf("{ Name: %s, Area: %f}\n", sh.GetName(), sh.CalculateArea())
}
