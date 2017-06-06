package main

import "fmt"
// START OMIT
func 交換(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := 交換("梓", "育")
	fmt.Println(a, b)
}
// END OMIT
