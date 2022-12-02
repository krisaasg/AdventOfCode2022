package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("02a_input.txt")
	totalPoints := 0

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		points := 0
		innString := strings.Split(fileScanner.Text(), " ")

		points = calcPoints(innString[0], innString[1])
		totalPoints = totalPoints + points

	}
	fmt.Println(totalPoints)
}

func calcPoints(opponent string, own string) int {
	points := 0

	switch own {
	case "X": // Rock
		points = points + 1

		if opponent == "A" {
			points = points + 3
		}
		if opponent == "C" {
			points = points + 6
		}

	case "Y": // Paper
		points = points + 2

		if opponent == "B" {
			points = points + 3
		}
		if opponent == "A" {
			points = points + 6
		}

	case "Z": // Scissors
		points = points + 3

		if opponent == "C" {
			points = points + 3
		}
		if opponent == "B" {
			points = points + 6
		}
	}
	return points
}
