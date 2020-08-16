package customerimporter

// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var m map[string]int = make(map[string]int)
var domains []string

//DomainCount represents a data structure of email domains and its occurreence
type DomainCount struct {
	Domain string
	Count  int
}

func processLine(email string) error {

	domain := strings.Split(email, "@")

	if len(domain) != 2 {
		return errors.New("variable does not contain a valid email")
	}

	value, exists := m[domain[1]]

	if exists == true {
		m[domain[1]] = value + 1
	} else {
		m[domain[1]] = 1
	}

	return nil

}

func sortKeys(records map[string]int) []DomainCount {

	keys := []string{}
	for key := range records {
		keys = append(keys, key)
	}
	var sortedrecords []DomainCount

	sort.Strings(keys)

	for _, domain := range keys {
		sortedrecords = append(sortedrecords[:], DomainCount{Domain: domain, Count: records[domain]})
	}

	return sortedrecords
}

func checkCSV(header []string) (bool, error) {
	if len(header) != 5 { //check if the csv file contains the required amount of fields
		return false, errors.New("Invalid csv header")
	}

	//todo: implement a more robust validation
	return strings.Contains(header[0], "first_name"), nil

}

//SortDomains returns a sorted []DomainCount and any encountered error.
func SortDomains(path string) ([]DomainCount, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	csvfile := csv.NewReader(file)

	for {
		row, err := csvfile.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err

		}
		//check if a row contains a header, which occurs every 1k lines. This check it is also used to detect invalid rows
		headerRow, err := checkCSV(row)

		if err != nil {
			return nil, err
		}

		//check if its a header
		if headerRow {
			continue
		}

		err = processLine(row[2])

		if err != nil {
			return nil, err
		}

	}

	return sortKeys(m), nil

}
