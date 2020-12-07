package tools

import (
	"os"
	"os/exec"
)

var clear map[string]func() // create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//InitBoard initializes the board, also calls render() for one time
func InitBoard() [3][3]string {
	var board [3][3]string

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "."
		}
	}

	return board
}
