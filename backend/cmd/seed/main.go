package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pepe454/joman-dictionary/internal/csv"
	"github.com/pepe454/joman-dictionary/internal/db"
	"github.com/pepe454/joman-dictionary/internal/repository"
)

type CSVFile struct {
	FilePath            string
	DefaultPartOfSpeech string
	CategoryID          int
}

func getOrInsertWord(ctx context.Context, q *repository.Queries, params repository.InsertWordParams) (int32, error) {
	id, err := q.GetWordID(ctx, repository.GetWordIDParams{
		WordText: params.WordText,
		Language: params.Language,
	})
	if err == nil {
		return id, nil
	}
	return q.InsertWord(ctx, params)
}

// uploadWordPair takes a pair of sourashtra word + english word and uploads them to the Database
func uploadWordPair(ctx context.Context, q *repository.Queries, wordRecord csv.WordRecord, categoryID int) error {
	// Step 1: Get or insert Sourashtra word
	sourashtraID, err := getOrInsertWord(ctx, q, repository.InsertWordParams{
		WordText:     wordRecord.SourashtraWord,
		Language:     "Sourashtra",
		PartOfSpeech: wordRecord.PartOfSpeech,
	})
	if err != nil {
		return err
	}

	// Step 2: Get or insert English word
	englishID, err := getOrInsertWord(ctx, q, repository.InsertWordParams{
		WordText:     wordRecord.EnglishWord,
		Language:     "English",
		PartOfSpeech: wordRecord.PartOfSpeech,
	})
	if err != nil {
		return err
	}

	// Step 3: Insert Translation if not already present
	_, getErr := q.GetTranslation(ctx, repository.GetTranslationParams{
		SourashtraWordID: sourashtraID,
		OtherWordID:      englishID,
	})
	if getErr != nil {
		err = q.InsertTranslation(ctx, repository.InsertTranslationParams{
			SourashtraWordID: sourashtraID,
			OtherWordID:      englishID,
			Context:          pgtype.Text{String: wordRecord.TranslationContext, Valid: wordRecord.TranslationContext != ""},
		})
		if err != nil {
			return err
		}
	}

	// Step 4: Insert Word Category if categoryID is valid (> 0) and not already present
	if categoryID > 0 {
		_, getErr := q.GetWordCategory(ctx, repository.GetWordCategoryParams{
			WordID:     sourashtraID,
			CategoryID: int32(categoryID),
		})
		if getErr != nil {
			err = q.InsertWordCategory(ctx, repository.InsertWordCategoryParams{
				WordID:     sourashtraID,
				CategoryID: int32(categoryID),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// initWordCategoryMap reads the categories from the table and creates a mapping from category string to category id.
func initWordCategoryMap(ctx context.Context, q *repository.Queries) (map[string]int, error) {
	categories, err := q.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	categoryMap := make(map[string]int)
	for _, category := range categories {
		categoryMap[category.Category] = int(category.CategoryID)
	}
	return categoryMap, nil
}

func uploadCSVFile(ctx context.Context, q *repository.Queries, f CSVFile) error {
	wordRecords, csvErr := csv.ReadCsvWordFile(f.FilePath, f.DefaultPartOfSpeech)
	if csvErr != nil {
		return csvErr
	}

	headerTitle := fmt.Sprintf(
		"Uploading csv file <%s> with PoS <%s> and category <%d>",
		f.FilePath, f.DefaultPartOfSpeech, f.CategoryID,
	)
	headerSeparator := "\n" + strings.Repeat("-", len(headerTitle)) + "\n"
	header := headerSeparator + headerTitle + headerSeparator
	fmt.Print(header)

	for _, wordRecord := range wordRecords[1:] {
		fmt.Printf(
			"Uploading pair: %s (%s), %s %s\n",
			wordRecord.SourashtraWord, wordRecord.PartOfSpeech,
			wordRecord.EnglishWord, wordRecord.TranslationContext,
		)
		err := uploadWordPair(ctx, q, wordRecord, f.CategoryID)
		if err != nil {
			log.Printf("Failed to upload word pair (%s, %s): %v", wordRecord.SourashtraWord, wordRecord.EnglishWord, err)
		}
	}
	return nil
}

// uploadCSVFiles uploads all the csv files
func uploadCSVFiles(ctx context.Context, q *repository.Queries, categoryMap map[string]int) error {
	wordsDir := os.Getenv("WORDS_DIR")

	// Setup CSV Files
	csvFiles := []CSVFile{
		{path.Join(wordsDir, "Adjectives.csv"), "adjective", categoryMap["adjectives"]},
		{path.Join(wordsDir, "Adverbs.csv"), "adverb", categoryMap["adverbs"]},
		{path.Join(wordsDir, "Animals.csv"), "noun", categoryMap["animals"]},
		{path.Join(wordsDir, "Body.csv"), "noun", categoryMap["body"]},
		{path.Join(wordsDir, "Business.csv"), "noun", categoryMap["business"]},
		{path.Join(wordsDir, "Clothing.csv"), "noun", categoryMap["clothing"]},
		{path.Join(wordsDir, "Knowledge.csv"), "noun", categoryMap["knowledge"]},
		{path.Join(wordsDir, "Colors.csv"), "adjective", categoryMap["colors"]},
		{path.Join(wordsDir, "Conjunctions.csv"), "conjunction", categoryMap["conjunctions"]},
		{path.Join(wordsDir, "Family.csv"), "noun", categoryMap["family"]},
		{path.Join(wordsDir, "Feelings.csv"), "noun", categoryMap["feelings"]},
		{path.Join(wordsDir, "Food.csv"), "noun", categoryMap["food"]},
		{path.Join(wordsDir, "House.csv"), "noun", categoryMap["house"]},
		{path.Join(wordsDir, "Nature.csv"), "noun", categoryMap["nature"]},
		{path.Join(wordsDir, "Numbers.csv"), "numeral", categoryMap["numbers"]},
		{path.Join(wordsDir, "People.csv"), "noun", categoryMap["people"]},
		{path.Join(wordsDir, "Places.csv"), "noun", categoryMap["places"]},
		{path.Join(wordsDir, "Possessives.csv"), "pronoun", categoryMap["pronouns"]},
		{path.Join(wordsDir, "Prepositions.csv"), "preposition", categoryMap["prepositions"]},
		{path.Join(wordsDir, "Pronouns.csv"), "pronoun", categoryMap["pronouns"]},
		{path.Join(wordsDir, "Questions.csv"), "interrogative", categoryMap["questions"]},
		{path.Join(wordsDir, "Taste.csv"), "adjective", categoryMap["taste"]},
		{path.Join(wordsDir, "Time.csv"), "noun", categoryMap["time"]},
		{path.Join(wordsDir, "TimeAdverbs.csv"), "adverb", categoryMap["time"]},
		{path.Join(wordsDir, "Verbs.csv"), "verb", categoryMap["verbs"]},
	}

	for _, csvFile := range csvFiles {
		csvErr := uploadCSVFile(ctx, q, csvFile)
		if csvErr != nil {
			log.Printf("Failed to upload csv file(%s): %v", csvFile.FilePath, csvErr)
		}
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
	categoryMap, err := initWordCategoryMap(ctx, q)

	if err != nil {
		log.Panic("Failed to initialize category map: ", err)
	}

	for category, id := range categoryMap {
		fmt.Printf("Category: %s, ID: %d\n", category, id)
	}

	uploadCSVFiles(ctx, q, categoryMap)
}
