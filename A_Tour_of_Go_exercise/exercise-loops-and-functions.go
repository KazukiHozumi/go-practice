package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 1.0
	for {
		tmp -= (z*z - x) / (2 * z)
		if math.Abs(z-tmp) > 1e-6 {
			fmt.Println(z)
			z = tmp
		} else {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
