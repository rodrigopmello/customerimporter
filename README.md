# customerimporter

customerimporter reads from the given csv file and returns a sorted (data structure of your choice) of email domains along with the number of customers with e-mail addresses for each domain.


## Run

```
$ go run main.go csv/file/path
```

## Test

ci_test.go contains 2 tests. The first one tests how my application handles different CSV files (empty, invalid format, invalid email, and valid). The second one tests the output of the application.

```
$ go test ./customerimporter -v
```


## Benchmark

ci_test.go also contains a simple benchmark of my application. To run the benchmark just insert the following command from the root src:

```
$ go test -bench=. ./customerimporter
```

In future versions, this benchmark could be improved. 

## Makefile

There is also a simple makefile that can be used instead of go commands. The available commands are: all, build, tests, and benchmark. If you use only make or make all, it is required to pass a csv file path in CSV variable. 


```
$ make CSV=customers.csv
```


## Files

The csv files used during tests can be found in customerimporter/csv_tests. 


## Important

During this implementation I've assumed a few things: (i) the header repetition could occur in any line, which causes some unnecessary checking; (ii) empty files does not cause any errors; (iii) if the CSV file contains an invalid email, the function returns an error.



