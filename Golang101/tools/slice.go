package main

import "fmt"
// START OMIT
func main() {
	names := []string{
			"John",
			"Paul",
			"George",
			"Ringo",
		}
		fmt.Println(names)

		a := names[0:2]
		b := names[1:3]
		fmt.Println(a, b)

		b[0] = "XXX" //注意這個改變！！
		fmt.Println(a, b)
		fmt.Println(names)
}
// END OMIT
