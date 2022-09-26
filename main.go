package main

import (
	"fmt"
)

func main() {
	c := circle{
		radius: 3,
	}

	s := square{
		length: 3,
	}

	o := output{}

	fmt.Println(o.HTML(c))
	fmt.Println(o.HTML(s))
	fmt.Println(o.JSON(c))
	fmt.Println(o.JSON(s))
}
