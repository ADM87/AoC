package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := loadLists("puzzleInput.txt")
	occurances := make([]int, len(left))

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	score := 0

	for i := 0; i < len(left); i++ {
		totalDistance += int(math.Abs(float64(left[i]) - float64(right[i])))
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				occurances[i]++
			}
		}
	}

	for i := 0; i < len(occurances); i++ {
		score += occurances[i] * left[i]
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Score: %d\n", score)
}

func loadLists(fileName string) ([]int, []int) {
	left := []int{}
	right := []int{}

	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "   ")
		value, err := strconv.Atoi(columns[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		left = append(left, value)

		value, err = strconv.Atoi(columns[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		right = append(right, value)
	}

	return left, right
}
