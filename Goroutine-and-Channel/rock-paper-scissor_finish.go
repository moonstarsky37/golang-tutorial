package main

import (
	"fmt"
	"time"
	"math/rand"
)

func robot(ch chan int){
	rand.Seed(time.Now().Unix())
	for true{
		ch <- rand.Intn(3)
	}
}

func main(){
	optionName := []string{
		"Rock",
		"Scissor",
		"Paper",
	}
	ch := make(chan int)
	go robot(ch)
	for true{
		var playerChoice int;
		fmt.Printf("(0:Rock, 1:Scissor, 2:Paper):")
		fmt.Scanln(&playerChoice);
		computerChoice := <- ch

		fmt.Println("Computer:", optionName[computerChoice])
		fmt.Println("Player:", optionName[playerChoice])
		//compare who's win
		switch (playerChoice - computerChoice + 3) % 3 {
		case 0:
			fmt.Println("draw!")
		case 1:
			fmt.Println("computer win!")
		case 2:
			fmt.Println("player win!")
		}
	}
}
