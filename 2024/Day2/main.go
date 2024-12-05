package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const threshold = 3

func main() {
	totalSafe := 0
	reports := loadReport("puzzleInput.txt")

	for _, report := range reports {
		if isSafe(report) || Dampen(0, report) {
			totalSafe++
		}
	}

	fmt.Println(totalSafe)
}

func isSafe(report []int) bool {
	sign := 0
	safe := true
	for i := 1; i < len(report); i++ {
		d := report[i] - report[i-1]
		s := getSign(d)
		a := int(math.Abs(float64(d)))
		if sign == 0 {
			sign = s
		}
		if d == 0 || sign != s || a > threshold {
			safe = false
			break
		}
	}
	return safe
}

func getSign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

func Dampen(i int, report []int) bool {
	if i >= len(report) {
		return false
	}

	modified := append([]int{}, report...)
	modified = append(modified[:i], modified[i+1:]...)

	if isSafe(modified) {
		return true
	}

	return Dampen(i+1, report)
}

func loadReport(fileName string) [][]int {
	report := [][]int{}

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

		report = append(report, []int{})
		for _, col := range columns {
			// Assuming the columns contain integers
			var num int
			fmt.Sscanf(col, "%d", &num)
			report[len(report)-1] = append(report[len(report)-1], num)
		}
	}

	return report
}
