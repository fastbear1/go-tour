package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var text string
	for i, v := range ip {
		if i == 0 {
			text = fmt.Sprintf("%v", v)
		} else {
			text = fmt.Sprintf("%v.%v", text, v)
		}
	}
	return fmt.Sprintf("%v", text)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
