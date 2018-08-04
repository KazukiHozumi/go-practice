package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// add
	primes2 := [7]int{1, 2, 3, 4, 5, 6, 7}
	var s2 []int
	fmt.Println(cap(s2))
	s2 = primes[0:1]
	fmt.Println(cap(s2))
	s2 = primes[0:2]
	fmt.Println(cap(s2))
	s2 = primes[0:3]
	fmt.Println(cap(s2))
	s2 = primes2[:2]
	fmt.Println(cap(s2))
}
