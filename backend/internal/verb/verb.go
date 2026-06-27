package verb

import (
	"fmt"
	"strings"
)

const INFINITIVE_ENDING = "attE"

type PersonType int

const (
	FirstPerson  PersonType = 1
	SecondPerson PersonType = 2
	ThirdPerson  PersonType = 3
)

type TenseType string

const (
	ImperativeTense TenseType = "imperative"
	GerundTense     TenseType = "gerun"
)

type ConjugationParams struct {
	Male    bool
	Plural  bool
	Respect bool
	Person  PersonType
	Tense   TenseType
}

// GetRoot finds the root of an infinitive verb. If the verb is not infinitive it cannot be conjugated.
func GetRoot(infinitive string) (string, error) {
	root, found := strings.CutSuffix(infinitive, "attE")
	var conjugationError error
	if !found {
		conjugationError = fmt.Errorf(
			"Cannnot find root for verb %s, it does not have the infinitive ending (%s)",
			infinitive, INFINITIVE_ENDING,
		)
	}
	return root, conjugationError
}

// ConjugateImperative is the simplest of all - just return the root and + "uvo" for respect.
func ConjugateImperative(root string, params ConjugationParams) string {
	if params.Respect || params.Plural {
		if strings.HasSuffix(root, "v") { // av -> avo
			return root + "o"
		}
		return root + "uvo" // aas -> aasuvo
	}
	return root
}

// Gerund is very simple - return root + "ile sE"
func ConjugateGerund(root string, params ConjugationParams) string {
	return root + "ile sE"
}

// ConjugateVerb takes an infinitive and params and conjugates the verb accordingly.
func ConjugateVerb(infinitive string, params ConjugationParams) (conjugatedVerb string, conjugationErr error) {
	fmt.Printf("Conjugating %s for params %+v\n", infinitive, params)

	// Setup the root first
	root, rootErr := GetRoot(infinitive)
	if rootErr != nil {
		conjugationErr = rootErr
		return
	}

	// Based on the tense , call different functions
	switch params.Tense {
	case ImperativeTense:
		conjugatedVerb = ConjugateImperative(root, params)
	case GerundTense:
		conjugatedVerb = ConjugateGerund(root, params)
	}
	return
}
