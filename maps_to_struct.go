package main

import "fmt"

type Coord struct {
	Lat, Long float64
}

var m map[string]Coord

func main() {
	m = make(map[string]Coord)
	m["Bell Labs"] = Coord{
		40.68433, -74.39967,
	}
	fmt.Println(m)
}
