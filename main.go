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

	domains, err := customerimporter.SortDomains(os.Args[1])
	if err != nil {
		log.Println(err)
	}

	for _, d := range domains {
		fmt.Printf("Domain: %s, Occurence: %d \n", d.Domain, d.Count)

	}

}
