package tools

// Check checks if any of the players have won yet
// This function only checks for the currently placed position
func Check(board Board, x, y int) bool {
	var hasEnded bool

	player := board[x][y]

	// check col
	for i := 0; i < 3; i++ {
		if board[x][i] != player {
			break
		}
		if i == 2 {
			hasEnded = true
		}
	}

	// check row
	for i := 0; i < 3; i++ {
		if board[i][y] != player {
			break
		}
		if i == 2 {
			hasEnded = true
		}
	}

	if x == y {
		// check diag
		for i := 0; i < 3; i++ {
			if board[i][i] != player {
				break
			}
			if i == 2 {
				hasEnded = true
			}
		}

		// check anti diag
		for i := 0; i < 3; i++ {
			if board[i][2-i] != player {
				break
			}
			if i == 2 {
				hasEnded = true
			}
		}
	}

	return hasEnded
}
