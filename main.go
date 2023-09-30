package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

// Chess board size
// Todo: should reconsider decaling variables and try to decrease them and use goroutines as much and as reasonably as possible.

const size = 8

var visited = make(map[int][][]int)

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

// Check if move is valid on board
func validMove(x, y int) bool {

	return (x >= 0 && x < size && y >= 0 && y < size) && !contains(x, y)
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

	var decisionTrees []DecisonTree

	for _, move := range searchlist {

		move1 := getKnightMoves(move[0], move[1])
		newTree := DecisonTree{
			coordinates: move1,
			Count:       len(move1),
			InsertedX:   move[0],
			InsertedY:   move[1],
		}

		decisionTrees = append(decisionTrees, newTree)

	}

	x, y := lowestFinder(decisionTrees)

	return x, y
}

func ReadInput() (int, int, error) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter x,y from 0 to ", size, " : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, err
		}

		input = strings.TrimSpace(input)
		coords := strings.Split(input, ",")

		if len(coords) != 2 {
			fmt.Println("Invalid input. Please Enter x,y from 0 to ", size, " : ")
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

func lowestFinder(coords []DecisonTree) (int, int) {

	if len(coords) == 0 {

		return 0, 0
	}

	smallest := coords[0]
	for _, item := range coords {
		if !contains(item.InsertedX, item.InsertedY) {
			if item.Count <= smallest.Count {
				smallest = item
			}
		}
		if contains(smallest.InsertedX, smallest.InsertedY) {
			sorted := make([]DecisonTree, len(coords))
			copy(sorted, coords)

			// Sort copy
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].Count < sorted[j].Count
			})

			// Find first not visited
			for _, s := range sorted {
				if !contains(s.InsertedX, s.InsertedY) {
					smallest = s
				}
			}

		}
	}

	return smallest.InsertedX, smallest.InsertedY
}

func addToVisited(x, y int) {
	if !contains(x, y) {
		visited[x] = append(visited[x], []int{x, y})
	}
}

func contains(x, y int) bool {
	for _, c := range visited {
		for _, coordinate := range c {
			if coordinate[0] == x && coordinate[1] == y {
				return true
			}
		}
	}
	return false
}

func findMissingNumbers(board [size][size]int) []int {
	// Create a set to store the numbers in the board.
	numbers := make(map[int]struct{})
	for _, row := range board {
		for _, number := range row {
			numbers[number] = struct{}{}
		}
	}

	// Create a slice to store the missing numbers.
	missingNumbers := []int{}
	for i := 1; i <= 64; i++ {
		if _, ok := numbers[i]; !ok {
			missingNumbers = append(missingNumbers, i)
		}
	}

	// Return the slice of missing numbers.
	fmt.Println(missingNumbers)
	return missingNumbers
}

func main() {

	var board [size][size]int
	level := 1

	// Starting position
	// Todo: read size from input and modify related functions

	var startX, startY, err = ReadInput()
	if err != nil {
		fmt.Println("Invalid input(from main funciotn). Please Enter x,y from 0 to ", size, " : ")
		log.Fatal(err)

	}

	board[startX][startY] = level
	addToVisited(startX, startY)

	for i := 1; i < (size * size); i++ {

		track := getKnightMoves(startX, startY)

		// Get next best move
		nextX, nextY := bestMove(track)

		// Update board with move and new level
		level++
		board[nextX][nextY] = level
		addToVisited(nextX, nextY)

		// Set start to new position
		startX = nextX
		startY = nextY

	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			fmt.Printf("%6d ", board[x][y])

		}
		fmt.Println()
	}

	findMissingNumbers(board)
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

}
