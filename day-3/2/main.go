package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// Helper functions for sorting strings:
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// Helper function for finding common element from 3 sorted strings
func findCommonElement(strs []string) string {
	// initialize three pointers to keep track of the current index in each string
	p1 := 0
	p2 := 0
	p3 := 0

	// loop until the end of any string is reached
	for p1 < len(strs[0]) && p2 < len(strs[1]) && p3 < len(strs[2]) {
		// if the characters at the current indexes in each string are the same,
		// return the character as a string
		if strs[0][p1] == strs[1][p2] && strs[1][p2] == strs[2][p3] {
			return string(strs[0][p1])
		}

		// find the minimum character among the three strings at their current indexes
		minChar := strs[0][p1]
		if strs[1][p2] < minChar {
			minChar = strs[1][p2]
		}
		if strs[2][p3] < minChar {
			minChar = strs[2][p3]
		}

		// move the pointer in the string with the minimum character forward
		if minChar == strs[0][p1] {
			p1++
		} else if minChar == strs[1][p2] {
			p2++
		} else {
			p3++
		}
	}

	// if the loop ends without finding a common element, return an empty string
	return ""
}

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	score := 0

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	var currentLines []string

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		sortedLine := SortString(line)

		currentLines = append(currentLines, sortedLine)

		if len(currentLines) == 3 {

			common := findCommonElement(currentLines)

			// Check if uppercase by converting to rune, then using unicode.IsUpper
			if unicode.IsUpper([]rune(common)[0]) {
				score += strings.Index(uppercase, string(common)) + 26 + 1
			} else {
				score += strings.Index(lowercase, string(common)) + 1
			}

			currentLines = make([]string, 0)

		}
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(score)
}
