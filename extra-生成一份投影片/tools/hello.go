package main

import (
	"fmt"
)

func main() {
	fmt.Println("33 plus 22 is", add(33, 22))
}

func add(a int, b int) int {
	return a + b
}
