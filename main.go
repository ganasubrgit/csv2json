package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

var (
	JSONFILE = "file.json"
	CSVFILE  = "file.csv"
)

func main() {
	csvFile, err := os.Open(CSVFILE)
	check(err, 0) //error check
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	check(err, 1) //error check
	hvalue := header(CSVFILE)
	JSONKEY := map[string]string{}

	for _, each := range csvData {
		for i := 0; i < len(hvalue); i++ {
			JSONKEY[hvalue[i]] = each[i]
		}
		printjson(JSONKEY)
	}
}

func printjson(m map[string]string) {
	// Convert to JSON
	jsonData, err := json.Marshal(m)
	check(err, 1) //error check
	fmt.Println(string(jsonData))
	jsonFile, err := os.Create(JSONFILE)
	check(err, 0) //error check
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()

}
