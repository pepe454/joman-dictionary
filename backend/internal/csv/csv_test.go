package csv

import (
	"fmt"
	"os"
	"testing"
)

func TestReadCsv(t *testing.T) {
	wordsDir := os.Getenv("WORDS_DIR")
	records := ReadCsv(fmt.Sprintf("%s/Adjectives.csv", wordsDir))
	header := records[0]
	fmt.Printf("Header: %v\n", header)
	for i, record := range records[1:] {
		fmt.Printf("Record %d: %v\n", i+1, record)
	}
}

func TestReadWordCsv(t *testing.T) {
	wordsDir := os.Getenv("WORDS_DIR")
	records, err := ReadCsvWordFile(fmt.Sprintf("%s/Adjectives.csv", wordsDir), "Adjective")
	if err != nil || len(records) == 0 {
		t.Errorf("Error reading word csv file: %v", err)
	}

	for _, record := range records {
		fmt.Printf(
			"Sourashtra: %s -> English: %s\n",
			record.SourashtraWord, record.EnglishWord,
		)
	}
}

func TestReadSentenceCsv(t *testing.T) {
	sentencesDir := os.Getenv("SENTENCES_DIR")
	records, err := ReadCsvSentenceFile(fmt.Sprintf("%s/Body-Sentences.csv", sentencesDir))
	if err != nil || len(records) == 0 {
		t.Errorf("Error reading sentence csv file: %v", err)
	}

	for _, record := range records {
		fmt.Printf(
			"Sourashtra: %s -> English: %s\n",
			record.SourashtraSentence, record.EnglishSentence,
		)
	}
}
