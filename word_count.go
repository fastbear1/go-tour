package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	pos := 0
	lit := make(map[string]int)
	for i, v := range s {
		if string(v) == " " {
			CheckInMap(lit, s[pos:i])
			pos = i + 1
		}
	}
	CheckInMap(lit, s[pos:])
	return lit
}

func CheckInMap(z map[string]int, idx string) {
	_, ok := z[idx]

	if !ok {
		z[idx] = 1
	} else {
		z[idx] += 1
	}
}

func main() {
	wc.Test(WordCount)
}
