package main

import (
	"customerimporter/customerimporter"
	"fmt"
	"os"

	"github.com/yanzay/log"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments. Insert a path to a csv file.")
		return
	}

	file, err := os.Open(os.Args[1])

	defer file.Close()

	if err != nil {
		log.Println(err)
		panic(err)

	}

	domains, err := customerimporter.SortDomains(file)
	if err != nil {
		log.Println(err)
	}

	for _, d := range domains {
		fmt.Printf("Domain: %s, Occurrence: %d \n", d.Domain, d.Count)

	}

}
