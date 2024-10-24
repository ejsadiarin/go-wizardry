package main

import "fmt"

func main() {
	m := map[string]string{
		"testkey": "testval",
		"twokey":  "twoval",
	}

	var keys []string
	var vals []string
	for k, v := range m {
		keys = append(keys, k)
		vals = append(vals, v)
		fmt.Printf("keys: %s\n", k)
		fmt.Printf("vals: %s\n", v)
	}

	fmt.Printf("\n----SLICE PRINT----\n")
	fmt.Printf("%v\n", keys)
	fmt.Printf("%v\n", vals)
}
