package tools

import (
	"fmt"
	"runtime"
)

// ClearScreen clears the screen using commands depending on the system
func ClearScreen() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// Render draws out the board
func Render(board *Board) {
	fmt.Println("  1 2 3")
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", rune(i+65))

		for j := 0; j < 3; j++ {
			fmt.Printf(" %s", board[i][j].String())
		}

		fmt.Print("\n")
	}
}
