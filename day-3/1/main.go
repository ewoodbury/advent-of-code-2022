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

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		r0 := line[:len(line)/2]
		r1 := line[len(line)/2:]

		s0 := SortString(r0)
		s1 := SortString(r1)

		// Search both sorted strings for first value in both:
		var common rune

		p0 := 0
		p1 := 0
		// loop until the end of either string is reached
		for p0 < len(s0) && p1 < len(s1) {
			// if the characters at the current indexes in each string are the same,
			// set the common character and break out of the loop
			if s0[p0] == s1[p1] {
				common = rune(s0[p0])
				break
			}

			// if the character in str1 is lexicographically smaller,
			// move the pointer in str1 forward
			if s0[p0] < s1[p1] {
				p0++
			} else {
				// otherwise, move the pointer in str2 forward
				p1++
			}
		}
		if unicode.IsUpper(common) {
			score += strings.Index(uppercase, string(common)) + 26 + 1
		} else {
			score += strings.Index(lowercase, string(common)) + 1
		}
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(score)
}
