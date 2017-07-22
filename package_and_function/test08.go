package main

import "fmt"

type Func1 func(int) int

func funcA() Func1 {
        x := 10
        return func(n int) int {
                return x + n
        }
}

func main() {
        fmt.Println(funcA()(2))  // 12
}