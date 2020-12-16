package tools

import "errors"

// Place marks the player's chosen position on the board
func Place(board *Board, player State, x, y int) error {
	var err error

	if board[x][y] == empty {
		board[x][y] = player
	} else {
		err = errors.New("Position Already Taken")
	}

	return err
}
