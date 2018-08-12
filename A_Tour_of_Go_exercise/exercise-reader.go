package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: add a Read([]byte) (int, error) method to MyReder.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})

	// add
	var r MyReader
	b := make([]byte, 8)
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
}
