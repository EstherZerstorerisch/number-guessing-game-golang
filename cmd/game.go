package cmd

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func RandomNum() uint8 {
	return uint8(rand.IntN(101))
}

func DiffLevel() (chances uint8) {
	var choice uint8
	var difficulty string
	
	// Prompt user to input number
	fmt.Print("Please select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\n")
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatalln("Invalid input!")
	}
	fmt.Println()
	// Use switch statement
	switch choice {
	case 1:
		difficulty = "Easy"
		chances = 10
	case 2:
		difficulty = "Medium"
		chances = 5
	case 3:
		difficulty = "Hard"
		chances = 3
	default:
		log.Fatal("Wrong! Please chose again!\n\n")
		DiffLevel()
	}
	if difficulty != "" {
		fmt.Printf("Great! You have selected the %s difficulty level\n", difficulty)
	}
	return 
}

func GuessNumber() (guess uint8) {
	fmt.Print("Enter your guess: ")
	_, err := fmt.Scanln(&guess)
	if err != nil {
		log.Fatal("Invalid input")
	}
	return
}

func CompareAnswer(userInput, answer uint8, chances *uint8) bool {
	if userInput == answer {
		fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", *chances)
		return true
	} else if userInput > answer {
		fmt.Printf("Incorrect! The number is less than %d.\n", userInput)
		*chances --
		return false
	} else {
		fmt.Printf("Incorrect! The number is greater than %d.\n", userInput)
		*chances --
		return false
	}
}
