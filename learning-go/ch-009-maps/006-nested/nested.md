# Nested

Maps can contain maps, creating a nested structure. For example:

```go
map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int
```

# Assignment

Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the `getNameCounts` function.

- It takes a slice of strings `names` and returns a nested map.
- The parent map's keys are all the unique first characters (see [runes](https://go.dev/blog/strings)) of the `names`,
  - the nested maps keys are all the `names` themselves, and the value is the count of each name.

For example:

```txt
billy
billy
bob
joe
```

Creates the following nested map:

```txt

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
```

# Tips

- The test suite is not printing the map you're returning directly, but instead checking a few specific keys.
- You can convert a string into a slice of runes as follows:

```go
runes := []rune(word)
```
