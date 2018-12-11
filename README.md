# csv2json
CSV to JSON converter 

Still in developement

### Usage

```csv2json
  -csvfile string
    	CSV file to convert (default "file.csv")
  -jsonfile string
    	JSON output file (default "file.json")
```

### Command line argument
```
csv2json -csvfile=file.csv -jsonfile=file.json

```

### Input files

![alt text](https://github.com/ganasubrgit/csv2json/blob/master/file.csv)

### Output 

```

{"address":"address","age":"age","name":"name","sex":"sex","working":"working"}
{"address":"2345 N ST","age":"12","name":"sam","sex":"male","working":"yes"}
{"address":"345 S Ln","age":"23","name":"john","sex":"male","working":"yes"}
{"address":"87 st","age":"34","name":"Jen","sex":"female","working":"no"}

```