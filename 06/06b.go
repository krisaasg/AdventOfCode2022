package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("06a_input.txt")
	windowSize := 14
	windowEnd := windowSize - 1 // Position of first window end

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		for windowStart := 0; windowStart <= windowEnd; windowStart++ {
			dupe := false

			if windowEnd <= len(fileScanner.Text()) { // Stops when there is not room for an entire new window
				for i := windowStart; i < windowEnd; i++ {
					for j := i + 1; j <= windowEnd; j++ {
						if fileScanner.Text()[i] == fileScanner.Text()[j] {
							dupe = true
						}
					}
				}
				windowEnd++

				if dupe == false {
					fmt.Println("Found start at ", windowEnd)
					break
				}

			}

		}
	}
}
