package main

func SnakesAndLadders(board, dice []int) int {
	gamePosition := 0

	for i, roll := range(dice) {
		if gamePosition + roll >= len(board) { continue }	//overshoot, roll again
		gamePosition += roll	// move based off the roll

		for board[gamePosition] != 0 {
			gamePosition += board[gamePosition]	// move based off snake/ladder
		}

		if (gamePosition == len(board)-1) || (i == (len(dice)-1)) {	// landed on last square or we are out of rolls
			break
		}
	}
	return gamePosition
}