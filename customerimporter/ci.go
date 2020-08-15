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

//Readcsv function responsible to read a csv file based on the path
func Readcsv(path string) error {

	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	csvfile := csv.NewReader(file)

	for {
		row, err := csvfile.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err

		}

		if strings.Contains(row[0], "first_name") == true { //check if its a header
			continue
		}

		processLine(row[2])

	}

	sortOccurences(m)

	return nil

}

//https://play.golang.org/p/SYG9dQV-ir

//ReadcsvConc function responsible to read a csv file based on the path
func ReadcsvConc(path string) error {

	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	csvfile := csv.NewReader(file)

	//TODO: CHECK DELIMITERS IN TEST

	for {
		row, err := csvfile.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err

		}

		if strings.Contains(row[0], "first_name") == true { //check if its a header
			continue
		}

		processline2(row[2])
	}
	sortoccurences2(m)
	//log.Println(aux)
	//log.Println(m)
	return nil

}

var m map[string]int = make(map[string]int)

//DomainCount struct that represents a pair of doamin and count
type DomainCount struct {
	Domain string
	Count  int
}

func processLine(email string) error {

	domain := strings.Split(email, "@")

	if len(domain) != 2 {
		return errors.New("variable does not contain a valid email")
	}

	//log.Println(domain)

	value, exists := m[domain[1]]

	if exists == true {
		m[domain[1]] = value + 1
	} else {
		m[domain[1]] = 1
	}

	return nil

}

var domains []string

func processline2(email string) {

	domain := strings.Split(email, "@")[1]

	value, exists := m[domain]

	if exists == true {
		m[domain] = value + 1
	} else {
		domains = append(domains[:], domain)
		m[domain] = 1
	}

}

func sortOccurences(records map[string]int) []DomainCount {

	// get the list of keys and sort them
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

func sortoccurences2(records map[string]int) []DomainCount {

	var sortedrecords []DomainCount

	sort.Strings(domains)

	for _, domain := range domains {
		sortedrecords = append(sortedrecords[:], DomainCount{Domain: domain, Count: records[domain]})
	}

	//log.Println(sortedrecords)
	return sortedrecords
}

func checkCSV(header []string) (bool, error) {
	if len(header) != 5 { //check if the csv file contains the required amount of fields
		return false, errors.New("Invalid csv header")
	}

	//todo: implement a more robust checking
	return strings.Contains(header[0], "first_name"), nil

}

func checkCSVHeader(header []string) bool {

	//todo: implement a more robust checking
	return len(header) != 5

}

//SortDomains search how they document
func SortDomains(path string) ([]DomainCount, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	csvfile := csv.NewReader(file)

	header, err := csvfile.Read()
	//executes initial validation, such as empty file
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !checkCSVHeader(header) {
		return nil, errors.New("Invalid csv")
	}

	for {
		row, err := csvfile.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err

		}

		if strings.Contains(row[0], "first_name") == true { //check if a row contains a header, which occurs every 1k lines
			continue
		}

		err = processLine(row[2])

		if err != nil {
			return nil, err
		}

	}

	return sortOccurences(m), nil

}

//SortDomains2 search how they document
func SortDomains2(path string) ([]DomainCount, error) {
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

	return sortOccurences(m), nil

}
