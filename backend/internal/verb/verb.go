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
	GerundTense     TenseType = "gerund"
)

type ConjugationParams struct {
	Male    bool
	Plural  bool
	Respect bool
	Person  PersonType
	Tense   TenseType
}

const Vowels = "aeiou"

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

// ConjugateImperative tense is for commands and requests.
func ConjugateImperative(root string, params ConjugationParams) string {
	suffix := string(root[len(root)-1])
	if params.Respect || params.Plural {
		if suffix == "v" { // av -> avo
			return root + "o"
		} else if strings.ContainsAny(suffix, Vowels) { // ja -> javo
			return root + "vo"
		}
		return root + "uvo" // aas -> aasuvo
	} else if suffix == "v" { // av -> av
		return root
	} else if strings.ContainsAny(suffix, Vowels) { // ja -> ja
		return root
	}
	return root + "i" // hing -> hingi
}

// Gerund is very simple - return root + "ile"
func ConjugateGerund(root string, params ConjugationParams) string {
	root = strings.TrimRight(root, "a") // for roots ending in a like "ja" -> "jile"
	return root + "ile"
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
