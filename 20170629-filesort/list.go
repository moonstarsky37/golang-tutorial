package main

import ("fmt" ; "container/list")

func main(){
var x list.List
x.PushBack(9)
x.PushBack(4)
x.PushBack(8)
x.PushBack(7)

for e := x.Front(); e != nil; e=e.Next(){
fmt.Println(e.Value.(int))

}
}
