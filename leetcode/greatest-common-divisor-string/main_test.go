package main

import (
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"AAAA", "AA", "AA"},
		{"ABCD", "ABCD", "ABCD"},
		{"", "", ""},
	}

	for _, test := range tests {
		result := gcdOfStrings(test.str1, test.str2)
		if result != test.expected {
			t.Errorf("gcdOfStrings(%q, %q) = %q; expected %q", test.str1, test.str2, result, test.expected)
		}
	}
}
