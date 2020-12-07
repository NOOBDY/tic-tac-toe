package main

import (
	"fmt"
	"xo/tools"
)

func main() {
	board := tools.InitBoard()
	hasEnded := false

	for round := 1; round <= 9; round++ {
		board = tools.Input(board, round)
		hasEnded = tools.Check(board)

		if hasEnded {
			fmt.Printf("Player %d has won", (round+1)%2+1)
			break
		}
	}
	if !hasEnded {
		fmt.Printf("Tie")
	}

	fmt.Scanln()
}
