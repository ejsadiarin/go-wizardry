package main

import "testing"

func TestEqual(t *testing.T) {
	// params
	word1 := "abc"
	word2 := "pqr"

	want := "apbqcr"
	got := MergeAlternately(word1, word2)

	if got != want {
		t.Fatalf("[TestEqual]: want: %s got: %s", want, got)
	}
}

func TestWord1Longer(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"

	want := "apbqcd"
	got := MergeAlternately(word1, word2)

	if got != want {
		t.Fatalf("[TestWord1Longer]: want: %s got: %s", want, got)
	}
}

func TestWord2Longer(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"

	want := "apbqrs"
	got := MergeAlternately(word1, word2)

	if got != want {
		t.Fatalf("[TestWord2Longer]: want: %s got: %s", want, got)
	}
}
