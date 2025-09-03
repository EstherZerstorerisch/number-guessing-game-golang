package main

import (
	"fmt"

	"github.com/EstherZerstorerisch/number-guessing-game-golang/cmd"
)

func main() {
	chosenNum := cmd.RandomNum()
	fmt.Printf("My number: %d", chosenNum)
}
