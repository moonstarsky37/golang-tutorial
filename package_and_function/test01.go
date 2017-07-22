package main

import "fmt"

func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}

func main() {
        maximum := max
        fmt.Println(max(10, 5))     // 10
        fmt.Println(maximum(10, 5)) // 10
}