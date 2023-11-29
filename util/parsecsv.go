package util

import (
	"encoding/csv"
	"log"
	"os"
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
