package main

import "fmt"

type Consumer func(int)

// START OMIT
func forEach(elems []int, consumer Consumer) {
        for _, elem := range elems {
                consumer(elem)
        }
}
func main() {
        numbers := []int {1, 2, 3, 4, 5}
        sum := 0
        forEach(numbers, func(elem int) {
                sum += elem
        })
        fmt.Println(sum)  // 15
// END OMIT
}