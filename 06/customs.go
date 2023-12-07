package main

import (
	"aoc-2020/util"
	"fmt"
	"strings"
)

const (
	inputPath = "06/input.txt"
	exPath    = "06/example.txt"
)

func main() {
	part1 := findTotalCount()
	fmt.Printf("part1: %v\n", part1)
}

func findTotalCount() int {
	totalCount := 0
	inputs := util.ReadInput(inputPath, "\n\n")

	for i, input := range inputs {
		inputs[i] = strings.ReplaceAll(input, "\n", "")
	}

	for _, input := range inputs {
		groupCount := findGroupCount(input)
		totalCount += groupCount
	}
	return totalCount
}

func findGroupCount(groupSet string) int {
	set := map[rune]bool{}
	for _, letter := range groupSet {
		_, ok := set[letter]
		if !ok {
			set[letter] = true
		}
	}
	return len(set)
}
