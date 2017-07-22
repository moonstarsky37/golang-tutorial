package main

import "fmt"

func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}

func main() {
        var maximum func(int, int) int
        fmt.Println(maximum)           // nil

        maximum = max
        fmt.Println(maximum(10, 5))    // 10
}