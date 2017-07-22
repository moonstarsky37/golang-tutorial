package main

import "fmt"
import "reflect"

func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}

func main() {
        maximum := max
        fmt.Println(reflect.TypeOf(max))     // func(int, int) int
        fmt.Println(reflect.TypeOf(maximum)) // func(int, int) int
}