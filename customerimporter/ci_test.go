package customerimporter

import (
	"log"
	"testing"
)

func TestCSVErrorHandling(t *testing.T) {
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
			log.Println("Executing test:" + tc.name)
			_, err := SortDomains2(tc.fileName)
			returnedErr := err != nil

			if returnedErr != tc.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v", tc.returnErr, returnedErr)
			}
		})
	}
}

func BenchmarkSortDomains2(b *testing.B) {
	var fileName = "csv_tests/valid.csv"
	for i := 0; i < b.N; i++ {
		SortDomains2(fileName)
	}
}
