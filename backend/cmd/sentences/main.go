package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pepe454/joman-dictionary/internal/csv"
	"github.com/pepe454/joman-dictionary/internal/db"
	"github.com/pepe454/joman-dictionary/internal/repository"
)

type Language string

const (
	English    Language = "English"
	Sourashtra Language = "Sourashtra"
)

func uploadSentence(
	ctx context.Context, q *repository.Queries, wordMap map[string]int32,
	sentence string, language Language) error {
	headerTitle := fmt.Sprintf("Uploading sentence <%s>", sentence)
	headerSeparator := "\n" + strings.Repeat("-", len(headerTitle)) + "\n"
	header := headerSeparator + headerTitle + headerSeparator
	fmt.Print(header)

	words := strings.SplitSeq(sentence, " ")
	for word := range words {
		// strip the punctuation from the word
		word = strings.Trim(word, ".,!?;:\"'()[]{}")
		decapitalized := strings.ToLower(word[0:1]) + word[1:]

		var depluralized string
		if language == English {
			depluralized = strings.TrimSuffix(decapitalized, "s")
		} else {
			depluralized = strings.TrimSuffix(decapitalized, "n")
		}
		candidates := []string{word, decapitalized, depluralized}

		var wordID int32
		var ok bool
		for _, candidate := range candidates {
			wordID, ok = wordMap[candidate]
			if ok {
				break
			}
		}
		if !ok {
			fmt.Printf("Could not identify word %s\n", word)
		} else {
			fmt.Printf("Found word %s with ID %d\n", word, wordID)
		}
	}
	return nil
}

func uploadCsvSentenceFile(ctx context.Context, q *repository.Queries, filePath string, sourashtraWordMap, englishWordMap map[string]int32) error {
	sentenceRecords, csvErr := csv.ReadCsvSentenceFile(filePath)
	if csvErr != nil {
		fmt.Printf("Error reading sentence csv file: %v\n", csvErr)
		return csvErr
	}
	fmt.Printf("Uploading csv file <%s>...\n", filePath)

	for _, sentenceRecord := range sentenceRecords {
		uploadSentence(ctx, q, sourashtraWordMap, sentenceRecord.SourashtraSentence, Sourashtra)
		uploadSentence(ctx, q, englishWordMap, sentenceRecord.EnglishSentence, English)
	}

	return nil
}

// uploadCsvSentenceFiles uploads all the csv files in the sentences/ directory
func uploadCsvSentenceFiles(
	ctx context.Context, q *repository.Queries,
	sourashtraWordMap, englishWordMap map[string]int32) error {
	sentencesDir := os.Getenv("SENTENCES_DIR")
	csvFiles := []string{
		// path.Join(sentencesDir, "Body-Sentences.csv"),
		path.Join(sentencesDir, "Greeting-Phrases.csv"),
	}

	for _, filePath := range csvFiles {
		uploadCsvSentenceFile(ctx, q, filePath, sourashtraWordMap, englishWordMap)
	}
	return nil
}

func main() {
	ctx := context.Background()
	pool, err := db.Connect(ctx)
	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}
	q := repository.New(pool)

	// Load sourashtra and english words into maps for quick lookup
	sourashtraWords, err := q.ListWordsForLanguage(ctx, string(Sourashtra))
	if err != nil {
		log.Panic("Failed to list sourashtra words: ", err)
	}
	sourashtraWordMap := make(map[string]int32)
	for _, word := range sourashtraWords {
		sourashtraWordMap[word.WordText] = word.WordID
	}
	fmt.Printf("Loaded %d sourashtra Words\n", len(sourashtraWords))

	englishWords, err := q.ListWordsForLanguage(ctx, string(English))
	if err != nil {
		log.Panic("Failed to list english words: ", err)
	}
	englishWordMap := make(map[string]int32)
	for _, word := range englishWords {
		englishWordMap[word.WordText] = word.WordID
	}
	fmt.Printf("Loaded %d english Words\n", len(englishWords))

	uploadCsvSentenceFiles(ctx, q, sourashtraWordMap, englishWordMap)
}
