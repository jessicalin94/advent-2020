package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2020/day/5

var lowestRow, highestRow int

func main() {
	part1 := highestID("05/input.txt")
	fmt.Printf("part1: %v\n", part1)

	part2 := myID("05/input.txt")
	fmt.Printf("part2: %v\n", part2)

	fmt.Println(lowestRow, highestRow)
}

// part 2
func myID(path string) int {
	records := readInput(path)
	otherIDs := make(map[int]struct{})
	myID := 0
	for _, record := range records {
		otherIDs[seatID(record)] = struct{}{}
	}

	for i := 1; i <= highestRow; i++ {
		for j := 0; j <= 7; j++ {
			id := i*8 + j
			_, ok := otherIDs[id]
			if !ok {
				myID = id
			}
		}
	}
	return myID
}

// part 1
func highestID(path string) int {
	biggestID := 0
	records := readInput(path)
	for _, record := range records {
		id := seatID(record)
		if biggestID < id {
			biggestID = id
		}
	}
	return biggestID
}

func seatID(record string) int {
	row := findPosition(record[:7], 0, 127)
	col := findPosition(record[7:], 0, 7)

	if row < lowestRow {
		lowestRow = row
	}

	if row > highestRow {
		highestRow = row
	}

	seatID := row*8 + col
	return seatID
}

func findPosition(code string, floor, ceiling int) int {
	if len(code) == 1 {
		if code == "F" || code == "L" {
			return floor
		} else {
			return ceiling
		}
	}

	if code[0] == 'F' || code[0] == 'L' {
		return findPosition(code[1:], floor, floor+((ceiling-floor)/2))
	}

	if code[0] == 'B' || code[0] == 'R' {
		return findPosition(code[1:], ((ceiling-floor)/2)+floor+1, ceiling)
	}

	return 0
}

func readInput(path string) []string {
	// Open file and create scanner on top of it
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// each dataSlice element is a record as a string
	dataSlice := strings.Split(string(data), "\n")

	return dataSlice
}
