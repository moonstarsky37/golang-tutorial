package main

import "fmt"
// START OMIT
func main() {
	defer fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("done")
}
// END OMIT
