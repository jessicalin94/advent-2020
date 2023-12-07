package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func ParseCSV(path string) [][]string {
	// open file
	f, err := os.Open(path)
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

	return data
}

// ReadInput takes blank line separated text input which returns it as a slice of strings
func ReadInput(path string, delim string) []string {
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
	dataSlice := strings.Split(string(data), delim)

	return dataSlice
}
