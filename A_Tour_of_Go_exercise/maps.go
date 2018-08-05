package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// add
	m2 := make(map[string]int)
	m2["aaa"] = 1
	fmt.Println(m2["aaa"])

	m3 := make(map[string]int)
	m3["bbb"] = 2
	fmt.Println(m3["bbb"])
}
