package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Chess board size
const size = 8

// Boolean board to track chosen squares
// at the beginning all of these are false
var validBoard [8][8]bool

type DecisonTree struct {
	coordinates [][]int
	Count       int
	InsertedX   int
	InsertedY   int
}

// Knight move offsets
var offsets = [][]int{
	{-2, 1}, {-2, -1},
	{-1, 2}, {-1, -2},
	{1, 2}, {1, -2},
	{2, 1}, {2, -1},
}

// func InithializeValidations() {
// 	for i := 0; i < 8; i++ {
// 		for j := 0; j < 8; j++ {
// 			validBoard[i][j] = true
// 		}
// 	}
// }

// Check if move is valid on board
func validMove(x, y int) bool {

	return x >= 0 && x < size && y >= 0 && y < size
}

func getKnightMoves(x, y int) [][]int {

	var moves [][]int

	for _, offset := range offsets {
		x1 := x + offset[0]
		y1 := y + offset[1]
		if validMove(x1, y1) {
			moves = append(moves, []int{x1, y1})
		}
	}

	return moves
}

func bestMove(searchlist [][]int) (int, int) {

	//bestMove := []Data{}
	//var posibbleMoves [][]int
	//cellCount := 0
	var decisionTrees []DecisonTree

	for _, move := range searchlist {
		//bestMove = getKnightMoves(move[0], move[1])
		//outputCounter := 8
		//posibbleMoves = getKnightMoves(move[0], move[1])
		move1 := getKnightMoves(move[0], move[1])
		newTree := DecisonTree{
			coordinates: move1,
			Count:       len(move1),
			InsertedX:   move[0],
			InsertedY:   move[1],
		}

		//tempPosibbleMoves := DecisonTree{}

		// tempPosibbleMoves.coordinates = append(tempPosibbleMoves.coordinates, move1...)
		// tempPosibbleMoves.Count = len(move1)

		//fmt.Println("Possible moves for", move[0], ", ", move[1], " : ", getKnightMoves(move[0], move[1]))
		//posibleMoves = append(posibleMoves, getKnightMoves(move[0], move[1]))
		//fmt.Println("posibleMoves", posibleMoves, "\n", "cellCount", cellCount)

		//fmt.Println("possiblemoves", newTree)

		decisionTrees = append(decisionTrees, newTree)

	}
	// todo : this is for find next move by lowest output to other cells
	// * sdfsdf ffdsfsdf
	// ? what is this
	// ! check this
	var lowest DecisonTree
	for _, tree := range decisionTrees {
		if (lowest.Count == 0 || tree.Count < lowest.Count) && (validBoard[lowest.InsertedX][lowest.InsertedY] == false) {
			fmt.Println("our condition is ", (lowest.Count == 0 || tree.Count < lowest.Count) && (validBoard[lowest.InsertedX][lowest.InsertedY] == false))
			lowest = tree
			validBoard[lowest.InsertedX][lowest.InsertedY] = true
		}
	}
	// ! check this

	// fmt.Println("Lowest Count:", lowest.Count)
	// fmt.Println("Coordinates:", lowest.coordinates)
	fmt.Println("Lowest is :", lowest)
	fmt.Println("Chosen Coordinates(x, y)(from bestmove function):", lowest.InsertedX, ", ", lowest.InsertedY)

	// fmt.Println("****************************************************************")
	// //fmt.Println("Best move Result:\n", "Possible moves: ", posibleMoves, "cellCount", cellCount)
	// Todo : this is for printing the decisiontrees slice in seperated lines
	// for _, tree := range decisionTrees {
	// 	fmt.Println("Possible Moves:", tree.coordinates)
	// 	fmt.Println("Count final:", tree.Count)
	// 	fmt.Println("x:", tree.InsertedX, "y:", tree.InsertedY)
	// }
	// "Chosen Coordinates(x, y):", lowest.InsertedX, ", ", lowest.InsertedY)
	return lowest.InsertedX, lowest.InsertedY
}

func ReadInput() (int, int, error) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter x,y: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, err
		}

		input = strings.TrimSpace(input)
		coords := strings.Split(input, ",")

		if len(coords) != 2 {
			fmt.Println("Invalid input. Please enter x,y")
			continue
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return 0, 0, err
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return 0, 0, err
		}

		// Todo this condition should validate entering coordinates
		// if validMove(x, y) != true {
		// 	fmt.Println("Invalid inpute. Please enter valid nmber between 0 and 7")
		// 	continue
		// }
		return x, y, nil
	}
}

func main() {

	var board [8][8]int
	level := 1

	// Starting position
	var startX, startY, err = ReadInput()
	if err != nil {
		fmt.Println("Invalid input(from main funciotn). Please valid enter x,y")
	}

	board[startX][startY] = level
	validBoard[startX][startY] = true

	// moves := getKnightMoves(startX, startY)
	// bestMove(moves)
	// fmt.Println("All knight moves are: ", moves, "Moves Count:", len((moves)))

	fmt.Println("****************************************************************")

	for i := 1; i < 10; i++ {

		track := getKnightMoves(startX, startY)
		fmt.Println("next possible moves for starting from(from main): ", startX, ", ", startY)
		fmt.Println("Tracked cells that one of them is valid", track)

		// Get next best move
		nextX, nextY := bestMove(track)
		fmt.Println("next best move: ", nextX, ", ", nextY)

		// Update board with move and new level
		level++
		board[nextX][nextY] = level

		// Print updated board
		//printBoard(board)

		// Set start to new position
		startX = nextX
		startY = nextY
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			fmt.Printf("%3d ", board[x][y])

		}
		fmt.Println()
		fmt.Println()
	}

}

//data := getKnightMoves(x, y)
//moves := data.Board
//count := data.Count

//fmt.Println("All Moves: ", moves, "\nAll Count", count)
//fmt.Println("**************************")

//for _, i := range moves {
//	fmt.Println("Move: ", i)
//	fmt.Print(" 0: ", i[0], " 1: ", i[1], "\n")
//}
