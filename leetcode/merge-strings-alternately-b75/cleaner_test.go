package main

import "testing"

func TestEqualCleaner(t *testing.T) {
	// params
	word1 := "abc"
	word2 := "pqr"

	want := "apbqcr"
	got := MergeAlternatelyCleaner(word1, word2)

	if got != want {
		t.Fatalf("[TestEqual - MergeAlternatelyCleaner]: want: %s got: %s", want, got)
	}
}

func TestWord1LongerCleaner(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"

	want := "apbqcd"
	got := MergeAlternatelyCleaner(word1, word2)

	if got != want {
		t.Fatalf("[TestWord1Longer - Cleaner]: want: %s got: %s", want, got)
	}
}

func TestWord2LongerCleaner(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"

	want := "apbqrs"
	got := MergeAlternatelyCleaner(word1, word2)

	if got != want {
		t.Fatalf("[TestWord2Longer - Cleaner]: want: %s got: %s", want, got)
	}
}
