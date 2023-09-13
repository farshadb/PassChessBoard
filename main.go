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

type Data struct {
	Board [][]int
	Count int
}

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

func getKnightMoves(x, y int) ([][]int, int) {
	var moves [][]int
	outputCount := 0
	for _, offset := range offsets {
		x1 := x + offset[0]
		y1 := y + offset[1]
		if validMove(x1, y1) {
			moves = append(moves, []int{x1, y1})
			outputCount++
		}
	}
	return moves, outputCount
	//return Data{
	//	Board: moves,
	//	Count: outputCount,
	//}
}

func bestMove(searchlist [][]int) {

	//bestMove := []Data{}
	var posibleMoves [][]int
	cellCount := 0

	for _, move := range searchlist {
		//bestMove = getKnightMoves()
		posibleMoves, cellCount = getKnightMoves(move[0], move[1])
		//fmt.Println("searchlist", getKnightMoves(move[0], move[1]))
		//posibleMoves = append(posibleMoves, getKnightMoves(move[0], move[1]))
		fmt.Println("posibleMoves", posibleMoves, "\n", "cellCount", cellCount)
	}
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

		return x, y, nil
	}
}

func main() {

	//var board [8][8]int
	//var moves [][]int
	var x, y, err = ReadInput()
	if err != nil {
		fmt.Println("Invalid input. Please enter x,y")
	}
	moves, count := getKnightMoves(x, y)
	//data := getKnightMoves(x, y)
	//moves := data.Board
	//count := data.Count
	fmt.Println("All Moves: ", moves, "\nAll Count", count)
	fmt.Println("**************************")
	bestMove(moves)
	//for _, i := range moves {
	//	fmt.Println("Move: ", i)
	//	fmt.Print(" 0: ", i[0], " 1: ", i[1], "\n")
	//}
}
