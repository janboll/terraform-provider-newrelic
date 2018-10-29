package newrelic

import (
	"testing"
)

func TestParseIDs_Basic(t *testing.T) {
	ids, err := parseIDs("1:2", 2)
	if err != nil {
		t.Fatal(err)
	}

	if len(ids) != 2 {
		t.Fatal(len(ids))
	}

	if ids[0] != 1 || ids[1] != 2 {
		t.Fatal(ids)
	}
}

func TestSerializeIDs_Basic(t *testing.T) {
	id := serializeIDs([]int{1, 2})

	if id != "1:2" {
		t.Fatal(id)
	}
}

func TestSortedTerms_Warn(t *testing.T) {
	testTerms := make([]map[string]interface{}, 0)
	testTerm := map[string]interface{}{
		"priority": "warning",
	}
	testTerms = append(testTerms, testTerm)
	testTerms = sortedTerms(testTerms)

	if testTerms[0]["priority"] != "warning" {
		t.Fatal("First element should be warning!")
	}
	if len(testTerms) != 1 {
		t.Fatal("Lenght return should have been 1!")
	}

}

func TestAsdasd(t *testing.T) {
	testTerms := make([]map[string]interface{}, 0)
	testTermW := map[string]interface{}{
		"priority": "warning",
	}
	testTermC := map[string]interface{}{
		"priority": "critical",
	}

	testTerms = append(testTerms, testTermC)
	testTerms = append(testTerms, testTermW)
	testTerms = sortedTerms(testTerms)

	if testTerms[0]["priority"] != "warning" {
		t.Fatal("First element should be warning!")
	}
	if testTerms[1]["priority"] != "critical" {
		t.Fatal("Second element should be critical!")
	}
	if len(testTerms) != 2 {
		t.Fatal("Lenght return should have been 1!")
	}

}
