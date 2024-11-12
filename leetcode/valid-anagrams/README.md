# [20241112 9:50PM]:

## isAnagram
made a map of rune (keys) with int (value), to indicate occurrence of the rune/char

loop the string s to construct the map

another loop for comparison, goal is to decrement the values in the map until 0 (assuming the key is the same)

always return false if:
- length of s != length of t
- length of map (keys) exceeds length of s (original length)
- the values in the map became negative

otherwise, return true for everything.

## isAnagram2

make both strings into slice of substrings
then sort them both

make the slice of substrings into a string again by joining them 
then compare results, if equal then return true, otherwise false
