package main

import (
	"fmt"
	"math"
)
// if statement starting with a short statement to execute before 
//the condition

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// cant use v here, though
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}