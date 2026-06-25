// csv
package csv

import (
	"encoding/csv"
	"log"
	"os"
)

// ReadCSV uses the encoding/csv package to read csv file and return a 2D array of strings
func ReadCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
