package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 3
}
