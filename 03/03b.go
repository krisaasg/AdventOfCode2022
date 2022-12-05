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

	var currentTeam [3]string
	i := 0
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// add each line to an array
		if i <= 2 {
			currentTeam[i] = fileScanner.Text()
			i++
		}
		if i == 3 { // Three lines present in array CurrentTeam
		out:
			for j := 0; j < len(currentTeam[0]); j++ {
				for k := 0; k < len(currentTeam[1]); k++ {
					if currentTeam[0][j] == currentTeam[1][k] { // Found duplicate between first and second bag
						for l := 0; l < len(currentTeam[2]); l++ {
							if currentTeam[0][j] == currentTeam[2][l] { // Found duplicate between all three bags
								var common, charValue byte
								common = (currentTeam[0][j])
								if common >= 65 && common <= 90 { // Uppercase
									charValue = common - 38
								} else if common >= 97 && common <= 122 { // Lowercase
									charValue = common - 96
								}
								totalPoints = totalPoints + int(charValue)
								break out
							}
						}
					}

				}
			}
			i = 0
		}
	}
	fmt.Println("Total Points: ", totalPoints)
}

/*
97-122 - lowercase
65-90 - uppercase
*/
