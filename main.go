package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var (
	JSONFILE = "file.json"
	CSVFILE  = "file.csv"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Sex       string
}

//error handeling
func check(err error, ext int) {
	if err != nil && ext == 1 {
		fmt.Println(err)
		os.Exit(1)
	} else if err != nil && ext == 0 {
		fmt.Println(err)
	}
}

func main() {
	csvFile, err := os.Open(CSVFILE)
	check(err, 0) //error check
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	check(err, 1) //error check
	var emp Employee
	var employees []Employee

	//fmt.Print(header(CSVFILE), "\n")
	hvalue := header(CSVFILE)

	//JSONKEY := map[int]string{}

	for value := range hvalue {

		fmt.Printf("  %v\n", hvalue[value])
	}

	for _, each := range csvData {
		emp.FirstName = each[0]
		emp.LastName = each[1]
		emp.Age, _ = strconv.Atoi(each[2])
		emp.Sex = each[3]
		employees = append(employees, emp)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(employees)
	check(err, 1) //error check
	fmt.Println(string(jsonData))
	jsonFile, err := os.Create(JSONFILE)
	check(err, 0) //error check
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
