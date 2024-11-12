package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"listen", "silent", true},
		{"triangle", "integral", true},
		{"apple", "papel", true},
		{"rat", "car", false},
		{"a", "aa", false},
		{"ab", "a", false},
		{"", "", true},
	}

	for _, tt := range tests {
		result := isAnagram(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("isAnagram(%q, %q) = %v; expected %v", tt.s, tt.t, result, tt.expected)
		}
	}
}

func TestIsAnagram2(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"listen", "silent", true},
		{"triangle", "integral", true},
		{"apple", "papel", true},
		{"rat", "car", false},
		{"a", "aa", false},
		{"ab", "a", false},
		{"", "", true},
	}

	for _, tt := range tests {
		result := isAnagram2(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("isAnagram2(%q, %q) = %v; expected %v", tt.s, tt.t, result, tt.expected)
		}
	}
}
