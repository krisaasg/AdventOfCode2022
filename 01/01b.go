package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var topCals, secondCals, thirdCals, totalCals, currentCals int
	readFile, err := os.Open("01a_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			intVar, intErr := strconv.Atoi(fileScanner.Text())
			// fmt.Println(intVar)
			if intErr == nil {
				currentCals = currentCals + intVar
			} else {
				fmt.Println(intErr)
			}

		} else { // Emtpy line = Calculation done for current Elf
			if currentCals > topCals {
				thirdCals = secondCals
				secondCals = topCals
				topCals = currentCals
			} else if currentCals > secondCals {
				thirdCals = secondCals
				secondCals = currentCals
			} else if currentCals > thirdCals {
				thirdCals = currentCals
			}
			currentCals = 0
		}
	}
	totalCals = topCals + secondCals + thirdCals
	fmt.Println(totalCals)
}
