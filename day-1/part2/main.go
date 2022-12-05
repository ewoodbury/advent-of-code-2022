package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	maxCals := []int{0, 0, 0}
	currCals := 0

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is blank, save the value as null
		if line == "" {
			if currCals > maxCals[0] {
				maxCals[0] = currCals
				sort.Ints(maxCals)
			}
			currCals = 0
			continue
		}

		// Parse the integer on the line
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			continue
		}

		currCals += num
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(maxCals)
	topThreeSum := 0
	for i := range maxCals {
		topThreeSum += maxCals[i]
	}
	fmt.Println(topThreeSum)
}
