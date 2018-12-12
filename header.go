package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func header(filename string) []string {
	// Load a CSV file.
	f, _ := os.Open(filename)
	var i int
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	record, err := r.Read()
	defer f.Close()
	for {
		//Break after reading CSV header first line
		if i == 1 {
			break
		}
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		i++
	}
	return record

}
