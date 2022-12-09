package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	var buffer []string
	// Part 1
	// bufferSize := 4
	// Part 2
	bufferSize := 14

	// Loop over each char in line:
	for i, char := range line {
		// If buffer has fewer than 4 chars, append to buffer:
		if len(buffer) < bufferSize {
			buffer = append(buffer, string(char))
			continue
		} else {
			// if buffer has 4, check if the last 4 are unique. If so, return current index
			if isUnique(buffer) {
				fmt.Println(i)
				return
			} else {
				// If not, remove the first char and append the new char
				buffer = buffer[1:]
				buffer = append(buffer, string(char))
			}
		}
	}
	return
}

func isUnique(s []string) bool {
	m := make(map[string]uint, len(s)) //preallocate the map size
	for _, r := range s {
		m[r]++
	}

	for _, r := range s {
		if m[r] > 1 {
			return false
		}
	}
	return true
}
