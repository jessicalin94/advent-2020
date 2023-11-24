package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const target = 2020

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
	fmt.Println(input, len(input))

	// part 1
	for i := 0; i < len(input); i++ {
		for j := i + 1; i < len(input)-j; j++ {
			if input[i]+input[j] == target {
				fmt.Println(input[i] * input[j])
			}
		}
	}

	// part 2
	for i := 0; i < len(input); i++ {
		for j := i + 1; i < len(input)-j; j++ {
			for k := j + 1; j < len(input)-k; k++ {
				if input[i]+input[j]+input[k] == target {
					fmt.Println(input[i] * input[j] * input[k])
				}
			}
		}
	}
}
