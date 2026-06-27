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
	testVerb := "avattE"

	// Test non-respect
	paramsInformal := ConjugationParams{Respect: false, Tense: ImperativeTense}
	informal, _ := ConjugateVerb(testVerb, paramsInformal)
	informalWant := "av"
	if informal != informalWant {
		t.Errorf("Informal commmand for %s doesn't match %s, got: %s", testVerb, informalWant, informal)
	}

	// Test plural without respect
	paramsPlural := ConjugationParams{Respect: false, Tense: ImperativeTense, Plural: true}
	plural, _ := ConjugateVerb(testVerb, paramsPlural)
	pluralWant := "avo"
	if plural != pluralWant {
		t.Errorf("Plural commmand for %s doesn't match %s, got: %s", testVerb, pluralWant, plural)
	}

	// Now test formal-> same as plural
	paramsFormal := ConjugationParams{Respect: true, Tense: ImperativeTense, Plural: false}
	formal, _ := ConjugateVerb(testVerb, paramsFormal)
	formalWant := "avo"
	if formal != formalWant {
		t.Errorf("formal commmand for %s doesn't match %s, got: %s", testVerb, formalWant, formal)
	}

	// informal not ending in v
	hingi, _ := ConjugateVerb("hingattE", paramsInformal)
	if hingi != "hingi" {
		t.Errorf("Informal command for hingattE should be hingi, got: %s", hingi)
	}

}
