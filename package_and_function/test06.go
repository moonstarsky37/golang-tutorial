package main

import "fmt"

type Predicate func(int) bool

func filter(origin []int, predicate Predicate) []int {
        filtered := []int {}
        for _, elem := range origin {
                if predicate(elem) {
                        filtered = append(filtered, elem)
                }
        }
        return filtered
}
// START OMIT
func _filter(origin []int, predicate Predicate) []int {
        filtered := []int {}
        // code

        // code
        return filtered
}

func greatThan5(elem int) bool {
        return elem > 5
}

func lessEqualsThan6(elem int) bool {
        return elem <= 6
}

func main() {
        data := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        fmt.Println(filter(data, greatThan5))       // [6 7 8 9 10]
        fmt.Println(filter(data, lessEqualsThan6))  // [1 2 3 4 5 6]
}
// END OMIT