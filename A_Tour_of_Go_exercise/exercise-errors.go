package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
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
		return z, nil
	} else {
		return 2, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(-2))
}
