package main

import (
	"fmt"
)

func main() {
	var x1, y1 int = 3, 4
	var x2, y2 float64 = float64(x1), float64(y1)
	fmt.Println(x1/y1)
	fmt.Println(x2/y2)
}
