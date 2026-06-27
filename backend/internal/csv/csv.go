// csv
package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type WordRecord struct {
	SourashtraWord     string
	EnglishWord        string
	TranslationContext string
	PartOfSpeech       string
}

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

// ReadCsvWordFile returns a slice of structs with translations from sourashtra to english.
func ReadCsvWordFile(filePath string, defaultPartOfSpeech string) ([]WordRecord, error) {
	var csvErr error
	records := ReadCsv(filePath)
	wordRecords := make([]WordRecord, 0, len(records)-1) // exclude header

	// Setup header mapping to get the index of each column
	header := records[0]
	headerMapping := make(map[string]int)
	for i, h := range header {
		headerMapping[h] = i
	}

	// Confirm sourashtra and english words available
	sourashtraWordIndex, sourashtraOk := headerMapping["sourashtra_word"]
	englishWordIndex, englishOk := headerMapping["english_word"]
	if !sourashtraOk || !englishOk {
		csvErr = fmt.Errorf("Could not find index for sourashtra or english word in csv file.")
		return wordRecords, csvErr
	}

	// populate the wordRecords slice with sourashtra->english translations
	for _, record := range records {
		sourashtraWord := record[sourashtraWordIndex]
		englishWord := record[englishWordIndex]
		translationContext := ""
		partOfSpeech := defaultPartOfSpeech

		translationIndex, translationOk := headerMapping["optional_context"]
		if translationOk {
			translationContext = record[translationIndex]
		}

		partOfSpeechIndex, partOfSpeechOk := headerMapping["part_of_speech"]
		if partOfSpeechOk {
			partOfSpeech = record[partOfSpeechIndex]
		}
		wordRecord := WordRecord{sourashtraWord, englishWord, translationContext, partOfSpeech}
		wordRecords = append(wordRecords, wordRecord)
	}

	return wordRecords, csvErr
}
