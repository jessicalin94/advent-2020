package main

import (
	"aoc-2020/util"
	"fmt"
)

func main() {
	input := util.ParseCSV("day-three/input.csv")

	// convert input into slice
	inputArr := make([]string, len(input))
	for i, row := range input {
		inputArr[i] = row[0]
	}

	fmt.Printf("part1 treeCount: %v\n", part1(inputArr, 3, 1))
	fmt.Printf("part2 treeCount): %v\n", part2(inputArr))

}

func part1(inputArr []string, latRule, vertRule int) int {
	var rowInd, colInd, treeCount int

	for i := 0; i < len(inputArr)-vertRule; i = i + vertRule {
		rowInd += vertRule
		colInd = (colInd + latRule) % 31

		if inputArr[rowInd][colInd] == '#' {
			treeCount++
		}

	}
	return treeCount
}

func part2(inputArr []string) int {
	return part1(inputArr, 1, 1) * part1(inputArr, 3, 1) * part1(inputArr, 5, 1) * part1(inputArr, 7, 1) * part1(inputArr, 1, 2)
}
