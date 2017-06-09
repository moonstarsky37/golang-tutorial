package main

import (
	"fmt"
)

func test(arg string) {
	fmt.Println(arg)
}

func main() {
	ch := make(chan func())
	ch <- test
	ch <- func(arg int) { fmt.Println(arg) }
}
