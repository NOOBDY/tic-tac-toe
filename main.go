package main

import (
	"fmt"
	"xo/tools"
)

func main() {
	board := tools.InitBoard()

	for round := 1; round <= 9; round++ {
		board = tools.CheckInput(board, round)
	}

	fmt.Scanln()
}
