package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("04a_input.txt")
	totalOverlaps := 0

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		innString := strings.Split(fileScanner.Text(), ",")
		range1 := strings.Split(innString[0], "-")
		range2 := strings.Split(innString[1], "-")
		range1From, _ := strconv.Atoi(range1[0])
		range1To, _ := strconv.Atoi(range1[1])
		range2From, _ := strconv.Atoi(range2[0])
		range2To, _ := strconv.Atoi(range2[1])

		if (range1From >= range2From && range1From <= range2To) || (range1To >= range2From && range1To <= range2To) {
			totalOverlaps++
		} else if (range2From >= range1From && range2From <= range1To) || (range2To >= range1From && range2To <= range1To) {
			totalOverlaps++
		}
	}
	fmt.Println("No overlaps: ", totalOverlaps)
}
