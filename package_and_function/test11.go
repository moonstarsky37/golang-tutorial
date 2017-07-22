package main

import "fmt"

type Getter func() int
type Setter func(int)

// START OMIT
func x_getter_setter(x int) (Getter, Setter) {
        fmt.Printf("the parameter :\tx (%p) = %d\n", &x, x)

        getter := func() int {
                fmt.Printf("getter invoked:\tx (%p) = %d\n", &x, x)
                return x
        }
        setter := func(n int) {
                x = n
                fmt.Printf("setter invoked:\tx (%p) = %d\n", &x, x)
        }
        return getter, setter
}

func main() {
        getX, setX := x_getter_setter(10)

        fmt.Println(getX())
        setX(20)
        fmt.Println(getX())
}
// END OMIT