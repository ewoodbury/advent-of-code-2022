package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	score := 0

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		s0, s1 := strings.Split(line, ",")[0], strings.Split(line, ",")[1]

		start0, end0 := strings.Split(s0, "-")[0], strings.Split(s0, "-")[1]
		start1, end1 := strings.Split(s1, "-")[0], strings.Split(s1, "-")[1]

		strs := []string{start0, end0, start1, end1}

		// initialize a slice of integers to store the converted values
		ints := make([]int, len(strs))

		// iterate over the slice of strings and convert them to integers
		for i, str := range strs {
			// unpack the two values returned by strconv.Atoi() into separate variables
			x, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println(err)
				return
			}
			ints[i] = x
		}

		// Check whether one interval is enveloped:
		if (ints[0] <= ints[2] && ints[1] >= ints[3]) || (ints[2] <= ints[0] && ints[3] >= ints[1]) {
			score += 1
		}
	}

	fmt.Println(score)
}
