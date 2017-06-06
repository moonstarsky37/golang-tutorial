package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pi(5000))
  sum := func(a, b int) int { return a+b } (3, 4)
  fmt.Println(sum)
}
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}
func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
