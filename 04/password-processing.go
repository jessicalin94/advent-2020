package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var fieldsCheck = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func main() {
	part1 := countValidCredentials("04/input.txt")
	fmt.Printf("part1: %v\n", part1)
}

func countValidCredentials(inputPath string) int {
	count := 0
	records := readTextFile(inputPath)

	for _, record := range records {
		if isValid(record) {
			count++
		}
	}
	return count
}

func isValid(record string) bool {
	fields := strings.Fields(record)
	fieldMap := make(map[string]int, len(fields))
	for i, field := range fields {
		fieldMap[field] = i
	}
	for _, check := range fieldsCheck {
		if _, ok := fieldMap[check]; !ok {
			if check == "cid" {
				continue
			}
			return false
		}
	}
	return true
}

// readTextTile converts text data into a slice of records (string)
func readTextFile(path string) []string {
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
	dataSlice := strings.Split(string(data), "\n\n")

	// clean up data
	for i, v := range dataSlice {
		dataSlice[i] = strings.ReplaceAll(v, "\n", " ")
		dataSlice[i] = strings.ReplaceAll(dataSlice[i], ":", " ")
	}
	return dataSlice
}
