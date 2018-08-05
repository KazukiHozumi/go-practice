package main

import "fmt"

func adder() func(int) (int, int) {
	sum := 0
	count := 0
	return func(n int) (int, int) {
		sum += n
		count++
		return count, sum
	}
}

func main() {
	adder1, adder2 := adder(), adder()
	adder1(5)
	adder1(5)
	c1, s1 := adder1(5)
	adder2(5)
	c2, s2 := adder2(5)

	fmt.Println("c1: ", c1, "s1: ", s1, "c2: ", c2, "s2:", s2)
}
