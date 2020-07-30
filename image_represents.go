package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	board := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		board[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			board[x][y] = uint8((x+y)/2 + (x * y))
		}
	}
	return board
}

func main() {
	Pic(10, 10)
	pic.Show(Pic)
}
