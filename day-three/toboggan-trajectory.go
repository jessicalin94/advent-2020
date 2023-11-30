package main

import (
	"aoc-2020/util"
	"fmt"
)

var mapWidth, treeCount int
var index = 4

func main() {
	input := util.ParseCSV("day-three/input.csv")

	// convert input into slice
	inputArr := make([]string, len(input))
	for i, row := range input {
		inputArr[i] = row[0]
	}
	// determine width of row
	mapWidth = len(inputArr[0])

	// start from row 1,
	for i := 1; i < len(inputArr); i++ {
		// fmt.Printf("i: %v, index: %v ---> ", i+1, index+1)
		// fmt.Printf("string(inputArr[i][index]): %v\n", string(inputArr[i][index]))
		if inputArr[i][index] == '#' {
			treeCount++
		}
		// maintain lateral index within 0 - 30
		if index = index + 3; index >= mapWidth {
			index = index - mapWidth
		}
	}

	fmt.Printf("treeCount: %v\n", treeCount)
}

// 29 + 3 = 32
// goal = 1

// danger: 28(31), 29(32), 30(33)
// compen: 0, 1, 2
// 39277
