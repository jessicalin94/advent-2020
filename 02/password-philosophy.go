package main

import (
	"aoc-2020/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ParseCSV("day-two/input.csv")

	// part 1
	fmt.Printf("part 1: there are %v valid passwords\n", validPasswordCount1(input))

	// part 2
	fmt.Printf("part 2: there are %v valid passwords\n", validPasswordCount2(input))
}

// part 1
func validPasswordCount1(input [][]string) int {
	var count int

	// given a set, break down components
	for _, set := range input {
		str := set[0]
		splitStr := strings.Split(str, " ")

		if validPassword1(splitStr) {
			count++
		}
	}
	return count
}

func validPassword1(set []string) bool {
	var min, max int
	var letter, str string

	for i, component := range set {
		if i == 0 {
			nums := strings.Split(component, "-")
			min, _ = strconv.Atoi(nums[0])
			max, _ = strconv.Atoi(nums[1])
		}

		if i == 1 {
			letter = strings.TrimSuffix(component, ":")
		}

		if i == 2 {
			str = component
		}

	}

	occ := strings.Count(str, letter)
	if occ >= min && occ <= max {
		return true
	}
	return false
}

// //////////////////////////////////////////////
// part 2
func validPasswordCount2(input [][]string) int {
	var count int

	for _, set := range input {
		str := set[0]
		splitStr := strings.Split(str, " ")

		if validPassword2(splitStr) {
			count++
		}
	}
	return count
}

func validPassword2(set []string) bool {
	var ind1, ind2 int
	var str string
	var char byte

	for i, component := range set {
		if i == 0 {
			indexes := strings.Split(component, "-")
			ind1, _ = strconv.Atoi(indexes[0])
			ind2, _ = strconv.Atoi(indexes[1])

			ind1--
			ind2--
		}
		if i == 1 {
			char = strings.TrimSuffix(component, ":")[0]
		}

		if i == 2 {
			str = component
		}
	}
	if str[ind1] == char && str[ind2] == char {
		return false
	}

	if str[ind1] != char && str[ind2] != char {
		return false
	}
	return true
}
