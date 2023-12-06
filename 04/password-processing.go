package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fieldsCheck = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func main() {
	fmt.Println(len("1994"))
	part1 := countValidCredentials("04/input.txt")
	fmt.Printf("part1: %v\n", part1)

	part2 := strictCheck("04/input.txt")
	fmt.Printf("part2: %v\n", part2)
}

// part 2
func strictCheck(inputPath string) int {
	count := 0
	records := readTextFile(inputPath)
	r, _ := regexp.Compile("^#[0-9a-f]{6}$")

	for _, record := range records {
		if !isValid(record) {
			continue
		}
		strictValid := true
		recordMap := make(map[string]string, len(record))
		split := strings.Split(record, " ")
		for i, v := range split {
			if i%2 == 0 {
				recordMap[string(v)] = string(split[i+1])
			}
		}
		for k, v := range recordMap {
			if k == "byr" {
				year, err := parseYear(v)
				if err != nil {
					strictValid = false
					break
				}
				if year < 1920 || year > 2002 {
					strictValid = false
					break
				}

			}

			if k == "iyr" {
				year, err := parseYear(v)
				if err != nil {
					strictValid = false
					break
				}
				if year < 2010 || year > 2020 {
					strictValid = false
					break
				}
			}

			if k == "eyr" {
				year, err := parseYear(v)
				if err != nil {
					strictValid = false
					break
				}
				if year < 2020 || year > 2030 {
					strictValid = false
					break
				}
			}

			if k == "hgt" {
				if !(strings.HasSuffix(v, "cm") || strings.HasSuffix(v, "in")) {
					strictValid = false
					break
				}

				if strings.HasSuffix(v, "cm") {
					cm := strings.TrimSuffix(v, "cm")
					n, _ := strconv.Atoi(cm)
					if n < 150 || n > 193 {
						strictValid = false
						break
					}
				}

				if strings.HasSuffix(v, "in") {
					in := strings.TrimSuffix(v, "in")
					n, _ := strconv.Atoi(in)
					if n < 59 || n > 76 {
						strictValid = false
						break
					}
				}
			}

			if k == "hcl" {
				if !r.MatchString(v) {
					strictValid = false
					break
				}
			}

			if k == "ecl" {
				if !(v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth") {
					strictValid = false
					break
				}
			}

			if k == "pid" {
				match, _ := regexp.MatchString("^[0-9]{9}$", v)
				if !match {
					strictValid = false
					break
				}
			}
		}
		if strictValid {
			count++
		}
	}
	return count
}

func parseYear(year string) (int, error) {
	if len(year) != 4 {
		return 0, errors.New("invalid year")
	}

	y, _ := strconv.Atoi(year)
	return y, nil
}

// part 1
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
	fieldMap := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		fieldMap[field] = struct{}{}
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
