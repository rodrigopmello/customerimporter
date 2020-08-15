package main

import (
	"rankByDomain/customerimporter"
	"time"

	"github.com/yanzay/log"
)

func main() {

	t1 := time.Now()
	domains, err := customerimporter.SortDomains2("customers.csv")

	if err != nil {
		log.Println(err)
	}
	log.Println(domains)
	log.Println(time.Now().Sub(t1))

}
