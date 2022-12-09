package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("../starting.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stacks := make(map[int][]string)

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			if char != ' ' {
				stacks[i+1] = append([]string{string(char)}, stacks[i+1]...)
			}
		}
	}

	// Read moves.txt
	file, err = os.Open("../moves.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner = bufio.NewScanner(file)

	// Loop over all moves, reading the format "move 2 from 8 to 4" into the variables
	// move, from, to
	for scanner.Scan() {
		line := scanner.Text()
		var move, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)

		// Pop `move` discs from `from` stack to `to` stack
		for i := 0; i < move; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
	fmt.Println(stacks)

	tops := []string{}
	for i := 1; i <= len(stacks); i++ {
		fmt.Println(i)
		tops = append(tops, stacks[i][len(stacks[i])-1])
	}
	fmt.Println(strings.Join(tops, ""))
}
