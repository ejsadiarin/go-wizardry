package main

import (
	"sort"
	"strings"
)

// Time Complexity: O(n)
// Space Complexity: ?
func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	m := make(map[rune]int, len(s))

	for _, sv := range s {
		m[sv] += 1 // record the occurrence of rune/char
	}
	for _, st := range t {
		m[st] -= 1
		if len(m) > len(s) || m[st] < 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(tt)

	return strings.Join(ss, "") == strings.Join(tt, "")
}
