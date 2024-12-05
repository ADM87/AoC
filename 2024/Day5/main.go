package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SortRule struct {
	Page     string
	Proceeds map[string]bool
}

func NewSortRule(page, order string) *SortRule {
	return &SortRule{
		Page:     page,
		Proceeds: map[string]bool{order: true},
	}
}

func main() {
	input := loadInput("puzzleInput.txt")
	rules := createRules(input[0])

	fmt.Println("Rules:")
	for _, rule := range rules {
		fmt.Printf("Page: %s, Proceeds: %v\n", rule.Page, rule.Proceeds)
	}

	// guides := getPageSequences(input[1])
	ordered, correct, updated := orderGuides(input[1], rules)

	fmt.Println("Correct Guides:")
	fmt.Println(correct)

	fmt.Println("Updated Guides:")
	fmt.Println(updated)

	fmt.Println("Ordered Guides:")
	for _, guide := range ordered {
		fmt.Println(guide)
	}

	fmt.Println("Total Middle:")
	fmt.Println(totalMiddle(ordered, updated))
}

func totalMiddle(guides [][]string, updated []int) int {
	total := 0
	for _, i := range updated {
		if i < 0 || i >= len(guides) {
			continue
		}

		m := len(guides[i]) / 2
		fmt.Println("Middle:", guides[i][m])

		v, err := strconv.Atoi(guides[i][m])
		if err != nil {
			fmt.Println(err)
			continue
		}

		total += v
	}
	return total
}

func orderGuides(input string, rules map[string]*SortRule) ([][]string, []int, []int) {
	ordered := [][]string{}
	updated := map[int]bool{}

	fmt.Println("Guides:")
	for g, guide := range getPageSequences(input) {
		fmt.Println(guide)

		updated[g] = false
		for i := 0; i < len(guide); i++ {
			for j := i + 1; j < len(guide); j++ {
				if rule, ok := rules[guide[j]]; ok && rule.Proceeds[guide[i]] {
					guide[i], guide[j] = guide[j], guide[i]
					updated[g] = true
				}
			}
		}

		ordered = append(ordered, guide)
	}

	fmt.Println("Updated Guides:")
	for g, guide := range ordered {
		if updated[g] {
			fmt.Printf("Guide %d: %v\n", g, guide)
		}
	}

	updatedIndices := []int{}
	correctIndices := []int{}

	for i, v := range updated {
		if v {
			updatedIndices = append(updatedIndices, i)
		} else {
			correctIndices = append(correctIndices, i)
		}
	}

	return ordered, correctIndices, updatedIndices
}

func createRules(input string) map[string]*SortRule {
	rules := map[string]*SortRule{}
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, "|")
		if r := rules[sp[0]]; r == nil {
			rules[sp[0]] = NewSortRule(sp[0], sp[1])
		} else {
			r.Proceeds[sp[1]] = true
		}
	}
	return rules
}

func getPageSequences(input string) [][]string {
	pages := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, ",")
		if sp[0] == "" {
			continue
		}
		pages = append(pages, sp)
	}

	return pages
}

func loadInput(fileName string) []string {
	report := ""

	var file *os.File
	var err error
	if file, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report = report + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	return strings.Split(report, "\n\n")
}
