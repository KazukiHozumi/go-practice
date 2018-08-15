package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 3
		ch <- 2
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
}
