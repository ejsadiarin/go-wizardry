# Merge Alternately - Leetcode #1768 

[2025-01-09 23:41:54]: Done first time - struggled to pass if word1 is longer than word2 without doing some brute force stuff (fokton of if statements) - will get back to this 2 days from now

# NOTES
- use `math.Max` to determine what word is longer
    - this returns the max length of the longer word
- need to `strings.Split` the words to get slices of both words
- loop condition should stop at length of the longer word (use the output of the `math.Max` here)
- loop body should append if and only if the current `i` is less than the length of the word (for both words)
    - so if its i is less then continue appending
    - no need to check if the i is at the length of the shorter word and create another loop starting from that length (of the shorter word)

- finally knew how to unit test (using the `*testing.T`), to write my first go test lmao (without braindead copy paste from external sources) lmao

# Cleaner Way

- get the `len` of both words
- assume that the max length is `word1` or any of the words, put in a variable like `longer`
    - just add an if statement that changes `longer` to `word2`, vice versa 
- use `strings.Builder` because its more memory efficient
- use the `WriteBytes` of the `strings.Builder` to the word to return the merged word
> [!IMPORTANT]
> NOTE: this assumes that 1 char == 1 byte, (try making it compatible with `runes` perhaps?)

- return in string with `.String()`
