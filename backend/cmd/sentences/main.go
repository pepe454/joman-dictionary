package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
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

func identifyCorrectWord(ctx context.Context, q *repository.Queries, word string, language Language) (correctWord string, wordID int32) {
	correctWord, wordID = "", -1
	searchResults, _ := q.WordSearch(ctx, repository.WordSearchParams{Language: string(language), Similarity: word})
	maxDisplay := min(5, len(searchResults))
	for i, result := range searchResults[:maxDisplay] {
		fmt.Printf("%d) Found similar word %s with ID %d\n", i+1, result.WordText, result.WordID)
	}

	// Let the user specify if they want to use a queried set of words or enter their own word.
	var choiceStr string
	fmt.Print("Enter the number of the correct word if listed above, -1 to enter correct word, or 0 to skip: ")
	fmt.Scanln(&choiceStr)
	choice, scanErr := strconv.Atoi(strings.TrimSpace(choiceStr))
	if scanErr != nil {
		fmt.Printf("Error reading input: %v\n", scanErr)
	} else if choice == 0 {
		fmt.Printf("Skipping word %s\n", word)
	} else if choice > 0 {
		if choice > len(searchResults)+1 {
			fmt.Printf("Invalid choice %d. Needed to enter a number between 1 and %d.\n", choice, len(searchResults))
		} else {
			word := searchResults[choice-1]
			correctWord, wordID = word.WordText, word.WordID
		}
	} else {
		var correctWord string
		fmt.Printf("Could not find similar word for %s. Please enter the correct matching word: ", word)
		fmt.Scanln(&correctWord)
		if len(correctWord) > 0 {
			return identifyCorrectWord(ctx, q, correctWord, language)
		}
	}
	return
}

func uploadSentence(
	ctx context.Context, q *repository.Queries, wordMap map[string]int32,
	sentence string, language Language) error {
	headerTitle := fmt.Sprintf("Uploading sentence <%s>", sentence)
	headerSeparator := "\n" + strings.Repeat("-", len(headerTitle)) + "\n"
	header := headerSeparator + headerTitle + headerSeparator
	fmt.Print(header)

	words := strings.SplitSeq(sentence, " ")
	contextWordIDs := make([]int32, 0, 5)

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

		// Try to get user input with a database search if we cannot identify the word.
		if !ok {
			fmt.Printf("Could not identify word: %s. Are any of the following correct?\n", word)
			correctWord, wordID := identifyCorrectWord(ctx, q, word, language)
			if wordID > 0 {
				fmt.Printf("Found word %s with ID %d\n", correctWord, wordID)
				contextWordIDs = append(contextWordIDs, wordID)
			}
		} else {
			fmt.Printf("Found word %s with ID %d\n", word, wordID)
			contextWordIDs = append(contextWordIDs, wordID)
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
