package main

import (
	"fmt"
)

type Duck interface {
	quack()
	walk()
}

type Dog struct {
	name    string
	footNum int
}

type Chicken struct {
	name    string
	footNum int
}

func (d Dog) quack() {
	fmt.Println(d.name, "say quack!")
}

func (d Dog) walk() {
	fmt.Println("the creature walk with", d.footNum, "feet")
}

func (c Chicken) quack() {
	fmt.Println(c.name, "say quack!")
}

func (c Chicken) walk() {
	fmt.Println("the creature walk with", c.footNum, "feet")
}

func main() {
	creature1 := Dog{"dog", 4}
	creature2 := Chicken{"chicken", 2}

	creature1.quack()
	creature1.walk()

	creature2.quack()
	creature2.walk()
}
