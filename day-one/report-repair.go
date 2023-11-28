package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const year = 2020

// parse csv file into array
func main() {
	// open file
	f, err := os.Open("day-one/puzzle-input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// extract data into slice
	var input []int
	for _, row := range data {
		n, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}

	// sort array
	sort.Ints(input)

	// using loops
	// part 1
	for i := 0; i < len(input); i++ {
		for j := i + 1; i < len(input)-j; j++ {
			if input[i]+input[j] == year {
				fmt.Println(input[i] * input[j])
			}
		}
	}

	// part 2
	for i := 0; i < len(input); i++ {
		for j := i + 1; i < len(input)-j; j++ {
			for k := j + 1; j < len(input)-k; k++ {
				if input[i]+input[j]+input[k] == year {
					fmt.Println(input[i] * input[j] * input[k])
				}
			}
		}
	}

	// using map
	// part 1
	x, y, err := findTarget2n(input, year)
	if err != nil {
		panic(err)
	}
	fmt.Printf("map part 1: result: %v\n", x*y)

	// part 2
	result, err := findTarget3n(input, year)
	if err != nil {
		panic(err)
	}
	fmt.Printf("map part 2: result: %v\n", result)

}

func findTarget2n(input []int, target int) (int, int, error) {
	// create map and populate with value and its complement
	var inputMap = make(map[int]int, len(input))
	for _, n := range input {
		inputMap[n] = target - n
	}

	// seach map to find if complement exists
	for n, comp := range inputMap {
		_, ok := inputMap[comp]
		if ok {
			return n, comp, nil
		}
	}
	return 0, 0, errors.New("unable to find a pair")
}

func findTarget3n(input []int, target int) (int, error) {
	// create map with value and subtarget
	var inputMap = make(map[int]int, len(input))
	for _, n := range input {
		inputMap[n] = target - n
	}

	for k, v := range inputMap {
		x, y, err := findTarget2n(input, v)
		if err == nil {
			return k * x * y, nil
		}
	}
	return 0, errors.New("unable to find complements")
}

// 2000 target

// 300, 700, 1000
// 300 n
// 1700 subTarget
//
// 1700 - 700 = 1000

// 2000 = n + x + y
// n = 2000 - x - y
// 2000 - n = x + y
