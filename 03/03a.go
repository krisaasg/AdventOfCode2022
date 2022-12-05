package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("03a_input.txt")
	totalPoints := 0

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		length := len(fileScanner.Text())
		compartment1 := fileScanner.Text()[0:(length / 2)]
		compartment2 := fileScanner.Text()[(length / 2):length]
		var misplaced, charValue byte
		var i, j int
		for i = 0; i < len(compartment1); i++ {
			for j = 0; j < len(compartment2); j++ {
				if compartment1[i] == compartment2[j] {
					misplaced = compartment1[i]
					break
				}
			}
		}
		if misplaced >= 65 && misplaced <= 90 { // Uppercase
			charValue = misplaced - 38
		} else if misplaced >= 97 && misplaced <= 122 { // Lowercase
			charValue = misplaced - 96
		}
		totalPoints = totalPoints + int(charValue)
	}
	fmt.Println("Total Points: ", totalPoints)
}

/*
97-122 - lowercase
65-90 - uppercase
*/
