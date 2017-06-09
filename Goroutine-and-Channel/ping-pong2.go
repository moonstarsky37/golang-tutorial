package main

import(
	"fmt"
	"time"
)

func ping_pong(ch chan time.Time){
	for true{
		ping := <-ch       //receive a ping
		fmt.Println("ping:", time.Since(ping))
		ch <- time.Now()   //return a pong
	}
}
func main(){
	ch := make(chan time.Time)
	go ping_pong(ch)
	go ping_pong(ch)
	ch <- time.Now()
	time.Sleep(time.Second*1)
	fmt.Println("End time:", <-ch)
}
