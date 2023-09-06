package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Chess board size
const size = 8

// Knight move offsets
var offsets = [][]int{
	{-2, 1}, {-2, -1},
	{-1, 2}, {-1, -2},
	{1, 2}, {1, -2},
	{2, 1}, {2, -1},
}

// Check if move is valid on board
func validMove(x, y int) bool {

	return x >= 0 && x < size && y >= 0 && y < size

}

// Get all valid knight moves from position (x,y)
func getKnightMoves(x, y int) ([][]int, int) {

	var moves [][]int
	count := 0

	for _, offset := range offsets {
		x1 := x + offset[0]
		y1 := y + offset[1]
		if validMove(x1, y1) {
			moves = append(moves, []int{x1, y1})
			count++
		}
	}

	return moves, count
}

func getBestMove(x, y int) []int {

	// Get initial moves
	moves, _ := getKnightMoves(x, y)

	// Track best move and smallest count
	var bestMove []int
	minCount := size * size

	// Check each initial move
	for _, m := range moves {

		// Get subsequent moves
		movesCount, _ := getKnightMoves(m[0], m[1])

		// Update if smaller count
		if len(movesCount) < minCount {
			bestMove = m
			minCount = len(movesCount)
		}
	}

	return bestMove
}

func ReadInput() (x int, y int) {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter X and Y Coordinates (comma-separated): ")
		input, _ := reader.ReadString('\n')

		// Remove leading/trailing whitespace and split input by comma
		coordinates := strings.Split(strings.TrimSpace(input), ",")

		if len(coordinates) != 2 {
			fmt.Println("Please enter two coordinates separated by a comma.")
			continue // Loop back to the beginning
		}

		x := strings.TrimSpace(coordinates[0])
		y := strings.TrimSpace(coordinates[1])

		fmt.Printf("X Coordinate: %s\n", x)
		fmt.Printf("Y Coordinate: %s\n", y)

		// If you want to exit the loop after successfully getting valid input, you can break here.
		break
	}
	// Returns x and y coordinates as integers
	return

}
func main() {

	var board [8][8]int
	//var moves [][]int
	var x, y = ReadInput()
	var nextMove []int
	count := 1
	board[x][y] = count
	getKnightMoves(x, y)
	fmt.Println(nextMove)
	moves1, movesCount := getKnightMoves(x, y)
	fmt.Println(moves1)
	fmt.Println("Moves Count: ", movesCount)
	//fmt.Println(getBestMove(x, y))
	//fmt.Println(getBestMove(x, y))
	// var counter int = 1
	// var x, y string
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)

	// // Ininite coordinates
	// for k := 2; k < 65; k++ {

	// }

	// x, y := 0, 0

	// for count <= 64 {

	// 	// Get best move
	// 	move := getBestMove(x, y)
	// 	fmt.Println(move)

	// 	// Mark move on board
	// 	board[move[0]][move[1]] = count

	// 	// Update current position
	// 	x, y = move[0], move[1]

	// 	count++
	// }

	// moves, count := getKnightMoves(x, y)
	// fmt.Println(moves)
	// fmt.Println("Number of possible moves:", count)

	// move := getBestMove(x, y)
	// fmt.Println(move)

	// for i := 0; i < 8; i++ {

	// 	for j := 0; j < 8; j++ {

	// 		// Print cell with width of 3, right aligned
	// 		fmt.Printf(" %3d", board[i][j])
	// 	}
	// 	fmt.Println("")
	// }

}
