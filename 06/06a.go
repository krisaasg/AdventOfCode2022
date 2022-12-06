package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("06a_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		for position, letter := range fileScanner.Text() {

			if position < len(fileScanner.Text())-4 {

				dupes := false

				for i := 1; i <= 4; i++ {
					if letter == rune(fileScanner.Text()[position+i]) {
						dupes = true
						break
					}
				}
				if fileScanner.Text()[position+1] == fileScanner.Text()[position+2] || fileScanner.Text()[position+1] == fileScanner.Text()[position+3] {
					dupes = true
				} else if fileScanner.Text()[position+2] == fileScanner.Text()[position+3] {
					dupes = true
				}

				if dupes == false {
					fmt.Println("Start detected at ", position+4)
					break
				}
			}

		}
	}
}
