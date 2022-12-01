package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var maxCals, currentCals int
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
			if currentCals > maxCals {
				maxCals = currentCals
			}
			currentCals = 0
		}
	}
	fmt.Println(maxCals)
}
