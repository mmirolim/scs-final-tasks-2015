// create simple console game
// find number computer generated randomly from 1 to 100
package main

import (
	"fmt"
	"math/rand"
	"scs-final-tasks/console"
	"strconv"
)

func main() {
	fmt.Println("Hello Let's Play")
	fmt.Println("Guess what number I think of")
	input := make(chan string)
	console.Start(input)

	var number int
	// make a rand number
	number = rand.Intn(100)
	var answer string
	// keep listening for inputs from console
	for {
		select {
		case in := <-input:
			guess, err := strconv.Atoi(in)
			if err != nil {
				answer = "it should be a number from 1 to 100"
				break
			}
			switch {
			case guess > number:
				answer = "less"
			case guess < number:
				answer = "more"
			case guess == number:
				answer = "Yep you are right"
			}
			fmt.Println(answer)
		}
	}
}
