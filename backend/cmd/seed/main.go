package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pepe454/joman-dictionary/internal/csv"
	"github.com/pepe454/joman-dictionary/internal/db"
	"github.com/pepe454/joman-dictionary/internal/repository"
)

type CSVFile struct {
	filepath     string
	partOfSpeech string
	categoryID   int
}

// uploadWordPair takes a pair of sourashtra word + english word and uploads them to the Database
// It also uploads the translation between them, and the category as well.
func uploadWordPair(ctx context.Context, pool *pgxpool.Pool, q *repository.Queries, sourashtraText, englishText, partOfSpeech, translationContext string, categoryID int) error {
	// Step 1 : Setup a transaction
	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Panic("Failed to begin transaction: ", err)
	}
	defer tx.Rollback(ctx)
	qtx := q.WithTx(tx)

	// Step 2: Insert Sourashtra Word
	sourashtraWordParams := repository.InsertWordParams{
		WordText:     sourashtraText,
		Language:     "Sourashtra",
		PartOfSpeech: partOfSpeech,
	}
	sourashtraID, err := qtx.InsertWord(ctx, sourashtraWordParams)
	if err != nil {
		return err
	}

	// Step 3: Insert English Word
	englishWordParams := repository.InsertWordParams{
		WordText:     englishText,
		Language:     "English",
		PartOfSpeech: partOfSpeech,
	}
	englishID, err := qtx.InsertWord(ctx, englishWordParams)
	if err != nil {
		return err
	}

	// Step 4: Insert Translation
	translationParams := repository.InsertTranslationParams{
		SourashtraWordID: sourashtraID,
		OtherWordID:      englishID,
		Context:          pgtype.Text{String: translationContext, Valid: translationContext != ""},
	}
	err = qtx.InsertTranslation(ctx, translationParams)
	if err != nil {
		return err
	}

	// Step 5: Insert Word Category if categoryID is valid (> 0)
	if categoryID > 0 {
		wordCategoryParams := repository.InsertWordCategoryParams{
			WordID:     sourashtraID,
			CategoryID: int32(categoryID),
		}
		err = qtx.InsertWordCategory(ctx, wordCategoryParams)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
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

// uploadSimpleWordPair takes a simple csv file with no context
func uploadCSVFile(ctx context.Context, pool *pgxpool.Pool, q *repository.Queries, f CSVFile) error {
	records := csv.ReadCsv(f.filepath)
	headerTitle := fmt.Sprintf("Uploading csv file <%s> with PoS <%s> and category <%d>", f.filepath, f.partOfSpeech, f.categoryID)
	headerSeparator := "\n" + strings.Repeat("-", len(headerTitle)) + "\n"
	header := headerSeparator + headerTitle + headerSeparator
	fmt.Print(header)

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		sourashtra := record[0]
		english := record[1]
		translationContext := ""

		if len(record) > 2 {
			translationContext = record[2]
		}
		fmt.Printf("Uploading pair: %s, %s %s\n", sourashtra, english, translationContext)

		// err := uploadWordPair(ctx, pool, q, sourashtra, english, f.partOfSpeech, translationContext, f.categoryID)
		// if err != nil {
		// 	log.Printf("Failed to upload word pair (%s, %s): %v", sourashtra, english, err)
		// }
	}
	return nil
}

// uploadCSVFiles uploads all the csv files
func uploadCSVFiles(ctx context.Context, pool *pgxpool.Pool, q *repository.Queries, categoryMap map[string]int) error {
	wordsDir := os.Getenv("WORDS_DIR")

	// Setup CSV Files
	csvFiles := []CSVFile{
		{path.Join(wordsDir, "Adjectives.csv"), "adjective", categoryMap["adjectives"]},
		{path.Join(wordsDir, "Adverbs.csv"), "adverb", categoryMap["adverbs"]},
		{path.Join(wordsDir, "Animals.csv"), "noun", categoryMap["animals"]},
		{path.Join(wordsDir, "Body.csv"), "noun", categoryMap["body"]},
		{path.Join(wordsDir, "Clothing.csv"), "noun", categoryMap["clothing"]},
		{path.Join(wordsDir, "Cognition.csv"), "noun", categoryMap["cognition"]},
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
		csvErr := uploadCSVFile(ctx, pool, q, csvFile)
		if csvErr != nil {
			log.Printf("Failed to upload csv file(%s): %v", csvFile.filepath, csvErr)
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

	uploadCSVFiles(ctx, pool, q, categoryMap)
}
