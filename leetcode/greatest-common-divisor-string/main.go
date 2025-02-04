/**
1071. Greatest Common Divisor String

For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

Constraints:
    (1) 1 <= str1.length, str2.length <= 1000
    (2) str1 and str2 consist of English uppercase letters.
*
* */

package main

/*
O(log(min(a,b)))

Base case: if b == 0

Recursive case: replace a with b, and b = a%b
--> until b equates to 0

Returns a as the length at which the longer string (in this case str1 or a) is "repeating"
*/
func gcd(a, b int) int {
	// uses Euclidean Algorithm
	for b != 0 {
		a, b = b, a%b // modulo is used for replacing the larger number with a % b
	}
	return a
}

// Assume that str1 is always the longest string
func gcdOfStrings(str1, str2 string) string {
	// if concatenating strings in different orders does produce different results, then no common divisor exists
	if str1+str2 != str2+str1 {
		return ""
	}

	// the gcdLength is the index at which the str1 is considered as "repeating"
	gcdLength := gcd(len(str1), len(str2))

	// the common divisor string is the prefix of str1, basically getting the difference
	return str1[:gcdLength]
}
