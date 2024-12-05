package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("puzzleInput.txt")
	muls := getEnabledMulStrings(input)

	total := 0
	for _, mul := range muls {
		total += execMul(mul)
	}

	fmt.Printf("Total: %d\n", total)
}

func execMul(mul string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	xy := re.FindStringSubmatch(mul)
	if len(xy) != 3 {
		fmt.Fprintf(os.Stderr, "Error: invalid mul string\n")
		os.Exit(1)
	}
	x, err := strconv.Atoi(xy[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	y, err := strconv.Atoi(xy[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	return x * y
}

func getEnabledMulStrings(input string) []string {
	dos := strings.Split(input, "do()")
	matches := []string{}

	for _, do := range dos {
		segments := strings.Split(do, "don't()")
		if len(segments) == 0 {
			continue
		}

		re := regexp.MustCompile(`mul\(\d+,\d+\)`) // this is the pattern Regex to match 'mul(x,y)'
		muls := re.FindAllString(input, -1)        // find all matches of the pattern in the input
		// muls is a list 'mul(x,y)' strings that match the pattern

		if len(muls) > 0 {
			matches = append(matches, muls...)
		}
	}

	return matches
}

func loadInput(fileName string) string {
	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	return string(data)
}
