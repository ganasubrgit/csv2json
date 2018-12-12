package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	CSVFILE  = flag.String("csvfile", "file.csv", "CSV file to convert")
	JSONFILE = flag.String("jsonfile", "file.json", "JSON output file")
)

func main() {

	flag.Parse()
	args := os.Args
	if len(args) <= 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	os.Remove(*JSONFILE)
	csvFile, err := os.Open(*CSVFILE)
	check(err, 0) //error check
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	csvData, err := reader.ReadAll()
	check(err, 1) //error check
	hvalue := header(*CSVFILE)
	JSONKEY := map[string]string{}
	for _, each := range csvData {
		for i := 0; i < len(hvalue); i++ {
			JSONKEY[hvalue[i]] = each[i]
		}
		printjson(JSONKEY)
	}
}

func printjson(m map[string]string) {
	//Convert to JSON
	jsonData, err := json.Marshal(m)
	check(err, 1) //error check
	fmt.Println(string(jsonData))
	jsonFile, err := os.OpenFile(*JSONFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err, 0) //error check
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()

}
