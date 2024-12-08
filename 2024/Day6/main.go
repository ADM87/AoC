package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Path  = "X"
	Obs   = "O"
	Block = "#"
	Empty = "."
)

// Directions are flipped because the grid is represented as [y][x] instead of [x][y]
var directions = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func main() {
	input := loadInput("testInput.txt")
	facing, pos := getFacing(input)
	visited, numLoops := patrol(input, pos, facing)

	fmt.Println("Visited:", visited)
	fmt.Println("Loops:", numLoops)
}

// Not the best, but it works
func patrol(grid [][]string, starting []int, facing int) (int, int) {
	visted := make(map[string]bool)
	numOfLoops := 0

	pos := []int{starting[0], starting[1]}

	// Check if redundant, but it keeps our loop going
	for checkBounds(grid, pos[0], pos[1]) {
		grid[pos[0]][pos[1]] = Path

		dir := directions[facing]
		next := []int{pos[0] + dir[0], pos[1] + dir[1]}

		if !checkBounds(grid, next[0], next[1]) {
			break
		}

		if isBlocked(grid, next) {
			facing = (facing + 1) % 4
			continue
		}

		pos[0] = next[0]
		pos[1] = next[1]

		visted[fmt.Sprintf("%d,%d", pos[0], pos[1])] = true
	}

	fmt.Println("Patrol Path:")
	for y := 0; y < len(grid); y++ {
		fmt.Println(grid[y])
	}

	return len(visted), numOfLoops
}

func getFacing(grid [][]string) (int, []int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[x][y] == "^" {
				return 0, []int{x, y}
			} else if grid[x][y] == ">" {
				return 1, []int{x, y}
			} else if grid[x][y] == "v" {
				return 2, []int{x, y}
			} else if grid[x][y] == "<" {
				return 3, []int{x, y}
			}
		}
	}
	return -1, []int{-1, -1}
}

func checkBounds(grid [][]string, x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func isBlocked(grid [][]string, pos []int) bool {
	return grid[pos[0]][pos[1]] == Block || grid[pos[0]][pos[1]] == Obs
}

func loadInput(fileName string) [][]string {
	report := [][]string{}

	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		report = append(report, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	return report
}
