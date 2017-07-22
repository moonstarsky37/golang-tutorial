package main

import "fmt"

type BiFunc func(int, int) int

func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}

func main() {
        var maximum BiFunc
        fmt.Println(maximum)         // nil

        maximum = max
        fmt.Println(maximum(10, 5))  // 10
}