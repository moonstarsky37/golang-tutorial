package main

import (
	"fmt"
	"time"
	"math/rand"
)

func robot(){
	rand.Seed(time.Now().Unix())
	for true{
		computerChoice := rand.Intn(3)

		/* return rock, scissor, or paper in int */

	}
}

func main(){
	optionName := []string{
		"Rock",
		"Scissor",
		"Paper",
	}
	for true{
		var playerChoice, computerChoice int;
		fmt.Printf("(0:Rock, 1:Scissor, 2:Paper):")
		fmt.Scanln(&playerChoice)

		/* receive computer choice */

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
