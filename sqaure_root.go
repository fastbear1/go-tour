package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  for i:=0;i<10;i++ {
    last_z := z
    z -= (z*z - x)/(2*z)
    fmt.Println(z)
    if last_z == z {
       break
    }
  }
  return z
}

func main() {
  fmt.Println(Sqrt(1))
}
