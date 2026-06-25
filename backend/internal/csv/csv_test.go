package csv

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	wordsDir := os.Getenv("WORDS_DIR")
	records := ReadCsv(fmt.Sprintf("%s/Adjectives.csv", wordsDir))
	header := records[0]
	fmt.Printf("Header: %v\n", header)
	for i, record := range records[1:] {
		fmt.Printf("Record %d: %v\n", i+1, record)
	}
}
