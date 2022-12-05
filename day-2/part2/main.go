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

	moveScoreMap := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 1,
			"Z": 2,
		},
		"B": {
			"X": 1,
			"Y": 2,
			"Z": 3,
		},
		"C": {
			"X": 2,
			"Y": 3,
			"Z": 1,
		},
	}

	resultScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	score := 0

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		p0 := line[0:1]
		p1 := line[2:3]

		score += moveScoreMap[p0][p1] + resultScoreMap[p1]
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(score)
}
