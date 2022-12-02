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
	case "X": // Lose

		if opponent == "A" {
			points = points + 3
		}
		if opponent == "B" {
			points = points + 1
		}
		if opponent == "C" {
			points = points + 2
		}

	case "Y": // Draw
		points = points + 3

		if opponent == "A" {
			points = points + 1
		}
		if opponent == "B" {
			points = points + 2
		}
		if opponent == "C" {
			points = points + 3
		}

	case "Z": // Win
		points = points + 6

		if opponent == "A" {
			points = points + 2
		}
		if opponent == "B" {
			points = points + 3
		}
		if opponent == "C" {
			points = points + 1
		}
	}
	return points
}

/*
X = loss
Y = draw
Z = victory'

A = Rock -> 1 point
B = Paper -> 2 points
C = Scissors -> 3 points
*/
