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

	resultScoreMap := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	moveScoreMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	score := 0

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		p0 := line[0:1]
		p1 := line[2:3]

		score += resultScoreMap[p0][p1] + moveScoreMap[p1]
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(score)
}
