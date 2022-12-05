package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("05a_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var stacks [9][]string
	stacks[0] = []string{"F", "T", "C", "L", "R", "P", "G", "Q"}
	stacks[1] = []string{"N", "Q", "H", "W", "R", "F", "S", "J"}
	stacks[2] = []string{"F", "B", "H", "W", "P", "M", "Q"}
	stacks[3] = []string{"V", "S", "T", "D", "F"}
	stacks[4] = []string{"Q", "L", "D", "W", "V", "F", "Z"}
	stacks[5] = []string{"Z", "C", "L", "S"}
	stacks[6] = []string{"Z", "B", "M", "V", "D", "F"}
	stacks[7] = []string{"T", "J", "B"}
	stacks[8] = []string{"Q", "N", "B", "G", "L", "S", "P", "H"}

	for fileScanner.Scan() {
		innString := strings.Split(fileScanner.Text(), " ") // Numbers found at innString[1/3/5]
		move, _ := strconv.Atoi(innString[1])
		from, _ := strconv.Atoi(innString[3])
		to, _ := strconv.Atoi(innString[5])

		for i := 0; i < move; i++ { // Move one item at a time to new slice
			stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-(move-i)])
		}
		for i := 0; i < move; i++ { // Remove from old slice
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		}
	}

	stackTops := ""
	for i := 0; i < len(stacks); i++ {
		stackTops = stackTops + stacks[i][len(stacks[i])-1]
	}
	fmt.Println("Top of stacks: ", stackTops)
}
