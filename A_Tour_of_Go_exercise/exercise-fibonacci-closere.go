package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1, f2 := 0, 0
	return func() int {
		f3 := f1 + f2
		f2 = f1
		f1 = f3
		if f1 == 0 {
			f2 = 1
		}
		return f1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
