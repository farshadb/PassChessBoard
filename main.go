package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Chess board size
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

	i := 1
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
	// * check this

	// for _, fighter := range decisionTrees {
	// 	fmt.Printf("%v %v %v %v\n", fighter.coordinates, fighter.Count, fighter.InsertedX, fighter.InsertedY)

	// }

	// fmt.Println(lowestFinder(decisionTrees))
	// fmt.Println("function lowestFinder called for ", i, "st times")
	i++
	x, y := lowestFinder(decisionTrees)
	// fmt.Println()
	// fmt.Println("----------------------------------")
	// fmt.Println(visited)
	// fmt.Println("----------------------------------")
	// fmt.Println("Recieved next move", x, ", ", y)
	// fmt.Println()

	return x, y
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
		//fmt.Println("from func inputreader: ", x, ", ", y)
		return x, y, nil
	}
}

func lowestFinder(coords []DecisonTree) (int, int) {

	if len(coords) == 0 {

		//fmt.Println("you have problem in you given data(form func lowestFinder)")
		return 0, 0
	}

	smallest := coords[0]
	for _, item := range coords {
		//fmt.Println("\ncheking ", item.InsertedX, ", ", item.InsertedY)
		if !contains(item.InsertedX, item.InsertedY) {
			// fmt.Println("+++++++++++++++++++++++++++++++++++")
			// fmt.Println(item.InsertedX, ", ", item.InsertedY, " Is not in visited")
			// fmt.Println(visited)
			// fmt.Println("+++++++++++++++++++++++++++++++++++")
			// ! fmt.Println(item.InsertedX, " ", item.InsertedY, " is not visited")
			if item.Count <= smallest.Count {
				//fmt.Println(item.InsertedX, ", ", item.InsertedY, " with output count ", item.Count, " is chosen")
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
					// fmt.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")
					// fmt.Println(s)
					// fmt.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")
					// fmt.Println()

					smallest = s
				}
			}

		}
	}

	//fmt.Println("lowest count(from func lowestfinder): ", smallest.Count)
	// fmt.Println("*****************************")
	// fmt.Println()
	// fmt.Println("Neeeext moves are: ", smallest.InsertedX, ", ", smallest.InsertedY)
	// fmt.Println("end of lowestfinder functoin\n")
	return smallest.InsertedX, smallest.InsertedY
}

func emptyValidCell(nextMin []DecisonTree) []DecisonTree {

	min := nextMin
	// fmt.Println("^^^^^^^^^^^^^^^^^^^^^^")
	// fmt.Println(min)
	// if contains(nextMin.InsertedX, nextMin.InsertedY) {
	// 	sorted := make([]DecisonTree, len(coords))
	// 	copy(sorted, nextMin)

	// 	// Sort copy
	// 	sort.Slice(sorted, func(i, j int) bool {
	// 		return sorted[i].Count < sorted[j].Count
	// 	})

	// 	// Find first not visited
	// 	for _, s := range sorted {
	// 		if !contains(s.InsertedX, s.InsertedY) {
	// 			fmt.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")
	// 			fmt.Println(s)
	// 			fmt.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")
	// 			fmt.Println()

	// 			min = s
	// 		}
	// 	}

	// }

	return min

}

func addToVisited(x, y int) {
	if !contains(x, y) {
		// fmt.Println("=============================================================")
		// fmt.Println(x, ", ", y, " Aded to visited")
		// fmt.Println("=============================================================")

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

func findMissingNumbers(board [8][8]int) []int {
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

	var board [8][8]int
	level := 1

	// Starting position
	var startX, startY, err = ReadInput()
	if err != nil {
		fmt.Println("Invalid input(from main funciotn). Please valid enter x,y")
		log.Fatal(err)

	}

	board[startX][startY] = level
	addToVisited(startX, startY)

	// fmt.Println("----------------------------------")
	// fmt.Println(visited)
	// fmt.Println("----------------------------------")

	// fmt.Println()
	// fmt.Println("#########")
	// fmt.Println("\n", startX, ", ", startY, " changed to visited")
	// fmt.Println(visited)
	// fmt.Println()

	// for x := 0; x < size; x++ {
	// 	for y := 0; y < size; y++ {
	// 		fmt.Printf("%3d ", board[x][y])

	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("**************************")
	// fmt.Println("")

	for i := 1; i < 64; i++ {

		track := getKnightMoves(startX, startY)
		// fmt.Println(i, " : next possible moves for starting from(from main): ", startX, ", ", startY)
		// fmt.Println("Tracked cells that one of them ", track)

		// Get next best move
		nextX, nextY := bestMove(track)
		//fmt.Println("next best move: ", nextX, ", ", nextY)

		// Update board with move and new level
		level++
		board[nextX][nextY] = level
		addToVisited(nextX, nextY)

		// fmt.Println("----------------------------------")
		// fmt.Println(visited)
		// fmt.Println("----------------------------------")

		// fmt.Println("\n", nextX, ", ", nextY, " changed to visited\n")
		// fmt.Println(visited)
		// fmt.Println()

		// Set start to new position
		startX = nextX
		startY = nextY

		// for x := 0; x < size; x++ {
		// 	for y := 0; y < size; y++ {
		// 		fmt.Printf("%3d ", board[x][y])

		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println("**************************")
		// fmt.Println("")

	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			fmt.Printf("%3d ", board[x][y])

		}
		fmt.Println()
	}

	findMissingNumbers(board)

}
