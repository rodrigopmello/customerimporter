package customerimporter

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestErrorHandling(t *testing.T) {
	cases := []struct {
		fileName  string
		returnErr bool
		name      string
	}{
		{
			fileName:  "csv_tests/empty.csv",
			returnErr: false,
			name:      "EmptyFile",
		},
		{
			fileName:  "csv_tests/valid.csv",
			returnErr: false,
			name:      "ValidFile",
		},
		{
			fileName:  "csv_tests/invalid.csv",
			returnErr: true,
			name:      "InvalidCSVFile",
		},
		{
			fileName:  "csv_tests/invalid2.csv",
			returnErr: true,
			name:      "InvalidEmail",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			file, err := os.Open(tc.fileName)

			defer file.Close()
			if err != nil {
				log.Println(err)
				panic(err)

			}
			log.Println("Executing test:" + tc.name)
			_, err = SortDomains(file)
			returnedErr := err != nil

			if returnedErr != tc.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v", tc.returnErr, returnedErr)
			}
		})
	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		fileName string
		name     string
		want     []DomainCount
	}{

		{
			fileName: "csv_tests/valid2.csv",
			name:     "ValidSort",
			want: []DomainCount{
				{
					Domain: "a.com",
					Count:  2,
				},
				{
					Domain: "b.com",
					Count:  1,
				},
				{
					Domain: "c.com",
					Count:  2,
				},
				{
					Domain: "d.com",
					Count:  1,
				},
				{
					Domain: "e.com",
					Count:  3,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			log.Println("Executing test:" + tc.name)

			file, err := os.Open(tc.fileName)
			if err != nil {
				log.Println(err)
				panic(err)

			}

			defer file.Close()

			got, _ := SortDomains(file)

			for i := 0; i < len(got); i++ {
				wantedCounter := tc.want[i].Count
				gotCounted := got[i].Count

				wantedDomain := tc.want[i].Domain
				gotDomain := got[i].Domain
				if wantedCounter != gotCounted {
					t.Fatalf("Expected Count: %v, got: %v", tc.want[i].Count, got[i].Count)

				}
				if strings.Compare(wantedDomain, gotDomain) != 0 {
					t.Fatalf("Expected Domain: %v, got: %v", wantedDomain, gotDomain)

				}

			}
		})
	}
}

func BenchmarkSortDomains(b *testing.B) {
	var fileName = "csv_tests/valid.csv" //TODO: use flag to pass a file path
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		panic(err)

	}

	defer file.Close()
	for i := 0; i < b.N; i++ {
		SortDomains(file)
	}
}
