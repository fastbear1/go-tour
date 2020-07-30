package main

import (
	"fmt"
)

func fibonachi() func() int {
	i := 0
	var last, pos int
	return func() int {
		pos = i
		if last == 0 && i == 0 {
			last = 0
			i = 1
		} else if last == 0 && i == 1 {
			last = 1
			i = 1
		} else {
			i = i + last
		}
		last = pos

		fmt.Println(last, i)
		return last
	}
}

func main() {
	fib := fibonachi()
	for i := 0; i < 10; i++ {
		fib()
		//fmt.Println(i, fib())
	}
}
