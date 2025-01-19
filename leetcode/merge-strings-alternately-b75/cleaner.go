package main

import "strings"

func MergeAlternatelyCleaner(word1 string, word2 string) string {
	a := len(word1)
	b := len(word2)
	ln := a
	if b > a {
		ln = b
	}
	var builder strings.Builder

	for i := 0; i < ln; i++ {
		if i < a {
			builder.WriteByte(word1[i])
		}
		if i < b {
			builder.WriteByte(word2[i])
		}
	}
	return builder.String()
}
