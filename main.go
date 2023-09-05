package main

import (
	"fmt"
)

// TODO: implement Knight moves
func knightMoves() {

}

// TODO: implement function to find all possible moves

// TODO: implement function to detect and determin best moves with lowest output to other cells

func main() {

	var board [8][8]int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {

			// Calculate i*j and assign to board
			board[i][j] = (i + 1) * (j + 1)
		}
	}

	for i := 0; i < 8; i++ {

		for j := 0; j < 8; j++ {

			// Print cell with width of 3, right aligned
			fmt.Printf(" %3d", board[i][j])
		}
		fmt.Println("")
	}
}
