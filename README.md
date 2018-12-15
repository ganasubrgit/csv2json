# csv2json
CSV to JSON converter 

***beta***

### git clone

```
git clone https://github.com/ganasubrgit/csv2json

```

### build

```

cd csv2json/
go build .

```


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

```
cat file.csv
name,address,age,sex,working
sam,2345 N ST,12,male,yes
john,345 S Ln,23,male,yes
Jen,87 st,34,female,no

```
### Output 

```

{"address":"address","age":"age","name":"name","sex":"sex","working":"working"}
{"address":"2345 N ST","age":"12","name":"sam","sex":"male","working":"yes"}
{"address":"345 S Ln","age":"23","name":"john","sex":"male","working":"yes"}
{"address":"87 st","age":"34","name":"Jen","sex":"female","working":"no"}

```