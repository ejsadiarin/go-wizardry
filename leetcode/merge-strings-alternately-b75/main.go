package main

import (
	"math"
	"strings"
)

func MergeAlternately(word1 string, word2 string) string {
	var m []string
	w1 := strings.Split(word1, "")
	w2 := strings.Split(word2, "")
	wordMax := math.Max(float64(len(word1)), float64(len(word2)))
	for i := 0; i < int(wordMax); i++ {
		if i < len(word1) {
			m = append(m, w1[i])
		}
		if i < len(word2) {
			m = append(m, w2[i])
		}
	}

	return strings.Join(m, "")
}
