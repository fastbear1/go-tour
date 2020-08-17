package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (f MyReader) Read(b []byte) (int, error) {
	for i := range b[:1] {
		b[i] = 'A'
	}
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
