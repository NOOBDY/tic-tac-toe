package tools

import (
	"errors"
	"fmt"
	"strings"
)

//Input validates the input before sending to place()
func Input(board Board, round int) Board {
	fmt.Printf("Player %d's turn\n", (round+1)%2+1)
	render(board)
	for {
		var err error
		var input string
		var player state

		fmt.Print("Enter Position: ")
		fmt.Scanln(&input)

		if round%2 == 1 {
			player = o
		} else {
			player = x
		}

		board, err = place(board, input, player)

		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}
	clearScreen()
	return board
}

// place marks the player's chosen position on the board
func place(board Board, input string, player state) (Board, error) {
	var err error
	if strings.Contains("ABC", strings.ToUpper(string(input[0]))) && strings.Contains("123", string(input[1])) {
		x := int(input[0]) - 97
		y := int(input[1]) - 49

		if board[y][x] == empty {
			board[y][x] = player
		} else {
			err = errors.New("Position Already Taken")
		}
	} else {
		err = errors.New("Invalid Input")
	}
	return board, err
}
