package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 7, 11, 12}

	var s []int = primes[0:1]
	fmt.Println(s)
}
