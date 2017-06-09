package main

import (
	"fmt"
	"math/rand"
	"time"
)

func a(x int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine no.", x, ";", i)
		j := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * j)
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		go a(i)
	}
	var a string
	fmt.Scanln(&a)
}
