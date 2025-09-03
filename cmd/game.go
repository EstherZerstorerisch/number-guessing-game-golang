package cmd

import (
	"fmt"
	"math/rand/v2"
)

func RandomNum() int {
	return rand.IntN(101)
}

func DiffLevel() (difficulty string) {
	fmt.Println(`
	Please select the difficulty level:
	1. Easy (10 chances)
	2. Medium (5 chances)
	3. Hard (3 chances)
	`)
	// Prompt use to input number
	var choice int
	fmt.Scanln(&choice)

	// Use switch statement
	switch choice {
	case 1:
		difficulty = "Easy"
	case 2:
		difficulty = "Medium"
	case 3:
		difficulty = "Hard"
	default:
		fmt.Println("Wrong! Please chose again!")
		DiffLevel()
	}
	fmt.Printf("Great! You have selected the %s difficulty level", difficulty)
	return 
}

