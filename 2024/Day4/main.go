package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const XMAX = "XMAS"
const MAS = "MAS"

func main() {
	input := loadReport("puzzleInput.txt")
	fmt.Println(wordSearch(input))
}

func wordSearch(input [][]string) string {
	xmasCount := 0
	masxCount := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			xmasCount += countXMAS(input, x, y)
			masxCount += countMasX(input, x, y)
		}
	}
	return fmt.Sprintf("XMAS: %d, MASX: %d", xmasCount, masxCount)
}

func countMasX(input [][]string, x, y int) int {
	if x-1 < 0 || x+1 >= len(input[0]) || y-1 < 0 || y+1 >= len(input) || input[x][y] != string(MAS[1]) {
		return 0
	}

	left := false
	right := false

	if input[x-1][y-1] == string(MAS[0]) && input[x+1][y+1] == string(MAS[2]) ||
		input[x-1][y-1] == string(MAS[2]) && input[x+1][y+1] == string(MAS[0]) {
		left = true
	}

	if input[x-1][y+1] == string(MAS[0]) && input[x+1][y-1] == string(MAS[2]) ||
		input[x-1][y+1] == string(MAS[2]) && input[x+1][y-1] == string(MAS[0]) {
		right = true
	}

	if left && right {
		return 1
	}

	return 0
}

func countXMAS(input [][]string, x, y int) int {
	count := 0
	if searchXMASDir(input, x, y, 1, 0) {
		count++
	}
	if searchXMASDir(input, x, y, -1, 0) {
		count++
	}
	if searchXMASDir(input, x, y, 0, 1) {
		count++
	}
	if searchXMASDir(input, x, y, 0, -1) {
		count++
	}
	if searchXMASDir(input, x, y, 1, 1) {
		count++
	}
	if searchXMASDir(input, x, y, -1, -1) {
		count++
	}
	if searchXMASDir(input, x, y, 1, -1) {
		count++
	}
	if searchXMASDir(input, x, y, -1, 1) {
		count++
	}
	return count
}

func searchXMASDir(input [][]string, x, y, dx, dy int) bool {
	for i := 0; i < len(XMAX); i++ {
		xx := x + dx*i
		yy := y + dy*i
		if xx < 0 || xx >= len(input[0]) || yy < 0 || yy >= len(input) {
			return false
		}
		if input[yy][xx] != string(XMAX[i]) {
			return false
		}
	}
	return true
}

func loadReport(fileName string) [][]string {
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
		columns := strings.Split(scanner.Text(), " ")

		report = append(report, []string{})
		for _, column := range columns {
			report[len(report)-1] = append(report[len(report)-1], strings.Split(column, "")...)
		}
	}

	return report
}
