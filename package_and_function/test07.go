package main

import "fmt"

type Predicate func(int) bool

// START OMIT
func filter(origin []int, predicate Predicate) []int {
        filtered := []int {}
        for _, elem := range origin {
                if predicate(elem) {
                        filtered = append(filtered, elem)
                }
        }
        return filtered
}
func main() {
        data := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        fmt.Println(filter(data, func(elem int) bool {
                return elem > 5
        }))
        fmt.Println(filter(data, func(elem int) bool {
                return elem <= 6
        })) 
}
// END OMIT