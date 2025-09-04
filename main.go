package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/EstherZerstorerisch/number-guessing-game-golang/cmd"
)

func main() {
	// Game introduction
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Print("I'm thinking of a number between 1 and 100.\n\n")
	
	restart := true
	for restart {
		// Create a random number
		chosenNum := cmd.RandomNum()
		fmt.Println(chosenNum)
		// Let user choose difficulty level then return remaining chances
		guessingChances := cmd.DiffLevel()
		fmt.Print("Let's start the game!\n\n")

		// User guessing number
		userAnswer := cmd.GuessNumber()
		
		// Compare to the result
		for !cmd.CompareAnswer(userAnswer, chosenNum, &guessingChances) {
			userAnswer = cmd.GuessNumber()
			if guessingChances == 0 {
				fmt.Print("You lose!\n\n")
			}
		}
		// Ask user if want to replay
		var play string
		fmt.Print("Do you want to start new game? (y/n) ")
		_, err := fmt.Scanln(&play)
		fmt.Println()
		if err != nil {
			log.Fatal("Invalid input.")
		}

		switch strings.ToLower(play) {
		case "yes": restart = true
		case "y": restart = true
		case "no": restart = false
		case "n": restart = false
		}
	}
}
