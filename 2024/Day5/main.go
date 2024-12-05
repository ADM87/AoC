package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SortRule represents a rule for sorting pages.
type SortRule struct {
	Page     string          // The page number.
	Proceeds map[string]bool // A map of pages that this page proceeds.
}

// NewSortRule creates a new SortRule with the given page number and order.
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

// totalMiddle returns the sum of the middle values of the guides at the given indices.
func totalMiddle(guides [][]string, indices []int) int {
	total := 0
	// Iterate over the indices and sum the middle values of the guides.
	for _, i := range indices {
		if i < 0 || i >= len(guides) {
			continue
		}

		// Find the middle value of the guide.
		m := len(guides[i]) / 2
		fmt.Println("Middle:", guides[i][m])

		// Convert the middle value to an integer and add it to the total.
		v, err := strconv.Atoi(guides[i][m])
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Add the middle value to the total.
		total += v
	}
	return total
}

// orderGuides orders the guide pages based on the rules.
func orderGuides(input string, rules map[string]*SortRule) ([][]string, []int, []int) {
	ordered := [][]string{}   // A new list of sorted guides.
	updated := map[int]bool{} // Tracks the indices of guides that have been updated.

	fmt.Println("Guides:")
	for g, guide := range getPageGroups(input) {
		fmt.Println(guide)

		updated[g] = false

		// Iterate over the guide pages twice.
		// The first iteration is to find the pages that need to be swapped.
		// The second iteration is to swap the pages.
		for i := 0; i < len(guide); i++ {
			for j := i + 1; j < len(guide); j++ {
				// If the page at index i proceeds the page at index j, swap the pages.
				if rule, ok := rules[guide[j]]; ok && rule.Proceeds[guide[i]] {
					// i proceeds j, swap the pages.
					guide[i], guide[j] = guide[j], guide[i]
					// Mark the guide as updated.
					updated[g] = true
				}
			}
		}

		// Append the sorted guide to the ordered list.
		ordered = append(ordered, guide)
	}

	fmt.Println("Updated Guides:")
	for g, guide := range ordered {
		if updated[g] {
			fmt.Printf("Guide %d: %v\n", g, guide)
		}
	}

	updatedIndices := []int{} // Holds all the indices of the guides that were updated.
	correctIndices := []int{} // Holds all the indices of the guides that were not updated.

	for i, v := range updated {
		if v {
			updatedIndices = append(updatedIndices, i)
		} else {
			correctIndices = append(correctIndices, i)
		}
	}

	return ordered, correctIndices, updatedIndices
}

// createRules builds a map of SortRules from the input string.
// Each SortRule contains a page number and all the pages that it proceeds.
func createRules(input string) map[string]*SortRule {
	rules := map[string]*SortRule{}
	for _, line := range strings.Split(input, "\n") {
		// Split the line on the pipe: Ex "75|93" -> ["75", "93"]
		// The first element is the page number and the second element is the page it proceeds.
		sp := strings.Split(line, "|")
		if r := rules[sp[0]]; r == nil {
			rules[sp[0]] = NewSortRule(sp[0], sp[1])
		} else {
			r.Proceeds[sp[1]] = true
		}
	}
	return rules
}

// getPageGroups returns a slice of slices of strings representing the the groups of pages.
func getPageGroups(input string) [][]string {
	pages := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		// Split the line on the comma: Ex "75,93",45,23" -> ["75", "93", "45", "23"]
		sp := strings.Split(line, ",")
		// If the split line is empty, skip it.
		if sp[0] == "" {
			continue
		}
		pages = append(pages, sp)
	}

	return pages
}

// loadInput reads the input file and returns a slice of strings representing the input data.
// Splits the input on the double newline character: \n\n
// The first element is the rules and the second element is the guides.
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
