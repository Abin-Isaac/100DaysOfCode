package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.8, 543.8 //math>Min only takes float argument 
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
}
