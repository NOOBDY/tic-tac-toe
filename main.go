package main

import (
	"fmt"
	"xo/tools"
)

func main() {
	board := tools.InitBoard()

	for round := 1; round <= 9; round++ {
		board = tools.Input(board, round)
	}

	fmt.Scanln()
}
