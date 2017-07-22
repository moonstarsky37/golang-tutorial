package main

import "fmt"

// START OMIT
type Getter func() int
type Setter func(int)

func x_getter_setter(x int) (Getter, Setter) {
        getter := func() int {
                return x
        }
        setter := func(n int) {
                x = n
        }
        return getter, setter
}
func main() {
        getX, setX := x_getter_setter(10)

        fmt.Println(getX())  // 10
        setX(20)
        fmt.Println(getX())  // 20
// END OMIT
}