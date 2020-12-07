package tools

import (
	"fmt"
	"runtime"
)

func clearScreen() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// render draws out the board
func render(board Board) {
	fmt.Println("  A B C")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", i+1)

		for j := 0; j < 3; j++ {
			fmt.Printf(" %s", board[i][j])
		}

		fmt.Print("\n")
	}
}
