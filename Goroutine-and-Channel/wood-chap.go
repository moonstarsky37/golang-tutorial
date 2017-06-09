package main
import (
	"fmt"
	"time"
)
func chap(c chan string, s string){
	i := 0;
	for i<len(s){
		if(s[i] == ' '){
			c <- s[:i+1]
		}
		i++
	}
	c <- s[:i]
	close(c)
}
func main(){
	c := make(chan string)
	go chap(c, "How much wood would a woodchuck chuck if a woodchuck could chuck wood?")
	var v, ok = <- c
	for ok{
		fmt.Println(v);
		time.Sleep(time.Millisecond*200)
		v, ok = <- c
	}
}
