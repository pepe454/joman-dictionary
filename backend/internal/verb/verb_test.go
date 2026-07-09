package verb

import (
	"fmt"
	"testing"
)

func TestGetRoot(t *testing.T) {
	testVerb := "jaattE"
	root, err := GetRoot(testVerb)
	want := "ja"

	fmt.Println("Got the root for jaatE:", testVerb, root, err)
	if root != want {
		t.Errorf("Root for %s doesn't match %s, got: %s", testVerb, want, root)
	}
	if err != nil {
		t.Errorf("Error conjugating verb: %v", err)
	}

	// now test error case -> should not conjugate
	testVerbConjugated := "jaariyo"
	_, rootErr := GetRoot(testVerbConjugated)
	if rootErr == nil {
		t.Errorf("Verb is not error, should not be conjugating")
	}
}

func TestConjugateImperative(t *testing.T) {
	paramsInformal := ConjugationParams{Respect: false, Tense: ImperativeTense}
	paramsPlural := ConjugationParams{Respect: false, Tense: ImperativeTense, Plural: true}
	paramsRespect := ConjugationParams{Respect: true, Tense: ImperativeTense, Plural: false}

	type TestCase struct {
		verb   string
		params ConjugationParams
		want   string
	}

	testCases := []TestCase{
		{"avattE", paramsInformal, "av"},
		{"avattE", paramsPlural, "avo"},
		{"avattE", paramsRespect, "avo"},
		{"hingattE", paramsInformal, "hingi"},
		{"hingattE", paramsRespect, "hinguvo"},
		{"jaattE", paramsRespect, "javo"},
		{"jaattE", paramsInformal, "ja"},
	}

	for _, tc := range testCases {
		got, _ := ConjugateVerb(tc.verb, tc.params)
		if got != tc.want {
			t.Errorf("Conjugation for %s with params %+v should be %s, got: %s", tc.verb, tc.params, tc.want, got)
		}
	}
}
