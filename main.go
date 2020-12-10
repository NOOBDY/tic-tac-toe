package main

import (
	"fmt"
	"xo/tools"
)

func main() {
	board := tools.InitBoard()
	hasEnded := false

	for round := 1; round <= 9; round++ {
		var player = tools.State((round+1)%2 + 1)
		var x, y int

		fmt.Printf("Player %d's turn\n", player)
		tools.Render(board)

		for {
			var err error
			x, y = tools.Input()
			board, err = tools.Place(board, player, x, y)

			if err == nil {
				break
			} else {
				fmt.Println(err)
			}
		}

		tools.ClearScreen()

		hasEnded = tools.Check(board, x, y)

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
